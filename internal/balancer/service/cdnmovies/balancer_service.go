package cdnmovies

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"errors"
	"fmt"
	tmdb "github.com/cyruzin/golang-tmdb"
	"github.com/go-chi/render"
	"github.com/gosimple/slug"
	"go-hdflex/internal/balancer/_struct"
	"go-hdflex/internal/balancer/service/helpers"
	"go-hdflex/internal/database/enum"
	"go-hdflex/internal/database/model"
	"go-hdflex/internal/database/repository"
	cr "go-hdflex/internal/database/repository/content"
	"log/slog"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"sync"
	"time"
)

type ResponseData struct {
	Data       []_struct.CdnMoviesContent `json:"data"`
	Pagination struct {
		NextPageUrl *string `json:"next_page_url"`
	} `json:"pagination"`
}

type BalancerService struct {
	log *slog.Logger

	cr         cr.ContentRepositoryInterface
	csr        repository.CreditRepositoryInterface
	tr         repository.TranslationRepositoryInterface
	gr         repository.GenreRepositoryInterface
	ctr        repository.CountryRepositoryInterface
	tmdbClient *tmdb.Client
	image      helpers.ImageServiceInterface
}

type BalancerServiceInterface interface {
	Parse(context.Context)
	Get(string) (ResponseData, error)
}

func NewBalancerService(
	log *slog.Logger,

	cr cr.ContentRepositoryInterface,
	csr repository.CreditRepositoryInterface,
	tr repository.TranslationRepositoryInterface,
	gr repository.GenreRepositoryInterface,
	ctr repository.CountryRepositoryInterface,
	tmdbClient *tmdb.Client,
	image helpers.ImageServiceInterface,
) *BalancerService {
	return &BalancerService{
		log:        log,
		cr:         cr,
		csr:        csr,
		tr:         tr,
		gr:         gr,
		ctr:        ctr,
		tmdbClient: tmdbClient,
		image:      image,
	}
}

func (p *BalancerService) Parse(ctx context.Context) {
	const op = "BalancerService.Parse() ->"

	log := p.log.With(
		"operation", op,
	)

	page := 1

	cdnmoviesUrl := fmt.Sprintf(
		"https://api.cdnmovies.net/v1/contents?order=created_at&order_type=desc&token=%s&page=%d",
		os.Getenv("CDNMOVIES_TOKEN"),
		page,
	)

	for {
		select {
		case <-ctx.Done():
			return
		default:
			responseData, err := p.Get(cdnmoviesUrl)
			if err != nil {
				log.Error("Error getting data", slog.Any("err", err))
				return
			}

			for _, data := range responseData.Data {
				_, err := p.cr.GetByCdnMoviesId(ctx, data.Id)
				if err != nil {
					if errors.Is(err, sql.ErrNoRows) {
						_, err := p.CreateContent(ctx, data)
						if err != nil {
							log.Error("Error creating content", slog.Any("err", err))
							return
						}
					}
				}
			}

			if responseData.Pagination.NextPageUrl == nil {
				break
			}

			page++

			cdnmoviesUrl = fmt.Sprintf(
				"https://api.cdnmovies.net/v1/contents?order=created_at&order_type=desc&token=%s&page=%d",
				os.Getenv("CDNMOVIES_TOKEN"),
				page,
			)
		}
	}
}

func (p *BalancerService) CreateContent(ctx context.Context, data _struct.CdnMoviesContent) (int64, error) {
	const op = "BalancerService.CreateContent() ->"

	wp, _ := time.Parse(time.RFC3339, data.WorldPremiere)
	rp, _ := time.Parse(time.RFC3339, data.RuPremiere)

	lgbt := false
	if data.Lgbt == 1 {
		lgbt = true
	}

	rand.NewSource(time.Now().UnixNano())

	content := model.Content{
		CdnMoviesId:   data.Id,
		RuTitle:       data.RuTitle,
		OrigTitle:     data.OrigTitle,
		EnTitle:       data.EnTitle,
		Url:           fmt.Sprintf("%d-%s", rand.Intn(1*1000*1000), slug.Make(data.RuTitle)),
		Slogan:        data.Slogan,
		Description:   data.Description,
		Duration:      data.Duration,
		IframeUrl:     data.IframeSrc,
		ContentType:   enum.CdnMoviesContentType(data.ContentType),
		Year:          data.Year,
		RatingAge:     data.RatingAge,
		RatingMpaa:    data.RatingMpaa,
		WorldPremiere: &wp,
		RuPremiere:    &rp,
		LastSeason:    data.LastSeason,
		LastEpisode:   data.LastEpisode,
		Lgbt:          lgbt,
	}

	if data.KinopoiskId != 0 {
		posterUrl, err := url.Parse(fmt.Sprintf("http://www.kinopoisk.ru/images/film_big/%d.jpg", data.KinopoiskId))
		if err != nil {
			return 0, fmt.Errorf("%s failed to parse kinopoisk poster url: %w", op, err)
		}

		content.Poster, _ = p.image.Download(_struct.Image{
			Url: *posterUrl,
		})
	}

	var similar []int64

	if data.TmdbId > 0 {
		var poster string
		var backdrop string
		var err error

		if content.ContentType.IsMovieContentType() {
			poster, backdrop, err = p._getImagesFromTmdbForMovies(data)
			if err != nil {
				p.log.Error("Failed to get images from tmdb", fmt.Errorf("%s %w", op, err))
			}

			similar, err = p._getSimilarForMovies(ctx, data)
			if err != nil {
				p.log.Error("Failed to get similar movies", fmt.Errorf("%s %w", op, err))
			}

			tmdbInfo, err := p.tmdbClient.GetMovieDetails(data.TmdbId, nil)
			if err != nil {
				p.log.Error("Failed to get movie details", fmt.Errorf("%s %w", op, err))
			}

			if tmdbInfo != nil {
				content.Popularity = float64(tmdbInfo.Popularity)
				content.Duration = tmdbInfo.Runtime

				if content.WorldPremiere.Second() == 0 {
					date, err := time.Parse("2006-01-02", tmdbInfo.ReleaseDate)
					if err != nil {
						p.log.Error("Failed to parse world premiere date", fmt.Errorf("%s %w", op, err))
					}

					content.WorldPremiere = &date
				}
			}
		} else if content.ContentType.IsTvSeriesContentType() {
			poster, backdrop, err = p._getImagesFromTmdbForTvSeries(data)
			if err != nil {
				return 0, fmt.Errorf("%s failed to get images from tmdb: %w", op, err)
			}

			similar, err = p._getSimilarForTvSeries(data)
			if err != nil {
				p.log.Error("Failed to get similar tv series", fmt.Errorf("%s %w", op, err))
			}

			tmdbInfo, err := p.tmdbClient.GetTVDetails(data.TmdbId, nil)
			if err != nil {
				p.log.Error("Failed to get tv series details", fmt.Errorf("%s %w", op, err))
			}

			if tmdbInfo != nil {
				content.Popularity = float64(tmdbInfo.Popularity)

				if len(tmdbInfo.EpisodeRunTime) != 0 {
					var avgDuration int
					for _, s := range tmdbInfo.EpisodeRunTime {
						avgDuration += s
					}

					content.Duration = avgDuration / len(tmdbInfo.EpisodeRunTime)
				}

				if content.WorldPremiere.Second() == 0 {
					date, err := time.Parse("2006-01-02", tmdbInfo.FirstAirDate)
					if err != nil {
						p.log.Error("Failed to parse world premiere date", fmt.Errorf("%s %w", op, err))
					}

					content.WorldPremiere = &date
				}
			}
		}

		if content.Poster == "" {
			content.Poster = poster
		}

		content.Backdrop = backdrop
	}

	id, err := p.cr.Create(ctx, content)
	if err != nil {
		return 0, fmt.Errorf("%s %w", op, err)
	}

	for _, s := range similar {
		if _, err := p.cr.CreateContentSimilar(ctx, model.ContentSimilar{
			ContentId: id,
			SimilarId: s,
		}); err != nil {
			p.log.Error("Failed to create similar content", fmt.Errorf("%s %w", op, err))
		}
	}

	if err := p.CreateContentExternalId(ctx, id, data); err != nil {
		p.log.Error("Error creating external id", fmt.Errorf("%s %w", op, err))
	}

	if err := p.CreateContentCountry(ctx, id, data); err != nil {
		p.log.Error("Error creating country", fmt.Errorf("%s %w", op, err))
	}

	if err := p.CreateContentTranslation(ctx, id, data); err != nil {
		p.log.Error("Error creating translation", fmt.Errorf("%s %w", op, err))
	}

	if err := p.CreateContentGenre(ctx, id, data); err != nil {
		p.log.Error("Error creating genre", fmt.Errorf("%s %w", op, err))
	}

	if err := p.CreateContentCast(ctx, id, data); err != nil {
		p.log.Error("Error creating cast", fmt.Errorf("%s %w", op, err))
	}

	return id, nil
}

func (p *BalancerService) CreateContentExternalId(ctx context.Context, contentId int64, data _struct.CdnMoviesContent) error {
	const op = "BalancerService.CreateContentExternalId() ->"

	externalIds := []model.ContentExternalId{
		{
			ContentId:           contentId,
			ExternalId:          strconv.Itoa(data.KinopoiskId),
			ExternalType:        enum.KinopoiskExternalId,
			ExternalRating:      data.KinopoiskRating,
			ExternalRatingVotes: data.KinopoiskRatingVotes,
		},
		{
			ContentId:           contentId,
			ExternalId:          data.ImdbId,
			ExternalType:        enum.ImdbExternalId,
			ExternalRating:      data.ImdbRating,
			ExternalRatingVotes: data.ImdbRatingVotes,
		},
		{
			ContentId:           contentId,
			ExternalId:          strconv.Itoa(data.TmdbId),
			ExternalType:        enum.TmdbExternalId,
			ExternalRating:      data.TmdbRating,
			ExternalRatingVotes: data.TmdbRatingVotes,
		},
	}

	for _, externalId := range externalIds {
		if len(externalId.ExternalId) < 2 {
			continue
		}

		if _, err := p.cr.CreateContentExternalId(ctx, externalId); err != nil {
			return fmt.Errorf("%s failed to create external id: %w", op, err)
		}
	}

	return nil
}

func (p *BalancerService) CreateContentTranslation(ctx context.Context, contentId int64, data _struct.CdnMoviesContent) error {
	const op = "BalancerService.CreateContentTranslation() ->"

	for _, tr := range data.Translations {
		translation, err := p.tr.FindFirstByExternalId(ctx, tr.Id)
		if err != nil {
			translation = model.Translation{
				ExternalId:  int(tr.Id),
				Title:       tr.Name,
				FormatTitle: tr.FormatName,
			}

			if translation.Id, err = p.tr.Create(ctx, translation); err != nil {
				return fmt.Errorf("%s failed to create translation: %w", op, err)
			}
		}

		ct := model.ContentTranslation{
			ContentId:     contentId,
			TranslationId: translation.Id,
			Quality:       tr.Quality,
			MaxQuality:    tr.MaxQuality,
		}

		if _, err := p.cr.CreateContentTranslation(ctx, ct); err != nil {
			return fmt.Errorf("%s failed to create content translation: %w", op, err)
		}
	}

	return nil
}

func (p *BalancerService) CreateContentGenre(ctx context.Context, contentId int64, data _struct.CdnMoviesContent) error {
	const op = "BalancerService.CreateContentGenre() ->"

	for _, gr := range data.Genres {
		genre, err := p.gr.FindFirstByTitle(ctx, gr.RuName)
		if err != nil {
			return fmt.Errorf("%s failed to find genre: %w", op, err)
		}

		if _, err := p.cr.CreateContentGenre(ctx, model.ContentGenre{
			ContentId: contentId,
			GenreId:   genre.Id,
		}); err != nil {
			return fmt.Errorf("%s failed to create genre: %w", op, err)
		}
	}

	return nil
}

func (p *BalancerService) CreateContentCountry(ctx context.Context, contentId int64, data _struct.CdnMoviesContent) error {
	const op = "BalancerService.CreateContentCountry() ->"

	for _, c := range data.Countries {
		country, err := p.ctr.FindCountryByTitle(ctx, c.RuName)
		if err != nil {
			return fmt.Errorf("%s failed to find country: %w", op, err)
		}

		if _, err := p.cr.CreateContentCountry(ctx, model.ContentCountry{
			ContentId: contentId,
			CountryId: country.Id,
		}); err != nil {
			return fmt.Errorf("%s failed to create country: %w", op, err)
		}
	}

	return nil
}

func (p *BalancerService) CreateContentCast(ctx context.Context, contentId int64, data _struct.CdnMoviesContent) error {
	const op = "BalancerService.CreateContentCast() ->"

	if enum.CdnMoviesContentType(data.ContentType).IsMovieContentType() {
		if err := p._storeCreditsForMovie(ctx, contentId, data); err != nil {
			return fmt.Errorf("%s failed to store credits for movie: %w", op, err)
		}
	} else if enum.CdnMoviesContentType(data.ContentType).IsTvSeriesContentType() {
		if err := p._storeCreditsForTvSeries(ctx, contentId, data); err != nil {
			return fmt.Errorf("%s failed to store credits for tv series: %w", op, err)
		}
	}

	return nil
}

func (p *BalancerService) _getSimilarForMovies(ctx context.Context, data _struct.CdnMoviesContent) ([]int64, error) {
	const op = "BalancerService._GetSimilarForMovies() ->"

	similar, err := p.tmdbClient.GetMovieSimilar(data.TmdbId, nil)
	if err != nil {
		return nil, fmt.Errorf("%s failed to get similar movies: %w", op, err)
	}

	var similarMovies []int64

	for _, movie := range similar.Results {
		similarMovies = append(similarMovies, movie.ID)
	}

	return similarMovies, nil
}

func (p *BalancerService) _getSimilarForTvSeries(data _struct.CdnMoviesContent) ([]int64, error) {
	const op = "BalancerService._GetSimilarForTvSeries() ->"

	similar, err := p.tmdbClient.GetTVSimilar(data.TmdbId, nil)
	if err != nil {
		return nil, fmt.Errorf("%s failed to get similar tv series: %w", op, err)
	}

	var similarTvSeries []int64

	for _, tvSeries := range similar.Results {
		similarTvSeries = append(similarTvSeries, tvSeries.ID)
	}

	return similarTvSeries, nil
}

func (p *BalancerService) _getImagesFromTmdbForMovies(data _struct.CdnMoviesContent) (string, string, error) {
	const op = "BalancerService._GetImagesFromTmdbForMovies() ->"

	var backdrop string
	var poster string

	images, err := p.tmdbClient.GetMovieImages(data.TmdbId, nil)
	if err != nil {
		return "", "", fmt.Errorf("failed to get movie images: %w", err)
	}

	if len(images.Backdrops) > 0 {
		sort.Slice(images.Backdrops, func(i, j int) bool {
			return images.Backdrops[i].Width > images.Backdrops[j].Width
		})

		imageUrl, err := url.Parse(tmdb.GetImageURL(images.Backdrops[0].FilePath, "original"))
		if err != nil {
			return "", "", fmt.Errorf("%s failed to parse backdrop url: %w", op, err)
		}

		if backdrop, err = p.image.Download(_struct.Image{
			Url: *imageUrl,
		}); err != nil {
			return "", "", fmt.Errorf("%s failed to download backdrop: %w", op, err)
		}
	}

	if len(images.Posters) > 0 {
		sort.Slice(images.Posters, func(i, j int) bool {
			return images.Posters[i].Width > images.Posters[j].Width
		})

		imageUrl, err := url.Parse(tmdb.GetImageURL(images.Posters[0].FilePath, "original"))
		if err != nil {
			return "", "", fmt.Errorf("%s failed to parse poster url: %w", op, err)
		}

		if poster, err = p.image.Download(_struct.Image{
			Url: *imageUrl,
		}); err != nil {
			return "", "", fmt.Errorf("%s failed to download poster: %w", op, err)
		}
	}

	return poster, backdrop, nil
}

func (p *BalancerService) _getImagesFromTmdbForTvSeries(data _struct.CdnMoviesContent) (string, string, error) {
	const op = "BalancerService._GetImagesFromTmdbForTvSeries() ->"

	var backdrop string
	var poster string

	images, err := p.tmdbClient.GetTVImages(data.TmdbId, nil)
	if err != nil {
		return "", "", fmt.Errorf("%s failed to get tv series images: %w", op, err)
	}

	if len(images.Backdrops) > 0 {
		sort.Slice(images.Backdrops, func(i, j int) bool {
			return images.Backdrops[i].Width > images.Backdrops[j].Width
		})

		imageUrl, err := url.Parse(tmdb.GetImageURL(images.Backdrops[0].FilePath, "original"))
		if err != nil {
			return "", "", fmt.Errorf("%s failed to parse backdrop url: %w", op, err)
		}

		if backdrop, err = p.image.Download(_struct.Image{
			Url: *imageUrl,
		}); err != nil {
			return "", "", fmt.Errorf("%s failed to download backdrop: %w", op, err)
		}
	}

	if len(images.Posters) > 0 {
		sort.Slice(images.Posters, func(i, j int) bool {
			return images.Posters[i].Width > images.Posters[j].Width
		})

		imageUrl, err := url.Parse(tmdb.GetImageURL(images.Posters[0].FilePath, "original"))
		if err != nil {
			return "", "", fmt.Errorf("%s failed to parse poster url: %w", op, err)
		}

		if poster, err = p.image.Download(_struct.Image{
			Url: *imageUrl,
		}); err != nil {
			return "", "", fmt.Errorf("%s failed to download poster: %w", op, err)
		}
	}

	return poster, backdrop, nil
}

func (p *BalancerService) _storeCreditsForMovie(ctx context.Context, contentId int64, data _struct.CdnMoviesContent) error {
	const op = "BalancerService._GetCastsFromMovie() ->"

	credits, err := p.tmdbClient.GetMovieCredits(data.TmdbId, nil)
	if err != nil {
		return fmt.Errorf("%s failed to get movie credits: %w", op, err)
	}

	var wg sync.WaitGroup

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		for _, cast := range credits.Cast {
			credit := model.Credit{
				ExternalId: cast.ID,
				Name:       cast.Name,
				OrigName:   cast.OriginalName,
				Popularity: cast.Popularity,
			}

			if cast.ProfilePath != "" {
				imageUrl, err := url.Parse(tmdb.GetImageURL(cast.ProfilePath, "original"))
				if err != nil {
					continue
				}

				h := sha256.New()
				h.Write([]byte(filepath.Base(imageUrl.String())))

				path, err := p.image.Download(_struct.Image{
					Url:      *imageUrl,
					Filename: fmt.Sprintf("%x.%s", h.Sum(nil), "webp"),
				})
				if err != nil {
					return
				}

				credit.Image = path
			}

			creditId, err := p._storeCreditIfNotExists(ctx, credit)
			if err != nil {
				return
			}

			if _, err := p.cr.CreateContentCast(ctx, model.ContentCast{
				ExternalId: cast.CreditID,
				ContentId:  contentId,
				CreditId:   creditId,
				Department: cast.KnownForDepartment,
				Character:  cast.Character,
				Sort:       cast.Order,
			}); err != nil {
				continue
			}
		}
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		for _, crew := range credits.Crew {
			credit := model.Credit{
				ExternalId: crew.ID,
				Name:       crew.Name,
				OrigName:   crew.OriginalName,
				Popularity: crew.Popularity,
			}

			if crew.ProfilePath != "" {
				imageUrl, err := url.Parse(tmdb.GetImageURL(crew.ProfilePath, "original"))
				if err != nil {
					continue
				}

				h := sha256.New()
				h.Write([]byte(filepath.Base(imageUrl.String())))

				path, err := p.image.Download(_struct.Image{
					Url:      *imageUrl,
					Filename: fmt.Sprintf("%x.%s", h.Sum(nil), "webp"),
				})
				if err != nil {
					return
				}

				credit.Image = path
			}

			creditId, err := p._storeCreditIfNotExists(ctx, credit)
			if err != nil {
				return
			}

			if _, err := p.cr.CreateContentCrew(ctx, model.ContentCrew{
				ExternalId: crew.CreditID,
				ContentId:  contentId,
				CreditId:   creditId,
				Department: crew.Department,
				Job:        crew.Job,
			}); err != nil {
				continue
			}

		}
		wg.Done()
	}(&wg)

	wg.Wait()

	return nil
}

func (p *BalancerService) _storeCreditsForTvSeries(ctx context.Context, contentId int64, data _struct.CdnMoviesContent) error {
	const op = "BalancerService._GetCastsFromMovie() ->"

	credits, err := p.tmdbClient.GetTVCredits(data.TmdbId, nil)
	if err != nil {
		return fmt.Errorf("%s failed to get tv series credits: %w", op, err)
	}

	var wg sync.WaitGroup

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		for _, cast := range credits.Cast {
			credit := model.Credit{
				ExternalId: cast.ID,
				Name:       cast.Name,
				OrigName:   cast.OriginalName,
				Popularity: cast.Popularity,
			}

			if cast.ProfilePath != "" {
				imageUrl, err := url.Parse(tmdb.GetImageURL(cast.ProfilePath, "original"))
				if err != nil {
					continue
				}

				h := sha256.New()
				h.Write([]byte(filepath.Base(imageUrl.String())))

				path, err := p.image.Download(_struct.Image{
					Url:      *imageUrl,
					Filename: fmt.Sprintf("%x.%s", h.Sum(nil), "webp"),
				})
				if err != nil {
					return
				}

				credit.Image = path
			}

			creditId, err := p._storeCreditIfNotExists(ctx, credit)
			if err != nil {
				return
			}

			if _, err := p.cr.CreateContentCast(ctx, model.ContentCast{
				ExternalId: cast.CreditID,
				ContentId:  contentId,
				CreditId:   creditId,
				Department: cast.KnownForDepartment,
				Character:  cast.Character,
				Sort:       cast.Order,
			}); err != nil {
				continue
			}
		}
		wg.Done()
	}(&wg)

	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		for _, crew := range credits.Crew {
			credit := model.Credit{
				ExternalId: crew.ID,
				Name:       crew.Name,
				OrigName:   crew.OriginalName,
				Popularity: crew.Popularity,
			}

			if crew.ProfilePath != "" {
				imageUrl, err := url.Parse(tmdb.GetImageURL(crew.ProfilePath, "original"))
				if err != nil {
					continue
				}

				h := sha256.New()
				h.Write([]byte(filepath.Base(imageUrl.String())))

				path, err := p.image.Download(_struct.Image{
					Url:      *imageUrl,
					Filename: fmt.Sprintf("%x.%s", h.Sum(nil), "webp"),
				})
				if err != nil {
					return
				}

				credit.Image = path
			}

			creditId, err := p._storeCreditIfNotExists(ctx, credit)
			if err != nil {
				return
			}

			if _, err := p.cr.CreateContentCrew(ctx, model.ContentCrew{
				ExternalId: crew.CreditID,
				ContentId:  contentId,
				CreditId:   creditId,
				Department: crew.Department,
				Job:        crew.Job,
			}); err != nil {
				continue
			}

		}
		wg.Done()
	}(&wg)

	wg.Wait()

	return nil
}

func (p *BalancerService) _storeCreditIfNotExists(ctx context.Context, credit model.Credit) (int64, error) {
	const op = "BalancerService._storeCreditIfNotExists() ->"

	var person model.Credit
	person = credit

	exist, err := p.csr.GetCreditByTmdbId(ctx, credit.Id)
	if err != nil {
		//log.Error("Failed to get cast", slog.Any("err", err))
	}

	if exist.Id == 0 {
		if person.Id, err = p.csr.Create(ctx, person); err != nil {
			return 0, fmt.Errorf("%s failed to create cast: %w", op, err)
		}
	} else {
		person.Id = exist.Id
	}

	return person.Id, nil
}

func (p *BalancerService) Get(url string) (ResponseData, error) {
	const op = "BalancerService.Get() ->"

	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ResponseData{}, fmt.Errorf("%s failed to create request: %w", op, err)
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return ResponseData{}, fmt.Errorf("%s failed to send request: %w", op, err)
	}

	defer response.Body.Close()

	var responseData ResponseData
	if err := render.DecodeJSON(response.Body, &responseData); err != nil {
		return ResponseData{}, fmt.Errorf("%s failed to decode json: %w", op, err)
	}

	return responseData, nil
}

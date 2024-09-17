package enum

import "slices"

type ContentType int

const (
	Movie ContentType = 1 + iota
	TvSeries
	Anime
	AnimeTvSeries
	TvShow
)

func CdnMoviesContentType(s string) ContentType {
	switch s {
	case "фильм":
		return Movie
	case "сериал":
		return TvSeries
	case "аниме":
		return Anime
	case "аниме сериал":
		return AnimeTvSeries
	case "тв телепередача":
		return TvShow
	default:
		return -1
	}
}

func (c ContentType) IsMovieContentType() bool {
	return slices.Contains([]ContentType{Movie, Anime}, c)
}

func (c ContentType) IsTvSeriesContentType() bool {
	return slices.Contains([]ContentType{TvSeries, AnimeTvSeries, TvShow}, c)
}

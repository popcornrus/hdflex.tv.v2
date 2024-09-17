package _struct

type CdnMoviesContent struct {
	Id                   string                  `json:"id"`
	IframeSrc            string                  `json:"iframe_src"`
	ContentType          string                  `json:"content_type"`
	RuTitle              string                  `json:"ru_title"`
	OrigTitle            string                  `json:"orig_title"`
	EnTitle              string                  `json:"en_title"`
	Year                 int                     `json:"year"`
	RatingAge            int                     `json:"rating_age"`
	RatingMpaa           string                  `json:"rating_mpaa"`
	KinopoiskId          int                     `json:"kinopoisk_id"`
	KinopoiskRating      float64                 `json:"kinopoisk_rating"`
	KinopoiskRatingVotes int                     `json:"kinopoisk_rating_votes"`
	ImdbId               string                  `json:"imdb_id"`
	ImdbRating           float64                 `json:"imdb_rating"`
	ImdbRatingVotes      int                     `json:"imdb_rating_votes"`
	TmdbId               int                     `json:"tmdb_id"`
	TmdbRating           float64                 `json:"tmdb_rating"`
	TmdbRatingVotes      int                     `json:"tmdb_rating_votes"`
	Slogan               string                  `json:"slogan"`
	Duration             int                     `json:"duration"`
	Description          string                  `json:"description"`
	FeelsWorld           string                  `json:"feels_world"`
	FeelsRu              string                  `json:"feels_ru"`
	FeelsUs              string                  `json:"feels_us"`
	Budget               float64                 `json:"budget"`
	WorldPremiere        string                  `json:"premiere_world"`
	RuPremiere           string                  `json:"premiere_russia"`
	Countries            []CdnMoviesCountry      `json:"countries"`
	Genres               []CdnMoviesGenre        `json:"genres"`
	Translations         []CdnMoviesTranslations `json:"translations"`
	Facts                []string                `json:"facts"`
	LastSeason           int                     `json:"last_season"`
	LastEpisode          int                     `json:"last_episode"`
	Lgbt                 int                     `json:"lgbt"`
}

type CdnMoviesCountry struct {
	RuName string `json:"ru_name"`
	EnName string `json:"en_name"`
}

type CdnMoviesGenre struct {
	RuName string `json:"ru_name"`
	EnName string `json:"en_name"`
}

type CdnMoviesTranslations struct {
	Id         int64  `json:"id"`
	Name       string `json:"name"`
	FormatName string `json:"format_name"`
	Quality    string `json:"quality"`
	MaxQuality int    `json:"max_quality"`
}

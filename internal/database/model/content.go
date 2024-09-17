package model

import (
	"go-hdflex/internal/database/enum"
	"time"
)

type Content struct {
	Id            int64            `json:"id"`
	CdnMoviesId   string           `json:"-" db:"cdnmovies_id"`
	RuTitle       string           `json:"ru_title" db:"ru_title"`
	OrigTitle     string           `json:"orig_title" db:"orig_title"`
	EnTitle       string           `json:"en_title" db:"en_title"`
	Url           string           `json:"url" db:"url"`
	Popularity    float64          `json:"popularity" db:"popularity"`
	Slogan        string           `json:"slogan" db:"slogan"`
	Description   string           `json:"description" db:"description"`
	Duration      int              `json:"duration" db:"duration"`
	IframeUrl     string           `json:"iframe_url" db:"iframe_url"`
	ContentType   enum.ContentType `json:"content_type" db:"content_type"`
	Year          int              `json:"year" db:"year"`
	Poster        string           `json:"poster" db:"poster"`
	Backdrop      string           `json:"backdrop" db:"backdrop"`
	RatingAge     int              `json:"rating_age" db:"rating_age"`
	RatingMpaa    string           `json:"rating_mpaa" db:"rating_mpaa"`
	WorldPremiere *time.Time       `json:"world_premiere" db:"world_premiere"`
	RuPremiere    *time.Time       `json:"ru_premiere" db:"ru_premiere"`
	LastSeason    int              `json:"last_season" db:"last_season"`
	LastEpisode   int              `json:"last_episode" db:"last_episode"`
	Lgbt          bool             `json:"lgbt" db:"lgbt"`
	LastWatchAt   *time.Time       `json:"last_watch_at" db:"last_watch_at"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`
}

type ContentExternalId struct {
	Id                  int64               `json:"id"`
	ContentId           int64               `json:"content_id" db:"content_id"`
	ExternalId          string              `json:"external_id" db:"external_id"`
	ExternalType        enum.ExternalIdType `json:"external_type" db:"external_type"`
	ExternalRating      float64             `json:"external_rating" db:"external_rating"`
	ExternalRatingVotes int                 `json:"external_rating_votes" db:"external_rating_votes"`
}

type ContentTranslation struct {
	Id            int64  `json:"id"`
	ContentId     int64  `json:"content_id" db:"content_id"`
	TranslationId int64  `json:"translation_id" db:"translation_id"`
	Quality       string `json:"quality" db:"quality"`
	MaxQuality    int    `json:"max_quality" db:"max_quality"`
}

type ContentGenre struct {
	Id        int64 `json:"id"`
	ContentId int64 `json:"content_id" db:"content_id"`
	GenreId   int   `json:"genre_id" db:"genre_id"`
}

type ContentCountry struct {
	Id        int64 `json:"id"`
	ContentId int64 `json:"content_id" db:"content_id"`
	CountryId int   `json:"country_id" db:"country_id"`
}

type ContentCast struct {
	Id         int64  `json:"id"`
	ExternalId string `json:"external_id" db:"external_id"`
	ContentId  int64  `json:"content_id" db:"content_id"`
	CreditId   int64  `json:"credit_id" db:"credit_id"`
	Department string `json:"department" db:"department"`
	Character  string `json:"character" db:"character"`
	Sort       int    `json:"sort" db:"sort"`
}

type ContentCrew struct {
	Id         int64  `json:"id"`
	ExternalId string `json:"external_id" db:"external_id"`
	ContentId  int64  `json:"content_id" db:"content_id"`
	CreditId   int64  `json:"credit_id" db:"credit_id"`
	Department string `json:"department" db:"department"`
	Job        string `json:"job" db:"job"`
}

type ContentSimilar struct {
	Id        int64 `json:"id"`
	ContentId int64 `json:"content_id" db:"content_id"`
	SimilarId int64 `json:"similar_id" db:"similar_id"`
}

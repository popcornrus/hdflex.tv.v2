package _struct

import (
	"go-hdflex/internal/database/enum"
	"go-hdflex/internal/database/model"
	"time"
)

type GetContent struct {
	Id          int64            `json:"id"`
	RuTitle     string           `json:"ru_title" db:"ru_title"`
	OrigTitle   string           `json:"orig_title" db:"orig_title"`
	EnTitle     string           `json:"en_title" db:"en_title"`
	Url         string           `json:"url" db:"url"`
	ContentType enum.ContentType `json:"content_type" db:"content_type"`
	Year        int              `json:"year" db:"year"`
	Poster      string           `json:"poster" db:"poster"`
	RatingAge   int              `json:"rating_age" db:"rating_age"`
	LastSeason  int              `json:"last_season" db:"last_season"`
	LastEpisode int              `json:"last_episode" db:"last_episode"`
	Lgbt        bool             `json:"lgbt" db:"lgbt"`
	CreatedAt   time.Time        `json:"created_at"`
	UpdatedAt   time.Time        `json:"updated_at"`

	Genres      []Genre      `json:"genres" db:"-"`
	Countries   []Country    `json:"countries" db:"-"`
	ExternalIds []ExternalId `json:"external_ids" db:"-"`
}

type ShowContent struct {
	Id            int64            `json:"id"`
	RuTitle       string           `json:"ru_title" db:"ru_title"`
	OrigTitle     string           `json:"orig_title" db:"orig_title"`
	EnTitle       string           `json:"en_title" db:"en_title"`
	Slogan        string           `json:"slogan" db:"slogan"`
	Description   string           `json:"description" db:"description"`
	IframeUrl     string           `json:"iframe_url" db:"iframe_url"`
	ContentType   enum.ContentType `json:"content_type" db:"content_type"`
	Year          int              `json:"year" db:"year"`
	Poster        string           `json:"poster" db:"poster"`
	Backdrop      string           `json:"backdrop" db:"backdrop"`
	Duration      int              `json:"duration" db:"duration"`
	RatingAge     int              `json:"rating_age" db:"rating_age"`
	RatingMpaa    string           `json:"rating_mpaa" db:"rating_mpaa"`
	LastSeason    int              `json:"last_season" db:"last_season"`
	LastEpisode   int              `json:"last_episode" db:"last_episode"`
	Lgbt          bool             `json:"lgbt" db:"lgbt"`
	WorldPremiere time.Time        `json:"world_premiere" db:"world_premiere"`
	CreatedAt     time.Time        `json:"created_at"`
	UpdatedAt     time.Time        `json:"updated_at"`

	Genres      []Genre      `json:"genres" db:"-"`
	Countries   []Country    `json:"countries" db:"-"`
	ExternalIds []ExternalId `json:"external_ids" db:"-"`
	Casts       []Cast       `json:"casts" db:"-"`
	Crew        []Crew       `json:"crew" db:"-"`
}

type Country struct {
	Id      int    `json:"id"`
	Title   string `json:"title"`
	EnTitle string `json:"en_title"`
}

type ExternalId struct {
	Id             int                 `json:"id"`
	ExternalId     string              `json:"external_id"`
	ExternalTypeId enum.ExternalIdType `json:"-"`
	ExternalType   string              `json:"external_type"`
	Rating         float64             `json:"rating"`
	Votes          int                 `json:"votes"`
}

type Cast struct {
	Id         int          `json:"id"`
	Person     model.Credit `json:"person"`
	Department string       `json:"department"`
	Character  string       `json:"character"`
}

type Crew struct {
	Id         int          `json:"id"`
	Person     model.Credit `json:"person"`
	Department string       `json:"department"`
	Job        string       `json:"job"`
}

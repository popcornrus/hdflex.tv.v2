package model

type Genre struct {
	Id          int    `json:"id"`
	CdnMoviesId int    `json:"-"`
	Title       string `json:"title"`
	EnTitle     string `json:"en_title"`
}

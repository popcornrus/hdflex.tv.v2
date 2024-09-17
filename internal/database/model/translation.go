package model

type Translation struct {
	Id          int64  `json:"id" db:"id" goqu:"skipinsert"`
	ExternalId  int    `json:"external_id" db:"cdnmovies_id"`
	Title       string `json:"title" db:"title"`
	FormatTitle string `json:"format_title" db:"format_title"`
}

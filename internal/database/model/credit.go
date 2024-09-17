package model

type Credit struct {
	Id         int64   `json:"id,omitempty"`
	ExternalId int64   `json:"external_id,omitempty"`
	Name       string  `json:"name"`
	OrigName   string  `json:"orig_name"`
	Popularity float32 `json:"popularity"`
	Image      string  `json:"image"`
}

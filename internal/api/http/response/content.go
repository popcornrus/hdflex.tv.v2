package response

import (
	"go-hdflex/internal/database/struct"
)

type ContentGetResponse struct {
	Items []_struct.GetContent `json:"items"`
}

type ContentDetailResponse struct {
	Item _struct.ShowContent `json:"item"`
}

type ContentCountResponse struct {
	Count int `json:"count"`
}

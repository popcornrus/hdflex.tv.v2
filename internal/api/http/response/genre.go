package response

import (
	"go-hdflex/internal/database/model"
)

type GenreGetResponse struct {
	Items []model.Genre `json:"items"`
}

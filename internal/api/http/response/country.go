package response

import (
	"go-hdflex/internal/database/model"
)

type CountryGetResponse struct {
	Items []model.Country `json:"items"`
}

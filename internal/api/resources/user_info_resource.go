package resources

import (
	"encoding/json"
	"go-boilerplate/internal/api/models"
)

type UserResources struct {
	User models.User
}

func (u UserResources) UserInfo() ([]byte, error) {
	var resource struct {
		ID        *uint64 `json:"id"`
		UUID      string  `json:"uuid"`
		Name      *string `json:"name"`
		AvatarURL *string `json:"avatar_url"`
	}

	resource.ID = u.User.ID
	resource.UUID = u.User.UUID
	resource.Name = u.User.Name
	resource.AvatarURL = u.User.AvatarURL

	return json.Marshal(&resource)
}

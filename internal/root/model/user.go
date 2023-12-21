package model

import "time"

type User struct {
	ID        int64     `json:"-"`
	UUID      string    `json:"uuid"`
	Username  string    `json:"username"`
	Password  string    `json:"-"`
	Email     string    `json:"email"`
	AvatarURL string    `json:"avatar_url"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"updated_at"`

	Token string `json:"-"`
}

package types

type User struct {
	ID        int64  `json:"id"`
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt int64  `json:"created_at"`
	UpdatedAt int64  `json:"updated_at"`
}

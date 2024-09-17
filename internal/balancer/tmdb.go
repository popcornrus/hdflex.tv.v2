package balancer

import (
	tmdb "github.com/cyruzin/golang-tmdb"
	"go-hdflex/external/config"
)

func NewTmdbClient(
	cfg *config.Config,
) *tmdb.Client {
	client, err := tmdb.Init(cfg.Tmdb.ApiKey)
	if err != nil {
		panic(err)
	}

	return client
}

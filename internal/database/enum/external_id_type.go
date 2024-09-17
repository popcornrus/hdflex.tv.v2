package enum

type ExternalIdType int

const (
	ImdbExternalId ExternalIdType = iota
	KinopoiskExternalId
	TmdbExternalId
)

func (e ExternalIdType) String() string {
	return [...]string{"imdb", "kinopoisk", "tmdb"}[e]
}

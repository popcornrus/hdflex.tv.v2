package mongodb

import "go.uber.org/fx"

func NewMongoDatabaseModule() fx.Option {
	return fx.Module(
		"mongodb",
		fx.Provide(
			NewMongoDatabase,
			fx.Annotate(
				NewMongo,
				fx.As(new(DatabaseInterface)),
			),
		),
	)
}

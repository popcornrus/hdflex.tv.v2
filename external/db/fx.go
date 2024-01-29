package db

import "go.uber.org/fx"

func NewDataBase() fx.Option {
	return fx.Module(
		"database",
		fx.Provide(
			NewMongoDatabase,
			fx.Annotate(
				NewMongo,
				fx.As(new(MongoDBInterface)),
			),
		),
	)
}

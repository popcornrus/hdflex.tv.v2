package mysql

import "go.uber.org/fx"

func NewMysqlDatabaseModule() fx.Option {
	return fx.Module(
		"mysql",
		fx.Provide(
			NewMysqlDatabase,
			fx.Annotate(
				NewMysql,
				fx.As(new(DatabaseInterface)),
			),
		),
	)
}

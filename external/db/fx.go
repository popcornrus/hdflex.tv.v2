package db

import (
	"go-hdflex/external/db/mongodb"
	"go-hdflex/external/db/mysql"
	"go.uber.org/fx"
)

func NewDataBase() fx.Option {
	return fx.Module(
		"database",
		fx.Options(
			mysql.NewMysqlDatabaseModule(),
			mongodb.NewMongoDatabaseModule(),
		),
	)
}

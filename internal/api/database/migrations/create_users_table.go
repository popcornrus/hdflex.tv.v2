package migrations

import (
	"go-boilerplate/internal/api/models"
	"log"
)

func CreateUsersTable() {
	var (
		table models.User
	)

	if !MySQL.Migrator().HasTable(table) {
		err = MySQL.AutoMigrate(table)
		if err != nil {
			log.Fatalln(err)
			return
		}
	}
}

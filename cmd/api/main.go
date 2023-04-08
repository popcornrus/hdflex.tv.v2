package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go-boilerplate/internal/api/database/migrations"
	"go-boilerplate/internal/api/routes"
	"log"
	"os"
)

func main() {
	var err error

	if err = godotenv.Load(".env"); err != nil {
		log.Fatalln(err)
		return
	}

	r := gin.Default()

	routes.RouteHandle(r)

	go migrations.Migrations{}.MigrateAll()

	if err = r.Run(":" + os.Getenv("GO_PORT")); err != nil {
		log.Fatalln(err)
		return
	}
}

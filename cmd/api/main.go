package main

import (
	"go-boilerplate/internal/api"
)

func main() {
	fx := api.NewApp()
	fx.Run()
}

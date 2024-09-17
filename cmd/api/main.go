package main

import (
	"go-hdflex/internal/api"
)

func main() {
	fx := api.NewApp()
	fx.Run()
}

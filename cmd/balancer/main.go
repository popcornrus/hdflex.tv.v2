package main

import (
	"go-hdflex/internal/balancer"
)

func main() {
	fx := balancer.NewApp()
	fx.Run()
}

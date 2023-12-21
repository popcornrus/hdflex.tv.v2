package main

import "go-boilerplate/internal/root"

func main() {
	fx := root.NewApp()
	fx.Run()
}

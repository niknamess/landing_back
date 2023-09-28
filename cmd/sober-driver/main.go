package main

import (
	"log"
	"sober_driver/internal/app"
)

// swag init --parseDependency --parseInternal --parseDepth 1 -g .\cmd\sober-driver\main.go --output .\cmd\sober-driver\docs
func main() {
	log.Println("start")
	app.App()
}

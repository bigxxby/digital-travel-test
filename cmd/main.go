package main

import (
	"log"

	"github.com/bigxxby/digital-travel-test/internal/app"
	"github.com/bigxxby/digital-travel-test/internal/config"
)

// set config first
func init() {
	err := config.SetConfig()
	if err != nil {
		log.Println(err)
		return
	}
}

// @title						basic CRUD API
// @version					1.0
// @description				This is a basic CRUD API for managing users orders and products
// @host						localhost:8081
// @BasePath					/
// @securityDefinitions.apikey	BearerAuth
// @in							header
// @name						Authorization
// @description				Provide your Bearer token in the format: Bearer <token>
func main() {
	app.App()
}

package main

import (
	"log"

	"github.com/lucamienert/Social-Media/config"
)

// @title API
// @version 1.0
// @description Backend
// @host localhost:8000
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	cfg, err := config.LoadConfig(".")

	if err != nil {
		log.Fatal("Could not load environment variables", err)
	}

	config.InitDB(&cfg)

}

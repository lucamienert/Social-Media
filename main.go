package main

import (
	"lucamienert/twitter-clone/config"
	"lucamienert/twitter-clone/routes"
)

// @title Twitter Clone API
// @version 1.0
// @description A simple Twitter clone backend in Go
// @host localhost:8080
// @BasePath /
func main() {
	config.InitDB()
	r := routes.SetupRouter()
	r.Run(":8080")
}

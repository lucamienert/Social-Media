package main

import (
	"lucamienert/twitter-clone/config"
	"lucamienert/twitter-clone/controllers"
	"lucamienert/twitter-clone/repository"
	"lucamienert/twitter-clone/routes"
	"lucamienert/twitter-clone/services"

	"github.com/gin-gonic/gin"
)

// @title Twitter Clone API
// @version 1.0
// @description A simple Twitter clone backend in Go
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	config.InitDB()

	userRepo := repository.NewUserRepository(config.DB)
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	postRepo := repository.NewPostRepository()
	postService := services.NewPostService(postRepo)
	postController := controllers.NewPostController(postService)

	r := gin.Default()

	routes.SetupUserRoutes(r, userController)
	routes.SetupPostRoutes(r, postController)
	routes.SetupRouter(r)

	r.Run(":8080")
}

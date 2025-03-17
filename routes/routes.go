package routes

import (
	"lucamienert/twitter-clone/controllers"
	"lucamienert/twitter-clone/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Auth routes
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	// Protected Routes
	auth := r.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/post", controllers.CreatePost)
		auth.GET("/posts", controllers.GetPosts)
		auth.POST("/comment", controllers.CreateComment)
	}

	return r
}

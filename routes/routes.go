package routes

import (
	"lucamienert/twitter-clone/controllers"
	"time"

	_ "lucamienert/twitter-clone/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupUserRoutes(router *gin.Engine, userController *controllers.UserController) {
	userRoutes := router.Group("/users")
	{
		userRoutes.POST("/register", userController.Register)
		userRoutes.POST("/login", userController.Login)
		userRoutes.DELETE("/:email", userController.DeleteUser)
	}
}

func SetupPostRoutes(router *gin.Engine, postController *controllers.PostController, commentController *controllers.CommentController) {
	postRoutes := router.Group("/")
	{
		postRoutes.POST("/post", postController.CreatePost)
		postRoutes.GET("/posts", postController.GetPosts)
		postRoutes.POST("/post/:post_id/like", postController.LikePost)

		postRoutes.POST("/post/:post_id/comment", commentController.AddComment)
		router.GET("/post/:post_id/comments", commentController.GetComments)

	}
}

func SetupRouter(router *gin.Engine) {
	// Enable Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Enable CORS
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Allow all domains (Change this to restrict access)
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lucamienert/Social-Media/config"
	"github.com/lucamienert/Social-Media/controllers"
	_ "github.com/lucamienert/Social-Media/docs"
	"github.com/lucamienert/Social-Media/routes"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	AuthController      controllers.AuthController
	AuthRouteController routes.AuthRouteController

	UserController      controllers.UserController
	UserRouteController routes.UserRouteController
)

// @title API
// @version 1.0
// @description Backend
// @host localhost:8080
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

	AuthController = controllers.NewAuthController(config.DB)
	AuthRouteController = routes.NewAuthRouteController(AuthController)

	UserController = controllers.NewUserController(config.DB)
	UserRouteController = routes.NewRouteUserController(UserController)

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router := r.Group("/api")

	r.OPTIONS("/*path", func(c *gin.Context) {
		c.AbortWithStatus(204)
	})

	router.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	AuthRouteController.AuthRoute(router)
	UserRouteController.UserRoute(router)

	log.Fatal(r.Run(":" + cfg.PORT))
}

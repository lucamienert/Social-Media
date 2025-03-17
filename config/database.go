package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"lucamienert/twitter-clone/models"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("twitter_clone.db"), &gorm.Config{})
	
	if err != nil {
		log.Fatal("Failed to connect to database")
	}

	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{})
}

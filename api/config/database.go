package config

import (
	"fmt"
	"log"

	"github.com/lucamienert/Social-Media/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(cfg *Config) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.DB_HOST,
		cfg.DB_USER,
		cfg.DB_PASSWORD,
		cfg.DB_NAME,
		cfg.DB_PORT,
	)

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = DB.AutoMigrate(
		&models.SignInInput{},
		&models.SignUpInput{},
		&models.User{},
	)

	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}
}

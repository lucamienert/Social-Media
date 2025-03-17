package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username  string     `json:"username" gorm:"unique"`
	FullName  string     `json:"full_name"`
	Email     string     `json:"email" gorm:"unique"`
	Password  string     `json:"-"`
	Birthday  *time.Time `json:"birthday,omitempty"`
	Address   string     `json:"address,omitempty"`
	Phone     string     `json:"phone,omitempty"`
	Gender    string     `json:"gender,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}

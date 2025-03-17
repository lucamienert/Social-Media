package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	Text   string `json:"text"`
	Likes  int    `json:"likes"`
}

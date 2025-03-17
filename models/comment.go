package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	UserID uint   `json:"user_id"`
	PostID uint   `json:"post_id"`
	Text   string `json:"text"`
}

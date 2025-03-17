package models

import (
	"time"
)

type Post struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Content   string    `json:"content"`
	Likes     int       `json:"likes" gorm:"default:0"`
	Comments  []Comment `json:"comments" gorm:"foreignKey:PostID"`
	CreatedAt time.Time `json:"created_at"`
}

type Like struct {
	ID     uint `json:"id" gorm:"primaryKey"`
	PostID uint `json:"post_id"`
}

type Comment struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	PostID    uint      `json:"post_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
}

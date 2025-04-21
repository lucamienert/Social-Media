package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primary_key"`
	Name        string    `gorm:"type:varchar(255);not null"`
	Email       string    `gorm:"uniqueIndex;not null"`
	Password    string    `gorm:"not null"`
	Role        string    `gorm:"type:varchar(255);not null"`
	Provider    string    `gorm:"not null"`
	Verified    bool      `gorm:"not null"`
	DisplayName string    `json:"displayName"`
	Bio         *string   `json:"bio,omitempty"`

	Posts     []Post     `gorm:"foreignKey:UserID" json:"posts,omitempty"`
	Following []Follow   `gorm:"foreignKey:FollowerID" json:"following,omitempty"`
	Followers []Follow   `gorm:"foreignKey:FollowingID" json:"followers,omitempty"`
	Likes     []Like     `gorm:"foreignKey:UserID" json:"likes,omitempty"`
	Bookmarks []Bookmark `gorm:"foreignKey:UserID" json:"bookmarks,omitempty"`
	Comments  []Comment  `gorm:"foreignKey:UserID" json:"comments,omitempty"`

	CreatedAt time.Time `gorm:"not null"`
	UpdatedAt time.Time `gorm:"not null"`
}

type Follow struct {
	FollowerID  string `json:"followerId"`
	Follower    User   `gorm:"foreignKey:FollowerID;constraint:OnDelete:CASCADE;" json:"-"`
	FollowingID string `json:"followingId"`
	Following   User   `gorm:"foreignKey:FollowingID;constraint:OnDelete:CASCADE;" json:"-"`
}

type Post struct {
	ID        string     `gorm:"primaryKey" json:"id"`
	Content   string     `json:"content"`
	UserID    string     `json:"userId"`
	User      User       `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	Likes     []Like     `gorm:"foreignKey:PostID" json:"likes,omitempty"`
	Bookmarks []Bookmark `gorm:"foreignKey:PostID" json:"bookmarks,omitempty"`
	Comments  []Comment  `gorm:"foreignKey:PostID" json:"comments,omitempty"`

	CreatedAt time.Time `json:"createdAt"`
}

type Comment struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	Content   string    `json:"content"`
	UserID    string    `json:"userId"`
	User      User      `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	PostID    string    `json:"postId"`
	Post      Post      `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

type Like struct {
	UserID string `json:"userId"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE;" json:"-"`
	PostID string `json:"postId"`
	Post   Post   `gorm:"foreignKey:PostID;constraint:OnDelete:CASCADE;" json:"-"`
}

type Bookmark struct {
	ID        string    `gorm:"primaryKey" json:"id"`
	UserID    string    `json:"userId"`
	User      User      `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	PostID    string    `json:"postId"`
	Post      Post      `gorm:"constraint:OnDelete:CASCADE;" json:"-"`
	CreatedAt time.Time `json:"createdAt"`
}

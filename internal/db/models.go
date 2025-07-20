package db

import (
	"time"
	"gorm.io/gorm"
	"gorm.io/datatypes"
)

type User struct { 
	ID uint `gorm:"primaryKey"`
	Username string `gorm:"uniqueIndex;not null"`
	Email string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role string `gorm:"default:user"` 
	CreatedAt time.Time
}

type Video struct {
	ID uint `gorm:"primaryKey"`
	Title string `gorm:"not null"`
	Description string
	URL string `gorm:"not null"`
	Views int `gorm:"default:0"`
	Metadata datatypes.JSON
	UserID uint `gorm:"not null"`
	Uploader User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
}

type Comment struct { 
	ID uint `gorm:"primaryKey"`
	VideoID uint `gorm:"not null"`
	UserID uint `gorm:"not null"`
	Text string `gorm:"not null"`
	Video Video `gorm:"foreignKey:VideoID"`
	User User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time 
}

type Like struct {
	ID uint `gorm:"primaryKey"`
	UserID uint `gorm:"not null"`
	VideoID uint `gorm:"not null"`
	CreatedAt time.Time
}

type WatchHistory struct {
	ID uint `gorm:"primaryKey"`
	UserID uint
	VideoID uint
	Progress float64
	WatchedAt time.Time
}

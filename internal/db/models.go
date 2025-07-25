package db

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type User struct { 
	ID       string `gorm:"type:uuid;primaryKey;"`
	Username string `gorm:"uniqueIndex;not null"`
	Email string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
	Role string `gorm:"default:user"` 
	CreatedAt time.Time
}

type Video struct {
	ID       string `gorm:"type:uuid;primaryKey;"`
	Title         string `gorm:"not null"`
	Description   *string
	URL           string `gorm:"not null"`
	Views         int `gorm:"default:0"`
	Metadata      datatypes.JSON
	UserID        string `gorm:"not null"`
	Uploader User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
}

type Comment struct { 
	ID       string `gorm:"type:uuid;primaryKey;"`
	VideoID   string `gorm:"not null"`
	UserID    string `gorm:"not null"`
	Text      string `gorm:"not null"`
	Video     Video `gorm:"foreignKey:VideoID"`
	User      User `gorm:"foreignKey:UserID"`
	CreatedAt time.Time 
}

type Like struct {
	ID       string `gorm:"type:uuid;primaryKey;"`
	UserID uint `gorm:"not null"`
	VideoID uint `gorm:"not null"`
	CreatedAt time.Time
}

type WatchHistory struct {
	ID       string `gorm:"type:uuid;primaryKey;"`
	UserID uint
	VideoID uint
	Progress float64
	WatchedAt time.Time
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
		u.ID = uuid.New().String()
		return
}

func (u *Video) BeforeCreate(tx *gorm.DB) (err error) {
		u.ID = uuid.New().String()
		return
}

func (u *Comment) BeforeCreate(tx *gorm.DB) (err error) {
		u.ID = uuid.New().String()
		return
}

func (u *Like) BeforeCreate(tx *gorm.DB) (err error) {
		u.ID = uuid.New().String()
		return
}

func (u *WatchHistory) BeforeCreate(tx *gorm.DB) (err error) {
		u.ID = uuid.New().String()
		return
}


package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	Id             int64  `gorm:"primaryKey" json:"id"`
	Title          string `gorm:"varchar(300)" json:"title"`
	Description    string `gorm:"text" json:"description"`
	Author         string `gorm:"varchar(300)" json:"author"`
	PublishDate    string `gorm:"date" json:"publish_date"`
	ImageCoverPath string `gorm:"varchar(300)" json:"image_cover_path"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func (b *Book) BeforeSave(tx *gorm.DB) {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}
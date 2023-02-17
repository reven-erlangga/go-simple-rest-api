package models

type Book struct {
	Id          int64  `gorm:"primaryKey" json:"id"`
	Title       string `gorm:"varchar(300)" json:"title"`
	Description string `gorm:"text" json:"description"`
	Author      string `gorm:"varchar(300)" json:"author"`
	PublishDate string `gorm:"date" json:"publish_date"`
}
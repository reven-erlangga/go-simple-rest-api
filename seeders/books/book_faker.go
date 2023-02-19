package seederBook

import (
	"github.com/bxcodec/faker/v3"
	"github.com/reven-erlangga/go-simple-rest-api/models"
	"gorm.io/gorm"
)

func BookFaker(*gorm.DB) *models.Book {
	return &models.Book{
		Title: faker.Name(),
		Description: faker.Paragraph(),
		Author: faker.FirstName(),
		PublishDate: faker.Date(),
		ImageCoverPath: "",
	}
}
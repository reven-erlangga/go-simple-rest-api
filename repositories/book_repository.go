package repositories

import "github.com/reven-erlangga/go-simple-rest-api/models"

type BookRepository interface {
	Save(book models.Book)
	Update(book models.Book)
	Delete(bookId int64)
	FindById(bookId int64)(models.Book, error)
	FindAll()[]models.Book
}
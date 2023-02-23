package repositories

import (
	"errors"

	"github.com/reven-erlangga/go-simple-rest-api/helpers"
	"github.com/reven-erlangga/go-simple-rest-api/models"
	"github.com/reven-erlangga/go-simple-rest-api/request"
	"gorm.io/gorm"
)

type BookRepositoryImpl struct {
	Db *gorm.DB
}

func NewBookRepositoryImpl(Db *gorm.DB) BookRepository {
	return &BookRepositoryImpl{Db: Db}
}

// Delete implements BookRepository
func (n *BookRepositoryImpl) Delete(bookId int64) {
	var book models.Book

	result := n.Db.Where("id = ?", bookId).Delete(&book)

	helpers.ErrorPanic(result.Error)
}

// FindAll implements BookRepository
func (n *BookRepositoryImpl) FindAll() []models.Book {
	var books []models.Book

	result := n.Db.Find(&books)

	helpers.ErrorPanic(result.Error)

	return books
}

// FindById implements BookRepository
func (n *BookRepositoryImpl) FindById(bookId int64) (models.Book, error) {
	var book models.Book

	result := n.Db.Find(&book, bookId)

	if result != nil {
		return book, nil
	} else {
		return book, errors.New("book not found")
	}
}

// Save implements BookRepository
func (n *BookRepositoryImpl) Save(book models.Book) {
	result := n.Db.Create(book)

	helpers.ErrorPanic(result.Error)
}

// Update implements BookRepository
func (n *BookRepositoryImpl) Update(book models.Book) {
	var requestUpdateBook = request.UpdateBookRequest{
		Id: book.Id,
		Title: book.Title,
		Description: book.Description,
		Author: book.Author,
		PublishDate: book.PublishDate,
	}

	result := n.Db.Model(&book).Updates(requestUpdateBook)

	helpers.ErrorPanic(result.Error)
}

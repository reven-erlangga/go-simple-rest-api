package services

import (
	"github.com/go-playground/validator/v10"
	"github.com/reven-erlangga/go-simple-rest-api/helpers"
	"github.com/reven-erlangga/go-simple-rest-api/models"
	"github.com/reven-erlangga/go-simple-rest-api/repositories"
	"github.com/reven-erlangga/go-simple-rest-api/request"
	"github.com/reven-erlangga/go-simple-rest-api/response"
)

type BookServiceImpl struct {
	BookRepository repositories.BookRepository
	Validate       *validator.Validate
}

func NewBookServiceImpl(bookRepository repositories.BookRepository, validate *validator.Validate) BookService {
	return &BookServiceImpl{
		BookRepository: bookRepository,
		Validate:       validate,
	}
}

// Create implements BookService
func (s *BookServiceImpl) Create(book request.CreateBookRequest) {
	err := s.Validate.Struct(book)

	helpers.ErrorPanic(err)

	bookModel := models.Book{
		Title: book.Title,
		Description: book.Description,
		Author: book.Author,
		PublishDate: book.PublishDate,
		ImageCoverPath: book.ImageCoverPath,
	}

	s.BookRepository.Save(bookModel)
}

// Delete implements BookService
func (s *BookServiceImpl) Delete(bookId int64) {
	s.BookRepository.Delete(bookId)
}

// FindAll implements BookService
func (s *BookServiceImpl) FindAll() []response.BookResponse {
	result := s.BookRepository.FindAll()
	var books []response.BookResponse

	for _, v := range result {
		book := response.BookResponse{
			Id: v.Id,
			Title: v.Title,
			Description: v.Description,
			Author: v.Author,
			PublishDate: v.PublishDate,
			ImageCoverPath: v.ImageCoverPath,
		}

		books = append(books, book)
	}

	return books
}

// FindById implements BookService
func (s *BookServiceImpl) FindById(bookId int64) response.BookResponse {
	bookData, err := s.BookRepository.FindById(bookId)

	helpers.ErrorPanic(err)

	bookResponse := response.BookResponse{
		Id: bookData.Id,
		Title: bookData.Title,
		Description: bookData.Description,
		Author: bookData.Author,
		PublishDate: bookData.PublishDate,
		ImageCoverPath: bookData.ImageCoverPath,
	}

	return bookResponse
}

// Update implements BookService
func (s *BookServiceImpl) Update(book request.UpdateBookRequest) {
	bookData, err := s.BookRepository.FindById(book.Id)

	helpers.ErrorPanic(err)

	bookData.Title = book.Title
	bookData.Description = book.Description
	bookData.Author = book.Author
	bookData.PublishDate = book.PublishDate
	bookData.ImageCoverPath = book.ImageCoverPath

	s.BookRepository.Update(bookData)
}

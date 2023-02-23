package services

import (
	"github.com/reven-erlangga/go-simple-rest-api/request"
	"github.com/reven-erlangga/go-simple-rest-api/response"
)

type BookService interface {
	Create(book request.CreateBookRequest)
	Update(book request.UpdateBookRequest)
	Delete(bookId int64)
	FindById(bookId int64) response.BookResponse
	FindAll() []response.BookResponse
}
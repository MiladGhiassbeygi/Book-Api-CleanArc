package command

import (
	"book-api-cleanarc/internal/domain"
	"book-api-cleanarc/internal/repository"
)

type BookCommandService struct {
	bookRepo repository.BookRepository
}

func NewBookCommandService(repo repository.BookRepository) *BookCommandService {
	return &BookCommandService{bookRepo: repo}
}

func (s *BookCommandService) CreateBook(title string, authorID uint) (*domain.Book, error) {
	book := domain.NewBook(title, authorID)
	err := s.bookRepo.Save(book)
	return book, err
}

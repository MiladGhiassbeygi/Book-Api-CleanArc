package query

import (
	"book-api-cleanarc/internal/domain"
	"book-api-cleanarc/internal/repository"
)

type BookQueryService struct {
	bookRepo repository.BookRepository
}

func NewBookQueryService(repo repository.BookRepository) *BookQueryService {
	return &BookQueryService{bookRepo: repo}
}

func (s *BookQueryService) GetAllBooks() ([]domain.Book, error) {
	return s.bookRepo.FindAll()
}

func (s *BookQueryService) GetBookByID(id uint) (*domain.Book, error) {
	return s.bookRepo.FindByID(id)
}

func (s *BookQueryService) GetBooksByAuthor(authorID uint) ([]domain.Book, error) {
	return s.bookRepo.FindByAuthorID(authorID)
}

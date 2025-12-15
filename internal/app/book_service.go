package app

import (
	"book-api-cleanarc/internal/domain"
	"book-api-cleanarc/internal/repository"
)

type BookService struct {
	bookRepo repository.BookRepository
}

func NewBookService(repo repository.BookRepository) *BookService {
	return &BookService{bookRepo: repo}
}

func (s *BookService) CreateBook(title string, authorID uint) (*domain.Book, error) {
	book := domain.NewBook(title, authorID)
	err := s.bookRepo.Save(book)
	return book, err
}

func (s *BookService) GetAllBooks() ([]domain.Book, error) {
	return s.bookRepo.FindAll()
}

func (s *BookService) GetBooksByAuthor(authorID uint) ([]domain.Book, error) {
	return s.bookRepo.FindByAuthorID(authorID)
}

func (s *BookService) GetBookByID(id uint) (*domain.Book, error) {
	return s.bookRepo.FindByID(id)
}

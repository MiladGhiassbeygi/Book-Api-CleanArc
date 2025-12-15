package repository

import "book-api-cleanarc/internal/domain"

type BookRepository interface {
	Save(book *domain.Book) error
	FindAll() ([]domain.Book, error)
	FindByID(id uint) (*domain.Book, error)
	FindByAuthorID(authorID uint) ([]domain.Book, error)
}

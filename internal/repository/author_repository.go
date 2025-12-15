package repository

import "book-api-cleanarc/internal/domain"

type AuthorRepository interface {
	Save(author *domain.Author) error
	FindAll() ([]domain.Author, error)
	FindByID(id uint) (*domain.Author, error)
	Delete(id uint) error
}

package persistence

import (
	"book-api-cleanarc/internal/domain"
	"book-api-cleanarc/internal/repository"

	"gorm.io/gorm"
)

type BookRepoImpl struct {
	DB *gorm.DB
}

func NewBookRepo(db *gorm.DB) repository.BookRepository {
	return &BookRepoImpl{DB: db}
}

func (r *BookRepoImpl) Save(book *domain.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepoImpl) FindAll() ([]domain.Book, error) {
	var books []domain.Book
	err := r.DB.Find(&books).Error
	return books, err
}

func (r *BookRepoImpl) FindByID(id uint) (*domain.Book, error) {
	var book domain.Book
	err := r.DB.First(&book, id).Error
	return &book, err
}

func (r *BookRepoImpl) FindByAuthorID(authorID uint) ([]domain.Book, error) {
	var books []domain.Book
	err := r.DB.Where("author_id = ?", authorID).Find(&books).Error
	return books, err
}

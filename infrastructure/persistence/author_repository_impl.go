package persistence

import (
	"book-api-cleanarc/internal/domain"
	"book-api-cleanarc/internal/repository"

	"gorm.io/gorm"
)

type AuthorRepoImpl struct {
	DB *gorm.DB
}

func NewAuthorRepo(db *gorm.DB) repository.AuthorRepository {
	return &AuthorRepoImpl{DB: db}
}

func (r *AuthorRepoImpl) Save(author *domain.Author) error {
	return r.DB.Create(author).Error
}

func (r *AuthorRepoImpl) Delete(id uint) error {
	var author domain.Author

	if err := r.DB.First(&author, id).Error; err != nil {
		return err
	}

	return r.DB.Delete(&author).Error
}

func (r *AuthorRepoImpl) FindAll() ([]domain.Author, error) {
	var authors []domain.Author
	err := r.DB.Preload("Books").Find(&authors).Error
	return authors, err
}

func (r *AuthorRepoImpl) FindByID(id uint) (*domain.Author, error) {
	var author domain.Author
	err := r.DB.Preload("Books").First(&author, id).Error
	return &author, err
}

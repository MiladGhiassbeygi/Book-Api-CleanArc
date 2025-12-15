package query

import (
	"book-api-cleanarc/internal/domain"
	"book-api-cleanarc/internal/repository"
)

type AuthorQueryService struct {
	authorRepo repository.AuthorRepository
}

func NewAuthorQueryService(repo repository.AuthorRepository) *AuthorQueryService {
	return &AuthorQueryService{authorRepo: repo}
}

func (s *AuthorQueryService) GetAllAuthors() ([]domain.Author, error) {
	return s.authorRepo.FindAll()
}

func (s *AuthorQueryService) GetAuthorByID(id uint) (*domain.Author, error) {
	return s.authorRepo.FindByID(id)
}

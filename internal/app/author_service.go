package app

import (
	"book-api-cleanarc/internal/domain"
	"book-api-cleanarc/internal/repository"
)

type AuthorService struct {
	authorRepo repository.AuthorRepository
}

func NewAuthorService(repo repository.AuthorRepository) *AuthorService {
	return &AuthorService{authorRepo: repo}
}

func (s *AuthorService) CreateAuthor(name string) (*domain.Author, error) {
	author := domain.NewAuthor(name)
	err := s.authorRepo.Save(author)
	return author, err
}

func (s *AuthorService) GetAllAuthors() ([]domain.Author, error) {
	return s.authorRepo.FindAll()
}

func (s *AuthorService) GetAuthorByID(id uint) (*domain.Author, error) {
	return s.authorRepo.FindByID(id)
}

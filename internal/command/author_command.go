package command

import (
	"book-api-cleanarc/internal/domain"
	"book-api-cleanarc/internal/repository"
)

type AuthorCommandService struct {
	authorRepo repository.AuthorRepository
}

func (s *AuthorCommandService) Delete(id uint) (domain.Author, error) {
	author, err := s.authorRepo.FindByID(id)
	if err != nil {
		return domain.Author{}, err
	}

	err = s.authorRepo.Delete(id)
	if err != nil {
		return domain.Author{}, err
	}

	return *author, nil
}

func NewAuthorCommandService(repo repository.AuthorRepository) *AuthorCommandService {
	return &AuthorCommandService{authorRepo: repo}
}

func (s *AuthorCommandService) CreateAuthor(name string) (*domain.Author, error) {
	author := domain.NewAuthor(name)
	err := s.authorRepo.Save(author)
	return author, err
}

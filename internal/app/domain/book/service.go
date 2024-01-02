package book

import (
	"context"

	"github.com/literalog/library/pkg/models"
)

type Service interface {
	Create(ctx context.Context, b *models.Book) error
	Update(ctx context.Context, b *models.Book) error
	Delete(ctx context.Context, id string) error
	GetById(ctx context.Context, id string) (*models.Book, error)
	GetAll(ctx context.Context) ([]models.Book, error)
}

type service struct {
	repository Repository
	validator  Validator
}

func NewService(repo Repository) Service {
	return &service{
		repository: repo,
	}
}

func (s *service) Create(ctx context.Context, b *models.Book) error {
	if err := s.validator.Validate(b); err != nil {
		return err
	}
	return s.repository.Create(ctx, b)
}

func (s *service) Update(ctx context.Context, b *models.Book) error {
	return s.repository.Update(ctx, b)
}

func (s *service) Delete(ctx context.Context, id string) error {
	if id == "" {
		return ErrEmptyId
	}
	return s.repository.Delete(ctx, id)
}

func (s *service) GetById(ctx context.Context, id string) (*models.Book, error) {
	if id == "" {
		return nil, ErrEmptyId
	}
	return s.repository.GetById(ctx, id)
}

func (s *service) GetAll(ctx context.Context) ([]models.Book, error) {
	return s.repository.GetAll(ctx)
}

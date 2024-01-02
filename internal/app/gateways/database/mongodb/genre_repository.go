package mongodb

import (
	"context"

	"github.com/literalog/library/internal/app/domain/genre"
	"github.com/literalog/library/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type GenreRepository struct {
	collection *mongo.Collection
}

func NewGenreRepository(collection *mongo.Collection) genre.Repository {
	return GenreRepository{
		collection: collection,
	}
}

func (r GenreRepository) Create(ctx context.Context, g *models.Genre) error {
	return nil
}

func (r GenreRepository) Update(ctx context.Context, g *models.Genre) error {
	return nil
}

func (r GenreRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r GenreRepository) GetById(ctx context.Context, id string) (*models.Genre, error) {
	return nil, nil
}

func (r GenreRepository) GetAll(ctx context.Context) ([]models.Genre, error) {
	return nil, nil
}

package mongodb

import (
	"context"

	"github.com/literalog/library/pkg/models"

	"github.com/literalog/library/internal/app/domain/series"

	"go.mongodb.org/mongo-driver/mongo"
)

type SeriesRepository struct {
	collection *mongo.Collection
}

func NewSeriesRepository(collection *mongo.Collection) series.Repository {
	return SeriesRepository{
		collection: collection,
	}
}

func (r SeriesRepository) Create(ctx context.Context, s *models.Series) error {
	return nil
}

func (r SeriesRepository) Update(ctx context.Context, s *models.Series) error {
	return nil
}

func (r SeriesRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r SeriesRepository) GetById(ctx context.Context, id string) (*models.Series, error) {
	return nil, nil
}

func (r SeriesRepository) GetAll(ctx context.Context) ([]models.Series, error) {
	return nil, nil
}

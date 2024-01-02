package mongodb

import (
	"context"

	"github.com/literalog/library/internal/app/domain/book"
	"github.com/literalog/library/pkg/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type BookRepository struct {
	collection *mongo.Collection
}

func NewBookRepository(collection *mongo.Collection) book.Repository {
	return BookRepository{
		collection: collection,
	}
}

func (r BookRepository) Create(ctx context.Context, b *models.Book) error {
	// coleção de genre se tem essa parada ai
	return nil
}

func (r BookRepository) Update(ctx context.Context, b *models.Book) error {
	return nil
}

func (r BookRepository) Delete(ctx context.Context, id string) error {
	return nil
}

func (r BookRepository) GetById(ctx context.Context, id string) (*models.Book, error) {
	return nil, nil
}
func (r BookRepository) GetAll(ctx context.Context) ([]models.Book, error) {
	return nil, nil
}

package repositories

import (
	"context"

	"devopsProjectModule.com/gl5/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Repository interface {
	GetAll(ctx context.Context) ([]*models.Product, error)
	GetByID(ctx context.Context, id string) (*models.Product, error)
	Create(ctx context.Context, product *models.Product) (*primitive.ObjectID, error)
	Update(ctx context.Context, product *models.Product) error
	Delete(ctx context.Context, id string) error
}

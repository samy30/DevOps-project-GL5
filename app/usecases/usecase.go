package usecases

import (
	"context"

	"devopsProjectModule.com/gl5/models"
)

type UseCase interface {
	GetProducts(ctx context.Context) ([]models.Product, error)
	GetProductByID(ctx context.Context, id string) (models.Product, error)
	CreateProduct(ctx context.Context, product models.Product) error
	UpdateProduct(ctx context.Context, product models.Product) error
	DeleteProduct(ctx context.Context, id string) error
}

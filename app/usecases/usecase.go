package usecases

import (
	"context"

	"devopsProjectModule.com/gl5/models"
	"devopsProjectModule.com/gl5/payload"
)

type UseCase interface {
	GetProducts(ctx context.Context) ([]*models.Product, error)
	GetProductByID(ctx context.Context, id string) (*models.Product, error)
	CreateProduct(ctx context.Context, product *models.Product) error
	UpdateProduct(ctx context.Context, product *models.Product) error
	DeleteProduct(ctx context.Context, id string) error
	BuyProduct(ctx context.Context, BuyRequest *payload.BuyRequest) error
	GetTransactions(ctx context.Context) ([]*models.Transaction, error)
}

package repositories

import (
	"context"

	"devopsProjectModule.com/gl5/models"
)

type TransactRepository interface {
	GetAll(ctx context.Context) ([]*models.Transaction, error)
	GetByID(ctx context.Context, id string) (*models.Transaction, error)
	Create(ctx context.Context, transaction *models.Transaction) error
}

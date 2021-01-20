package repositories

import (
	"context"
	"errors"
	"strconv"

	"devopsProjectModule.com/gl5/models"
)

//ProductRepositoryTest in memory repo
type transactionRepositoryTest struct {
	m map[string]*models.Transaction
}

//ProductRepositoryTest create new repository
func NewTransactionRepositoryTest() TransactRepository {
	var m = map[string]*models.Transaction{}
	return &transactionRepositoryTest{
		m: m,
	}
}

// Get all the products from the database.
func (r transactionRepositoryTest) GetAll(ctx context.Context) ([]*models.Transaction, error) {
	var transactions []*models.Transaction

	for _, j := range r.m {
		transactions = append(transactions, j)
	}
	return transactions, nil
}

// Get product with the specified ID from the database.
func (r transactionRepositoryTest) GetByID(ctx context.Context, id string) (*models.Transaction, error) {
	if r.m[id] == nil {
		return nil, errors.New("mongo: no documents in result")
	}
	return r.m[id], nil
}

// Create product with the specified object
func (r transactionRepositoryTest) Create(ctx context.Context, transaction *models.Transaction) error {
	transaction.ID = strconv.Itoa(len(r.m))
	r.m[transaction.ID] = transaction
	return nil
}

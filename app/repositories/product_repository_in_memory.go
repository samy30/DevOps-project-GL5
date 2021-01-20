package repositories

import (
	"context"
	"errors"
	"strconv"

	"devopsProjectModule.com/gl5/models"
)

//ProductRepositoryTest in memory repo
type productRepositoryTest struct {
	m map[string]*models.Product
}

//ProductRepositoryTest create new repository
func NewProductRepositoryTest() Repository {
	var m = map[string]*models.Product{}
	return &productRepositoryTest{
		m: m,
	}
}

// Get all the products from the database.
func (r productRepositoryTest) GetAll(ctx context.Context) ([]*models.Product, error) {
	var products []*models.Product

	for _, j := range r.m {
		products = append(products, j)
	}
	return products, nil
}

// Get product with the specified ID from the database.
func (r productRepositoryTest) GetByID(ctx context.Context, id string) (*models.Product, error) {
	if r.m[id] == nil {
		return nil, errors.New("mongo: no documents in result")
	}
	return r.m[id], nil
}

// Create product with the specified object
func (r productRepositoryTest) Create(ctx context.Context, product *models.Product) error {
	product.ID = strconv.Itoa(len(r.m) + 1)
	r.m[product.ID] = product
	return nil
}

// Update saves the changes to a product in the database.
func (r productRepositoryTest) Update(ctx context.Context, product *models.Product) error {

	if r.m[product.ID] == nil {
		return errors.New("mongo: no documents in result")
	}

	r.m[product.ID] = product
	return nil
}

// Delete deletes a product with the specified ID from the database.
func (r productRepositoryTest) Delete(ctx context.Context, id string) error {
	if r.m[id] == nil {
		return errors.New("mongo: no documents in result")
	}
	delete(r.m, id)
	return nil
}

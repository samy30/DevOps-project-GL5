package usecases

import (
	"context"
	"errors"
	"fmt"
	"time"

	"devopsProjectModule.com/gl5/logger"
	"devopsProjectModule.com/gl5/models"
	"devopsProjectModule.com/gl5/payload"
	"devopsProjectModule.com/gl5/repositories"
)

type productUseCase struct {
	productRepository     repositories.Repository
	transactionRepository repositories.TransactRepository
}

// NewProductUseCase : construct a new productUseCase
func NewProductUseCase(
	productRepository repositories.Repository,
	transactionRepository repositories.TransactRepository) UseCase {

	return &productUseCase{
		productRepository:     productRepository,
		transactionRepository: transactionRepository,
	}
}

// GetProducts : call the repository to retreive all products
func (p productUseCase) GetProducts(ctx context.Context) ([]*models.Product, error) {
	logger.Info("get products request sent")
	return p.productRepository.GetAll(ctx)
}

// GetProductByID : call the repository to retreive product from database by id
func (p productUseCase) GetProductByID(ctx context.Context, id string) (*models.Product, error) {
	logger.Info(fmt.Sprintf("get product by id %s request sent\n", id))
	return p.productRepository.GetByID(ctx, id)
}

// CreateProduct : call the repository to persist a product to the database & contains creation logic
func (p productUseCase) CreateProduct(ctx context.Context, product *models.Product) error {
	product.Quantity = product.InitialQuantity
	logger.Info(fmt.Sprintf("create product %#v request sent\n", product))
	return p.productRepository.Create(ctx, product)
}

// UpdateProduct : call the repository to persist a product update to the datbase
func (p productUseCase) UpdateProduct(ctx context.Context, product *models.Product) error {
	logger.Info(fmt.Sprintf("update product %#v request sent\n", product))
	return p.productRepository.Update(ctx, product)
}

// DeleteProduct : call the repository to delete a product from the database
func (p productUseCase) DeleteProduct(ctx context.Context, id string) error {
	logger.Info(fmt.Sprintf("delete product by id %s request sent\n", id))
	return p.productRepository.Delete(ctx, id)
}

// BuyProduct : contains the logic of the buying operation
func (p productUseCase) BuyProduct(ctx context.Context, buyRequest *payload.BuyRequest) error {
	product, err1 := p.productRepository.GetByID(ctx, buyRequest.ProductId)

	if err1 != nil {
		return err1
	}

	newQuantity, err2 := p.calculateNewProductsQuantity(product.Quantity, buyRequest.Quantity)

	product.Quantity = newQuantity
	if err2 != nil {
		return err2
	}

	err3 := p.productRepository.Update(ctx, product)

	if err3 != nil {
		return err3
	}

	var transaction *models.Transaction
	transaction.Date = time.Now().String()
	transaction.Quantity = buyRequest.Quantity
	transaction.Product = product

	return p.transactionRepository.Create(ctx, transaction)
}

// GetTransactions : call the transactions repository to retreive all trnsactions from the database
func (p productUseCase) GetTransactions(ctx context.Context) ([]*models.Transaction, error) {
	return p.transactionRepository.GetAll(ctx)
}

// calculateNewProductsQuantity : business method to calculate the new quantity of a product after a buying operation
func (p productUseCase) calculateNewProductsQuantity(initialQuantity int, boughtQuantity int) (int, error) {
	newQuantity := initialQuantity - boughtQuantity
	if newQuantity < 0 {
		return -1, errors.New("out of stock")
	}
	return newQuantity, nil
}

package usecases

import (
	"context"
	"errors"
	"time"

	"devopsProjectModule.com/gl5/models"
	"devopsProjectModule.com/gl5/payload"
	"devopsProjectModule.com/gl5/repositories"
)

type ProductUseCase struct {
	productRepository     repositories.Repository
	transactionRepository repositories.TransactRepository
}

func NewProductUseCase(productRepository repositories.Repository, transactionRepository repositories.TransactRepository) UseCase {
	return &ProductUseCase{
		productRepository:     productRepository,
		transactionRepository: transactionRepository,
	}
}

func (p ProductUseCase) GetProducts(ctx context.Context) ([]models.Product, error) {
	return p.productRepository.GetAll(ctx)
}

func (p ProductUseCase) GetProductByID(ctx context.Context, id string) (models.Product, error) {
	return p.productRepository.GetByID(ctx, id)
}

func (p ProductUseCase) CreateProduct(ctx context.Context, product models.Product) error {
	product.InitialQuantity = product.Quantity
	return p.productRepository.Create(ctx, product)
}

func (p ProductUseCase) UpdateProduct(ctx context.Context, product models.Product) error {
	return p.productRepository.Update(ctx, product)
}

func (p ProductUseCase) DeleteProduct(ctx context.Context, id string) error {
	return p.productRepository.Delete(ctx, id)
}

func (p ProductUseCase) BuyProduct(ctx context.Context, buyRequest payload.BuyRequest) error {
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

	var transaction models.Transaction
	transaction.Date = time.Now().String()
	transaction.Quantity = buyRequest.Quantity
	transaction.Product = &product

	return p.transactionRepository.Create(ctx, transaction)
}

func (p ProductUseCase) GetTransactions(ctx context.Context) ([]models.Transaction, error) {
	return p.transactionRepository.GetAll(ctx)
}

func (p ProductUseCase) calculateNewProductsQuantity(initialQuantity int, boughtQuantity int) (int, error) {
	newQuantity := initialQuantity - boughtQuantity
	if newQuantity < 0 {
		return -1, errors.New("out of stock")
	}
	return newQuantity, nil
}

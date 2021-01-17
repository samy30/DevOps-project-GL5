package usecases

import (
	"context"

	"devopsProjectModule.com/gl5/models"
	"devopsProjectModule.com/gl5/repositories"
)

type ProductUseCase struct {
	productRepository repositories.Repository
}

func NewProductUseCase(productRepository repositories.Repository) UseCase {
	return &ProductUseCase{
		productRepository: productRepository,
	}
}

func (p ProductUseCase) GetProducts(ctx context.Context) ([]models.Product, error) {
	return p.productRepository.GetAll(ctx)
}

func (p ProductUseCase) GetProductByID(ctx context.Context, id string) (models.Product, error) {
	return p.productRepository.GetByID(ctx, id)
}

func (p ProductUseCase) CreateProduct(ctx context.Context, product models.Product) error {
	return p.productRepository.Create(ctx, product)
}

func (p ProductUseCase) UpdateProduct(ctx context.Context, product models.Product) error {
	return p.productRepository.Update(ctx, product)
}

func (p ProductUseCase) DeleteProduct(ctx context.Context, id string) error {
	return p.productRepository.Delete(ctx, id)
}

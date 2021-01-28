package usecases

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"devopsProjectModule.com/gl5/logger"
	"devopsProjectModule.com/gl5/models"
	"devopsProjectModule.com/gl5/repositories"
	"github.com/stretchr/testify/assert"
)

var productUsecase UseCase

// Test the creation, listing, deletion and updating of products using the mockDatabase ( inmemory db)
func TestCreateProduct(t *testing.T) {
	productUsecase = NewProductUseCase(repositories.NewProductRepositoryTest(), nil)
	product := &models.Product{
		Title:           "Product 1",
		Price:           50,
		InitialQuantity: 45,
	}

	_, err := productUsecase.CreateProduct(context.TODO(), product)
	assert.Nil(t, err)
}

func TestCreateProductLog(t *testing.T) {
	productUsecase = NewProductUseCase(
		repositories.NewProductRepositoryTest(), nil)
	buffer := bytes.NewBuffer(nil)
	log := logger.NewLogger(buffer)
	logger.SetDefaultLogger(log)
	product := &models.Product{
		Title:           "Product 1",
		Price:           50,
		InitialQuantity: 45,
	}

	productUsecase.CreateProduct(context.Background(), product)
	s := buffer.String()
	if !strings.Contains(s, "Info:create product &models.Product{ID:\"\", Title:\"Product 1\", Price:50, Quantity:45, InitialQuantity:45, Category:(*models.Category)(nil)} request sent") {
		t.Fail()
	}
}

func TestGetEmptyProducts(t *testing.T) {
	productUsecase = NewProductUseCase(
		repositories.NewProductRepositoryTest(), nil)
	products, err := productUsecase.GetProducts(context.TODO())
	if err != nil {
		t.Fail()
	}

	if len(products) != 0 {
		t.Fail()
	}
}

func TestGetProducts(t *testing.T) {
	productUsecase = NewProductUseCase(
		repositories.NewProductRepositoryTest(), nil)
	product1 := &models.Product{
		Title:           "Product 1",
		Price:           50,
		InitialQuantity: 45,
	}
	product2 := &models.Product{
		Title:           "Product 2",
		Price:           20,
		InitialQuantity: 20,
	}
	productUsecase.CreateProduct(context.TODO(), product1)
	productUsecase.CreateProduct(context.TODO(), product2)
	products, err := productUsecase.GetProducts(context.TODO())
	if err != nil {
		t.Fail()
	}

	if len(products) != 2 {
		t.Fail()
	}

	if products[0].Title != "Product 1" ||
		products[0].Price != 50 ||
		products[0].InitialQuantity != 45 ||
		products[0].Quantity != 45 {
		t.Fail()
	}

	if products[1].Title != "Product 2" ||
		products[1].Price != 20 ||
		products[1].InitialQuantity != 20 ||
		products[1].Quantity != 20 {
		t.Fail()
	}
}

func TestGetProductByID(t *testing.T) {
	productUsecase = NewProductUseCase(
		repositories.NewProductRepositoryTest(), nil)
	product1 := &models.Product{
		Title:           "Product 1",
		Price:           50,
		InitialQuantity: 45,
	}
	product2 := &models.Product{
		Title:           "Product 2",
		Price:           20,
		InitialQuantity: 20,
	}
	id1, _ := productUsecase.CreateProduct(context.TODO(), product1)
	productUsecase.CreateProduct(context.TODO(), product2)

	product, err := productUsecase.GetProductByID(context.TODO(), id1.Hex())

	if err != nil {
		t.Fail()
	}

	if product == nil {
		t.Fail()
	}

	if product.Title != "Product 1" ||
		product.ID != id1.Hex() ||
		product.Price != 50 ||
		product.InitialQuantity != 45 ||
		product.Quantity != 45 {
		t.Fail()
	}
}

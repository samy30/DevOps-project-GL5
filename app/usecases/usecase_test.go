package usecases

import (
	"bytes"
	"context"
	"fmt"
	"testing"

	"devopsProjectModule.com/gl5/logger"
	"devopsProjectModule.com/gl5/models"
	"devopsProjectModule.com/gl5/repositories"
	"github.com/stretchr/testify/assert"
)

var productUsecase UseCase

func init() {
	fmt.Println("initing")
	productUsecase = NewProductUseCase(
		repositories.NewProductRepositoryTest(), nil)
}

func TestCreateProduct(t *testing.T) {

	product := &models.Product{
		Title:           "Product 1",
		Price:           50,
		InitialQuantity: 45,
	}

	err := productUsecase.CreateProduct(context.TODO(), product)
	assert.Nil(t, err)
}

func TestCreateProductLog(t *testing.T) {
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
	fmt.Println(s)
}

func TestGetProducts(t *testing.T) {
	products, err := productUsecase.GetProducts(context.TODO())
	if err != nil {
		fmt.Printf("error")
	}
	fmt.Printf("%+v\n", products[0])
}

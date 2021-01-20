package usecases

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"devopsProjectModule.com/gl5/models"
	"devopsProjectModule.com/gl5/repositories"
	"github.com/stretchr/testify/assert"
)

var productUsecase UseCase

func init() {
	fmt.Println("initing")
	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	productUsecase = NewProductUseCase(
		repositories.NewProductRepositoryTest(), nil, logger, logger, logger)
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

func TestGetProducts(t *testing.T) {
	products, err := productUsecase.GetProducts(context.TODO())
	if err != nil {
		fmt.Printf("error")
	}
	fmt.Printf("%+v\n", products[0])
}

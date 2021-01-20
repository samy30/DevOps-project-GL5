package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	helper "devopsProjectModule.com/gl5/helpers"
	"devopsProjectModule.com/gl5/models"
	"devopsProjectModule.com/gl5/payload"
	"devopsProjectModule.com/gl5/repositories"
	"devopsProjectModule.com/gl5/usecases"
	"github.com/gorilla/mux"
)

type ProductController struct {
	productUseCase usecases.UseCase
	ctx            context.Context
}

func NewProductController(
	databaseName string,
	username string,
	password string,
	warningLogger *log.Logger,
	infoLogger *log.Logger,
	errorLogger *log.Logger) *ProductController {
	db := helper.ConnectDB("products", databaseName, username, password)
	db_transactions := helper.ConnectDB("transactions", databaseName, username, password)
	productRepository := repositories.NewProductRepository(db)
	transactionRepository := repositories.NewTransactionRepository(db_transactions)
	return &ProductController{
		productUseCase: usecases.NewProductUseCase(
			productRepository,
			transactionRepository,
			warningLogger,
			infoLogger,
			errorLogger),
		ctx: context.TODO(),
	}
}

func (p ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	products, err := p.productUseCase.GetProducts(p.ctx)

	if err != nil {
		helper.GetError(errors.New("Something went wrong Please retry later"), w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

func (p ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := params["id"]

	product, err := p.productUseCase.GetProductByID(p.ctx, id)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			helper.GetError(errors.New("No product found with the specified id"), w, http.StatusBadRequest)
			return
		}
		helper.GetError(errors.New("Something went wrong Please retry later"), w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (p ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product *models.Product

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&product)

	if product.Title == "" {
		helper.GetError(errors.New("No 'title' specified"), w, http.StatusBadRequest)
		return
	}

	if product.Price == 0 {
		helper.GetError(errors.New("No 'price' specified"), w, http.StatusBadRequest)
		return
	}

	if product.InitialQuantity == 0 {
		helper.GetError(errors.New("No 'initial_quantity' specified"), w, http.StatusBadRequest)
		return
	}

	err := p.productUseCase.CreateProduct(p.ctx, product)

	if err != nil {
		helper.GetError(errors.New("Something went wrong Please retry later"), w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Product created successfully")
}

func (p ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product *models.Product

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := params["id"]

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&product)

	product.ID = id
	err := p.productUseCase.UpdateProduct(p.ctx, product)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			helper.GetError(errors.New("No product found with the specified id"), w, http.StatusBadRequest)
			return
		}
		helper.GetError(errors.New("Something went wrong Please retry later"), w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Product updated successfully")
}

func (p ProductController) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := params["id"]

	err := p.productUseCase.DeleteProduct(p.ctx, id)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			helper.GetError(errors.New("No product found with the specified id"), w, http.StatusBadRequest)
			return
		}
		helper.GetError(errors.New("Something went wrong Please retry later"), w, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	json.NewEncoder(w).Encode("Product deleted successfully")
}

func (p ProductController) BuyProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var buyRequest *payload.BuyRequest

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&buyRequest)

	if buyRequest.ProductId == "" {
		helper.GetError(errors.New("No 'product_id' specified"), w, http.StatusBadRequest)
		return
	}

	if buyRequest.Quantity == 0 {
		helper.GetError(errors.New("product attribute 'quantity' must be specified"), w, http.StatusBadRequest)
		return
	}

	err := p.productUseCase.BuyProduct(p.ctx, buyRequest)

	if err != nil {
		if err.Error() == "mongo: no documents in result" {
			helper.GetError(errors.New("No product found with the specified id"), w, http.StatusBadRequest)
			return
		}
		if err.Error() == "out of stock" {
			helper.GetError(err, w, http.StatusBadRequest)
			return
		}
		helper.GetError(errors.New("Something went wrong Please retry later"), w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Transaction finished successfully")
}

func (p ProductController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := p.productUseCase.GetTransactions(p.ctx)

	if err != nil {
		helper.GetError(errors.New("Something went wrong Please retry later"), w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transactions)
}

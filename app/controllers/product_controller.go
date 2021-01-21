package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	helper "devopsProjectModule.com/gl5/helpers"
	"devopsProjectModule.com/gl5/models"
	"devopsProjectModule.com/gl5/payload"
	"devopsProjectModule.com/gl5/repositories"
	"devopsProjectModule.com/gl5/usecases"
	"github.com/gorilla/mux"
)

// ProductController : a struct that declares the app's logic handler
type ProductController struct {
	productUseCase usecases.UseCase
	ctx            context.Context
}

// NewProductController : constructs the ProductController struct
func NewProductController(
	databaseName string,
	username string,
	password string) *ProductController {
	db := helper.ConnectDB(databaseName, username, password)
	productsCollection := helper.GetCollection("products", db)
	transactionsCollection := helper.GetCollection("transactions", db)
	productRepository := repositories.NewProductRepository(productsCollection)
	transactionRepository := repositories.NewTransactionRepository(transactionsCollection)
	return &ProductController{
		productUseCase: usecases.NewProductUseCase(
			productRepository,
			transactionRepository),
		ctx: context.TODO(),
	}
}

// GetProducts : an http handler to list products
func (p ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, err := p.productUseCase.GetProducts(p.ctx)

	if err != nil {
		helper.GetError(errors.New("Something went wrong Please retry later"), w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(products)
}

// GetProduct : an http handler to get a product by specified id in request
func (p ProductController) GetProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

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

// CreateProduct : http handler to post a new Product
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

// UpdateProduct : an http handler to update a given product with the given attributes
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

// DeleteProduct : an http handler to delete a product by id
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

// BuyProduct : an http handler to the buy product functionality
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

// GetTransactions : an http handler to list all transactions
func (p ProductController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := p.productUseCase.GetTransactions(p.ctx)

	if err != nil {
		helper.GetError(errors.New("Something went wrong Please retry later"), w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transactions)
}

package controllers

import (
	"context"
	"encoding/json"
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

func NewProductController() *ProductController {
	db := helper.ConnectDB("products")
	db_transactions := helper.ConnectDB("transactions")
	productRepository := repositories.NewProductRepository(db)
	transactionRepository := repositories.NewTransactionRepository(db_transactions)
	return &ProductController{
		productUseCase: usecases.NewProductUseCase(productRepository, transactionRepository),
		ctx:            context.TODO(),
	}
}

func (p ProductController) GetProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	products, err := p.productUseCase.GetProducts(p.ctx)

	if err != nil {
		helper.GetError(err, w, http.StatusInternalServerError)
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
		helper.GetError(err, w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(product)
}

func (p ProductController) CreateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&product)

	err := p.productUseCase.CreateProduct(p.ctx, product)

	if err != nil {
		helper.GetError(err, w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Product created successfully")
}

func (p ProductController) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := params["id"]

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&product)

	product.ID = id
	err := p.productUseCase.UpdateProduct(p.ctx, product)

	if err != nil {
		helper.GetError(err, w, http.StatusInternalServerError)
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
		helper.GetError(err, w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Product deleted successfully")
}

func (p ProductController) BuyProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var buyRequest payload.BuyRequest

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&buyRequest)

	err := p.productUseCase.BuyProduct(p.ctx, buyRequest)

	if err != nil {
		helper.GetError(err, w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode("Transaction finished successfully")
}

func (p ProductController) GetTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	transactions, err := p.productUseCase.GetTransactions(p.ctx)

	if err != nil {
		helper.GetError(err, w, http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(transactions)
}

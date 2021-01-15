package main

import (
	"log"
	"net/http"

	productService "gl5-project/services"

	"github.com/gorilla/mux"
)

func main() {
	//Init Router
	r := mux.NewRouter()

	// arrange our routes
	r.HandleFunc("/api/products", productService.GetProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", productService.GetProduct).Methods("GET")
	r.HandleFunc("/api/products", productService.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", productService.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/products/{id}", productService.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/api/products/{id}/buy", productService.BuyProduct).Methods("POST")

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}

func Hello() string {
	return "Hello, world."
}

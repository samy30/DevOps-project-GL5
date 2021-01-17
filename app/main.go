package main

import (
	"log"
	"net/http"

	"devopsProjectModule.com/gl5/controllers"
	"github.com/gorilla/mux"
)

func main() {
	controller := controllers.NewProductController()
	//Init Router
	r := mux.NewRouter()

	// arrange our routes
	r.HandleFunc("/api/products", controller.GetProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", controller.GetProduct).Methods("GET")
	r.HandleFunc("/api/products", controller.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", controller.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/products/{id}", controller.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/api/products/buy", controller.BuyProduct).Methods("POST")

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}

func Hello() string {
	return "Hello, world."
}

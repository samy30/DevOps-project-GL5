package main

import (
	helper "DevOps-project-GL5/app/helpers"
	"DevOps-project-GL5/app/models"
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Connection mongoDB with helper class
// var collection *mongo.Collection
var collection = helper.ConnectDB()

// ************************************
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&product)

	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), product)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// ************************************

func main() {

	//Init Router
	r := mux.NewRouter()

	// arrange our route
	/*r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")*/
	r.HandleFunc("/api/products", createProduct).Methods("POST")
	/*r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")*/

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}

package main

import (
	helper "DevOps-project-GL5/app/helpers"
	"DevOps-project-GL5/app/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//Connection mongoDB with helper class
// var collection *mongo.Collection
var collection = helper.ConnectDB("products")
var transactionsCollection = helper.ConnectDB("transactions")

// ******************************************************************************************

func getProducts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// we created Book array
	var products []models.Product

	// bson.M{},  we passed empty filter. So we want to get all data.
	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.GetError(err, w)
		return
	}

	// Close the cursor once finished
	/*A defer statement defers the execution of a function until the surrounding function returns.
	simply, run cur.Close() process but after cur.Next() finished.*/
	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var product models.Product
		// & character returns the memory address of the following variable.
		err := cur.Decode(&product) // decode similar to deserialize process.
		if err != nil {
			log.Fatal(err)
		}

		// add item our array
		products = append(products, product)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(products) // encode similar to serialize process.
}

// ******************************************************************************************

func getProduct(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&product)

	if err != nil {
		helper.GetError(err, w)
		// return
	}

	json.NewEncoder(w).Encode(product)
}

// ******************************************************************************************

func updateProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var params = mux.Vars(r)

	//Get id from parameters
	id, _ := primitive.ObjectIDFromHex(params["id"])

	var product models.Product

	// Create filter
	filter := bson.M{"_id": id}

	// Read update model from body request
	_ = json.NewDecoder(r.Body).Decode(&product)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"title", product.Title},
			{"price", product.Price},
			{"initial_quantity", product.InitialQuantity},
			{"category", bson.D{
				{"name", product.Category.Name},
			}},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&product)

	err = collection.FindOne(context.TODO(), filter).Decode(&product)
	if err != nil {
		helper.GetError(err, w)
		return
	}

	product.ID = id

	json.NewEncoder(w).Encode(product)
}

// ******************************************************************************************

func deleteProduct(w http.ResponseWriter, r *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// get params
	var params = mux.Vars(r)

	// string to primitve.ObjectID
	id, err := primitive.ObjectIDFromHex(params["id"])

	// prepare filter.
	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)
}

// ******************************************************************************************
func createProduct(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var product models.Product

	// we decode our body request params
	_ = json.NewDecoder(r.Body).Decode(&product)

	product.Quantity = product.InitialQuantity
	// insert our book model.
	result, err := collection.InsertOne(context.TODO(), product)

	if err != nil {
		helper.GetError(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)
}

// ************************************

func buyProduct(w http.ResponseWriter, r *http.Request) {
	// set header.
	w.Header().Set("Content-Type", "application/json")

	var product models.Product
	var transaction models.Transaction
	// we get params with mux.
	var params = mux.Vars(r)

	// string to primitive.ObjectID
	id, _ := primitive.ObjectIDFromHex(params["id"])

	// We create filter. If it is unnecessary to sort data for you, you can use bson.M{}
	filter := bson.M{"_id": id}

	err := collection.FindOne(context.TODO(), filter).Decode(&product)

	// Read update model from body request
	// _ = json.NewDecoder(r.Body).Decode(&product)

	// prepare update model.
	update := bson.D{
		{"$set", bson.D{
			{"quantity", product.Quantity - 1},
		}},
	}

	err = collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&product)

	err = collection.FindOne(context.TODO(), filter).Decode(&product)

	transaction.Quantity = 1
	transaction.Date = time.Now().String()
	transaction.Product = &product
	result, err2 := transactionsCollection.InsertOne(context.TODO(), transaction)

	if err != nil {
		helper.GetError(err, w)
		return
	}
	if err2 != nil {
		helper.GetError(err2, w)
		return
	}

	fmt.Println("InsertOne() API result:", result)

	product.ID = id

	json.NewEncoder(w).Encode(product)
}

func main() {

	//Init Router
	r := mux.NewRouter()

	// arrange our route
	r.HandleFunc("/api/products", getProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", getProduct).Methods("GET")
	r.HandleFunc("/api/products", createProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", updateProduct).Methods("PUT")
	r.HandleFunc("/api/products/{id}", deleteProduct).Methods("DELETE")

	r.HandleFunc("/api/products/{id}/buy", buyProduct).Methods("POST")

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))

}

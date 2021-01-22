package helper

// fmt: Package to format I/O, similar to printf & scanf.
// log: Logging package.
import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ConnectDB : This is helper function to connect mongoDB
// If you want to export your function. You must to start upper case function name. Otherwise you won't see your function when you import that on other class.
func ConnectDB() *mongo.Client {

	fmt.Println("Connecting to MongoDB...")
	credentials := options.Credential{
		Username: os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		Password: os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
	}

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://" + os.Getenv("MONGO_DB_HOST") + "/" + os.Getenv("MONGO_INITDB_DATABASE")).SetAuth(credentials)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(errors.New("Connection to database failed"))
	}

	fmt.Println("Connected to MongoDB!")

	return client
}

// GetCollection : This is a helper function to get a collection from the provided database client
func GetCollection(collectionName string, client *mongo.Client) *mongo.Collection {
	return client.Database("devopsProjectDB").Collection(collectionName)
}

// ErrorResponse : This is the error model.
type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"message"`
}

// GetError : This is helper function to prepare error model.
func GetError(err error, w http.ResponseWriter, statusCode int) {

	log.Printf(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   statusCode,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

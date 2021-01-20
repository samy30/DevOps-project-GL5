package main

import (
	"log"
	"os"

	"github.com/codegangsta/negroni"

	"devopsProjectModule.com/gl5/controllers"
	"devopsProjectModule.com/gl5/metric"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	warningLogger *log.Logger
	infoLogger    *log.Logger
	errorLogger   *log.Logger
)

func main() {
	initLogger()
	controller := controllers.NewProductController(
		os.Getenv("MONGO_INITDB_DATABASE"),
		os.Getenv("MONGO_INITDB_ROOT_USERNAME"),
		os.Getenv("MONGO_INITDB_ROOT_PASSWORD"),
		warningLogger,
		infoLogger,
		errorLogger)

	n := negroni.New()
	r := mux.NewRouter()
	m := metric.NewMiddleware("Product Service")

	n.Use(m)
	// metrics
	r.Path("/metrics").Handler(promhttp.Handler())
	// api
	r.HandleFunc("/api/products", controller.GetProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", controller.GetProduct).Methods("GET")
	r.HandleFunc("/api/products", controller.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", controller.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/products/{id}", controller.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/api/products/buy", controller.BuyProduct).Methods("POST")
	r.HandleFunc("/api/transactions", controller.GetTransactions).Methods("GET")

	n.UseHandler(r)

	n.Run(":8000")

}

func Hello() string {
	return "Hello, world."
}

func initLogger() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		log.Fatal(err)
	}

	infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	warningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

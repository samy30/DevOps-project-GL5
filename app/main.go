package main

import (
	"log"
	"os"

	"github.com/codegangsta/negroni"

	"devopsProjectModule.com/gl5/controllers"
	"devopsProjectModule.com/gl5/logger"
	"devopsProjectModule.com/gl5/metric"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	file, err := os.Open("logs/logs.txt")
	if err != nil {
		log.Fatal("file not found")
	}
	log := logger.NewLogger(file)
	logger.SetDefaultLogger(log)
	logger.Info("Server starting...")
	controller := controllers.NewProductController()

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

// Hello : a test function
func Hello() string {
	return "Hello, world."
}

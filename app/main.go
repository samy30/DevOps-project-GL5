package main

import (
	
	"github.com/codegangsta/negroni"
	
	"devopsProjectModule.com/gl5/metric"
	"devopsProjectModule.com/gl5/controllers"
	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	
	
)

func main() {
	controller := controllers.NewProductController()
	//Init Router
	/*r := mux.NewRouter()
    n := negroni.New()
	metricService := metric.NewMiddleware("serviceName")
	n.Use(metricService)
	
	
	r.Path("/metrics").Handler(promhttp.Handler())
	r.HandleFunc("/ok", func(w http.ResponseWriter, r *http.Request) {
		sleep := rand.Intn(4999) + 1
		time.Sleep(time.Duration(sleep) * time.Millisecond)
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "slept %d milliseconds\n", sleep)
	})

	// arrange our routes
	r.HandleFunc("/api/products", controller.GetProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", controller.GetProduct).Methods("GET")
	r.HandleFunc("/api/products", controller.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", controller.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/products/{id}", controller.DeleteProduct).Methods("DELETE")
	r.HandleFunc("/api/products/buy", controller.BuyProduct).Methods("POST")
	r.HandleFunc("/api/transactions", controller.GetTransactions).Methods("GET")

	// set our port address
	log.Fatal(http.ListenAndServe(":8000", r))*/
	n := negroni.New()
	r := mux.NewRouter()
	m:= metric.NewMiddleware("Product Service")
	
	// if you want to use other buckets than the default (300, 1200, 5000) you can run:
	// m := negroniprometheus.NewMiddleware("serviceName", 400, 1600, 700)

	n.Use(m)

	r.Path("/metrics").Handler(promhttp.Handler())
	//r := http.NewServeMux()
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

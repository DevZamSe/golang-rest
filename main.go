package main

import (
	"log"
	"net/http"

	"microservices/src/middleware"

	"github.com/gorilla/mux"
)

func main() {

	// Init Router
	r := mux.NewRouter()

	// Route Handlers
	r.HandleFunc("/api/products", middleware.GetProducts).Methods("GET")
	r.HandleFunc("/api/products/{id}", middleware.GetProduct).Methods("GET")
	r.HandleFunc("/api/products", middleware.CreateProduct).Methods("POST")
	r.HandleFunc("/api/products/{id}", middleware.UpdateProduct).Methods("PUT")
	r.HandleFunc("/api/products/{id}", middleware.DeleteProduct).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

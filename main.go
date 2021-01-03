package main

import (
	"fmt"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		fmt.Println("default")
		port = "8000"
	}

	fmt.Println("todo listo!!")
	fmt.Println(port)
	// log.Fatal(http.ListenAndServe(":"+port, r))
	http.ListenAndServe(":"+port, r)

}

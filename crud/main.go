package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/riadafridishibly/rest-mux-example/crud/api"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/products", api.GetAllProducts).Methods("GET")
	r.HandleFunc("/products", api.CreateNewProduct).Methods("POST")
	r.HandleFunc("/product/{id}", api.UpdateProduct).Methods("PUT")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

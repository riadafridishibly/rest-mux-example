package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// READ: get all the products
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

// CREATE: create and new product and add to database
func CreateNewProduct(w http.ResponseWriter, r *http.Request) {
	var prod Product
	err := json.NewDecoder(r.Body).Decode(&prod)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	products = append(products, prod)
}

// UPDATE: update a product
func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	var prod Product
	err := json.NewDecoder(r.Body).Decode(&prod)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	id := mux.Vars(r)["id"]

	fmt.Printf("%#v\n", id)
	fmt.Printf("%#v\n", prod)

	for _, product := range products {
		if product.ID == id {
			product.Name = prod.Name
		}
	}
}

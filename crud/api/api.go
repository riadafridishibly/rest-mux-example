package api

import (
	"encoding/json"
	"net/http"
)

// Get requests
func GetAllProducts(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(products)
}

func CreateNewProduct(w http.ResponseWriter, r *http.Request) {
	var prod Product
	err := json.NewDecoder(r.Body).Decode(&prod)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	products = append(products, prod)
}

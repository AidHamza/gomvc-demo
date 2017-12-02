package controllers

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/AidHamza/gomvc-demo/controllers/product"
	"github.com/AidHamza/gomvc-demo/controllers/web"
)

// New .
func New() http.Handler {
	r := mux.NewRouter()

	// website
	r.HandleFunc("/", web.Root).Methods("GET")

	// product
	r.HandleFunc("/product", product.New).Methods("POST")
	r.HandleFunc("/product/{id}", product.Get).Methods("GET")

	return r
}

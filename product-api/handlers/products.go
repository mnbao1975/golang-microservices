package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// KeyProduct is the key used for the Product in the context
type KeyProduct struct{}

// Products defines ...
type Products struct {
	l *log.Logger
}

// NewProducts is pattern of dependency injection, a logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func getProductID(r *http.Request) int {
	// Parse product id from URL
	vars := mux.Vars(r)

	// Convert to int and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
	}

	return id
}

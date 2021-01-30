//Package handlers Product API
//	Documentation Product API.
//
//		Schemes: http, https
//		Host: localhost
//		BasePath: /v2
//		Version: 0.0.1
//
//		Consumes:
//		- application/json
//		- application/xml
//
//		Produces:
//		- application/json
//		- application/xml
//
//	swagger:meta
package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// KeyProduct is the key used for the Product in the context
type KeyProduct struct{}

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

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

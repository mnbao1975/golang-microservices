//Package handlers Product API
//
//	Documentation Product API.
//
//		Schemes: http, https
//		Host: localhost
//		BasePath: /v2
//		Version: 1.0.0
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

	"github.com/mnbao1975/microservices/product-api/data"

	"github.com/gorilla/mux"
)

//	A list of products return in the response
//	swagger:response productsResponse
type productsResponseWrapper struct {
	//	in: body
	Body []data.Product
}

//	A product returns in the response
//	swagger:response productResponse
type productResponseWrapper struct {
	//	in: body
	Body data.Product
}

//	An error returns in the response
//	swagger:response errorResponse
type errorResponseWrapper struct {
	//	in: body
	Body GenericError
}

//	swagger:parameters listOneProduct
type productIDParam struct {
	//	The id of product given on the URI path
	//	in: path
	//	required: true
	ID int `json:"id"`
}

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

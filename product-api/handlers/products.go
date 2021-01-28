package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/mnbao1975/microservices/product-api/data"
)

// Products defines ...
type Products struct {
	l *log.Logger
}

// NewProducts is pattern of dependency injection, a logger
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

// GetProducts return list of products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET products")
	lp := data.GetProducts()

	// Stream the JSON to client, no need to complete parsing it and send to client
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	//rw.Write(d)
}

// AddProduct will add a new product
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST a product")

	// Get KeyProduct from middleware,MiddlewareProductValidation, and cast to data.Product
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

// UpdateProduct will update a product by id
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		p.l.Println("[ERROR] unable to convert id")
		http.Error(rw, "Invalid id", http.StatusBadRequest)
		return
	}
	p.l.Println("PUT a product: ", id)

	// Get KeyProduct from middleware,MiddlewareProductValidation, and cast to data.Product
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	// Other update error
	if err != nil {
		p.l.Println("[ERROR] internal server error")
		http.Error(rw, "Server error", http.StatusInternalServerError)
		return
	}
}

// KeyProduct is
type KeyProduct struct{}

// MiddlewareValidateProduct is used to validate product data
func (p *Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		prod := data.Product{}
		err := prod.FromJSON(r.Body)
		if err != nil {
			p.l.Println("[ERROR] decentializing product")
			http.Error(rw, "Error reading proudct", http.StatusBadRequest)
			return
		}

		// Validate the product
		err = prod.Validate()
		if err != nil {
			p.l.Println("[ERROR] validating product", err)
			http.Error(
				rw,
				fmt.Sprintf("Error validating proudct: %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// Copy the valid product(prod) to the KeyProduct and add it to the context
		// so that the next handdler can access it
		ctx := context.WithValue(r.Context(), KeyProduct{}, prod)
		r = r.WithContext(ctx)

		// Call the next handdler
		next.ServeHTTP(rw, r)
	})
}

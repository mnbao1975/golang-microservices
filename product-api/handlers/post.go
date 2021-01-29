package handlers

import (
	"net/http"

	"github.com/mnbao1975/microservices/product-api/data"
)

// AddProduct will add a new product
func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST a product")

	// Get KeyProduct from middleware,MiddlewareProductValidation, and cast to data.Product
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}

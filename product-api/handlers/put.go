package handlers

import (
	"net/http"

	"github.com/mnbao1975/microservices/product-api/data"
)

// UpdateProduct will update a product by id
func (p *Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Println("PUT a product with id: ", id)

	// Fetch KeyProduct from middleware,MiddlewareValidateProduct, and cast to data.Product
	prod := r.Context().Value(KeyProduct{}).(data.Product)

	err := data.UpdateProduct(id, &prod)
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

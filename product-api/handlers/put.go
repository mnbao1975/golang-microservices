package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/mnbao1975/microservices/product-api/data"
)

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

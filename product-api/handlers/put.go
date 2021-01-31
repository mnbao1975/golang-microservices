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
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(GenericError{Message: "Product not found"}, rw)
		return
	}
	// Other update error
	if err != nil {
		p.l.Println("[ERROR] internal server error")
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(GenericError{Message: "Server error"}, rw)
		return
	}
}

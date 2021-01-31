package handlers

import (
	"net/http"

	"github.com/mnbao1975/microservices/product-api/data"
)

// DeleteProduct returns a product by id
func (p *Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Println("DELETE one product with id: ", id)

	prod, _, err := data.GetOneProduct(id)
	if err == data.ErrProductNotFound {
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(GenericError{Message: "Product not found"}, rw)
	}
	// Stream the JSON to client, no need to complete parsing it and send to client
	err = data.ToJSON(prod, rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

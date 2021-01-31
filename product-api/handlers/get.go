package handlers

import (
	"net/http"

	"github.com/mnbao1975/microservices/product-api/data"
)

// GetProducts return list of products
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET products")
	lp := data.GetProducts()

	// Stream the JSON to client, no need to complete parsing it and send to client
	err := data.ToJSON(lp, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(GenericError{Message: err.Error()}, rw)
		return
	}
}

//	swagger:route GET /products products listProducts
//	Returns a list of products

// GetOneProduct returns a product by id
func (p *Products) GetOneProduct(rw http.ResponseWriter, r *http.Request) {
	id := getProductID(r)
	p.l.Println("GET one product with id: ", id)

	prod, _, err := data.GetOneProduct(id)
	if err == data.ErrProductNotFound {
		rw.WriteHeader(http.StatusNotFound)
		data.ToJSON(GenericError{Message: "Product not found"}, rw)
		return
	}
	// Stream the JSON to client, no need to complete parsing it and send to client
	err = data.ToJSON(prod, rw)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		data.ToJSON(GenericError{Message: err.Error()}, rw)
		return
	}
}

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
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}
}

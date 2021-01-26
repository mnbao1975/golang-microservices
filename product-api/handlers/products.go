package handlers

import (
	"log"
	"net/http"

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

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
	}
	// Catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

// GET method
func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("GET products")
	lp := data.GetProducts()
	//d, err := json.Marshal(lp)

	// Stream the JSON to client, no need to complete parsing it and send to client
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
		return
	}

	//rw.Write(d)
}

// POST method
func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("POST a product")

	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	p.l.Printf("Prod: %#v", prod)
}

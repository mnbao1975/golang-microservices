package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

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

	if r.Method == http.MethodPut {
		p.l.Println(r.RequestURI)
		// Expect the id in the URI
		reg := regexp.MustCompile(`/([0-9]+)`)
		grp := reg.FindAllStringSubmatch(r.RequestURI, -1)

		if len(grp) != 1 {
			p.l.Println("Invalid URI more that one id")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		if len(grp[0]) != 2 {
			p.l.Println("Invalid URI more that one capture group")
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := grp[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		p.l.Println("PUT id: ", id)
		p.updateProduct(id, rw, r)
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

	//p.l.Printf("Prod: %#v", prod)
	data.AddProduct(prod)
}

// PUT method
func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("PUT a product")
	prod := &data.Product{}

	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, prod)
	if err == data.ErrProductNotFound {
		http.Error(rw, "Product not found", http.StatusNotFound)
		return
	}
	// Other update error
	if err != nil {
		http.Error(rw, "Server error", http.StatusInternalServerError)
		return
	}
}

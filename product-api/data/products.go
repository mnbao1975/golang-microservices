package data

import (
	"encoding/json"
	"io"
	"time"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"` //change name of ID to id
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float32 `json:"price"`
	SKU         string  `json:"sku"`
	CreatedOn   string  `json:"-"` // ignore it
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// FromJSON read POST data from request and parse it and assign it to Product
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// Products type is a custom type and a collection of Product
// We can create a func for it
type Products []*Product

// ToJSON serializes the contents of the products collection to JSON
// NewEncoder provides better performance than json.Marshal() as it does not buffer
// the output into memory
func (p *Products) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

// GetProducts will return products list
func GetProducts() Products {
	return productList
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc232",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String()},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and stong coffee without milk",
		Price:       1.99,
		SKU:         "fgd324",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String()}}

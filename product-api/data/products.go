package data

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// Product defines the structure for an API product
type Product struct {
	ID          int     `json:"id"` //change name of ID to id
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"` // ignore it
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// FromJSON read POST data from request and parse it and assign it to Product
func (p *Product) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(p)
}

// Validate will validate the json data of product
func (p *Product) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("sku", validateSKU)
	return validate.Struct(p)
}

func validateSKU(fl validator.FieldLevel) bool {
	// SKU is of format abc-cfdf-dfsgh
	reg := regexp.MustCompile(`[a-z]+-[a-z]+-[a-z]+`)
	matches := reg.FindAllString(fl.Field().String(), -1)
	fmt.Println(matches)
	if len(matches) != 1 {
		return false
	}

	return true
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

// AddProduct will add a new product submited from client
func AddProduct(p *Product) {
	p.ID = getNextID()
	productList = append(productList, p)
}

func getNextID() int {
	lp := productList[len(productList)-1]
	return lp.ID + 1
}

// UpdateProduct a product
func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}

	p.ID = id
	productList[pos] = p

	return nil
}

// ErrProductNotFound ..
var ErrProductNotFound = fmt.Errorf("Product not found")

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

var productList = []*Product{
	&Product{
		ID:          1,
		Name:        "Latte",
		Description: "Frothy milky coffee",
		Price:       2.45,
		SKU:         "abc232",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
	&Product{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and stong coffee without milk",
		Price:       1.99,
		SKU:         "fgd324",
		CreatedOn:   time.Now().UTC().String(),
		UpdatedOn:   time.Now().UTC().String(),
	},
}

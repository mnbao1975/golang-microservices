package data

import (
	"fmt"
	"regexp"
	"time"

	"github.com/go-playground/validator"
)

// ErrProductNotFound ..
var ErrProductNotFound = fmt.Errorf("Product not found")

//Product defines the structure for an API product
//	swagger:model
type Product struct {
	// the id for this product
	//
	// required: true
	// min: 1
	ID          int     `json:"id"` //change name of ID to id
	Name        string  `json:"name" validate:"required"`
	Description string  `json:"description"`
	Price       float32 `json:"price" validate:"gt=0"`
	SKU         string  `json:"sku" validate:"required,sku"`
	CreatedOn   string  `json:"-"` // ignore it
	UpdatedOn   string  `json:"-"`
	DeletedOn   string  `json:"-"`
}

// Products type is a custom type and a collection of Product
// We can create a func for it
type Products []*Product

// FromJSON read POST data from request and parse it and assign it to Product
// func (p *Product) FromJSON(r io.Reader) error {
// 	e := json.NewDecoder(r)
// 	return e.Decode(p)
// }

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

// ToJSON serializes the contents of the products collection to JSON
// NewEncoder provides better performance than json.Marshal() as it does not buffer
// the output into memory
// func (p *Products) ToJSON(w io.Writer) error {
// 	e := json.NewEncoder(w)
// 	return e.Encode(p)
// }

// // ToJSONOneProduct serialized the content of one product
// func (p *Product) ToJSONOneProduct(w io.Writer) error {
// 	e := json.NewEncoder(w)
// 	return e.Encode(p)
// }

// GetProducts returns products list
func GetProducts() Products {
	return productList
}

// GetOneProduct returns one product by id
func GetOneProduct(id int) (*Product, int, error) {
	return findProduct(id)
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

// DeleteProduct a product
func DeleteProduct(id int) error {
	_, i, _ := findProduct(id)
	if i == -1 {
		return ErrProductNotFound
	}

	productList = append(productList[:i], productList[i+1])

	return nil
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range productList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProductNotFound
}

// Fake dataset
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

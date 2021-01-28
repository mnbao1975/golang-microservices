package data

import (
	"testing"
)

func TestCheckValidation(t *testing.T) {
	p := &Product{
		Name:  "item",
		Price: 1.9,
		SKU:   "acv-abc-aff",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}

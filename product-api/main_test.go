package main

import (
	"fmt"
	"testing"

	"github.com/mnbao1975/microservices/product-api/sdk/client/products"

	"github.com/mnbao1975/microservices/product-api/sdk/client"
)

func TestOurClient(t *testing.T) {
	//cfg := client.DefaultTransportConfig().WithHost("localhost:3000")
	cfg := client.DefaultTransportConfig()
	cfg.Host = "localhost:3000"

	c := client.NewHTTPClientWithConfig(nil, cfg)
	// params := products.NewListProductsParams()
	// prod, err := c.Products.ListProducts(params)

	// products/2
	params := products.NewListOneProductParams().WithID(2)
	prod, err := c.Products.ListOneProduct(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(prod)
}

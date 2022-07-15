package main

import (
	"testing"

	"github.com/microservices/sdk/client/products"

	"github.com/microservices/sdk/client"
)

func TestOurClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)

	params := products.NewListProductsParams()
	_, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

}

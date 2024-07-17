package bigcartel

import (
	"encoding/json"

	"github.com/Sn1perdog/bigcartel-go-api/types"
)

// GetProducts retrieves products from BigCartel
func (c *Client) GetProducts() ([]types.Product, error) {
	url := "/products"
	respBody, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var products []types.Product
	if err := json.Unmarshal(respBody, &products); err != nil {
		return nil, err
	}

	return products, nil
}

// CreateProduct creates a new product in BigCartel
func (c *Client) CreateProduct(product types.Product) (*types.Product, error) {
	url := "/products"
	respBody, err := c.doRequest("POST", url, product)
	if err != nil {
		return nil, err
	}

	var newProduct types.Product
	if err := json.Unmarshal(respBody, &newProduct); err != nil {
		return nil, err
	}

	return &newProduct, nil
}

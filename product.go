package bigcartel

import (
	"encoding/json"
	"fmt"

	"bigcartel/types"
)

// GetProducts retrieves products from BigCartel
func (c *Client) GetProducts() ([]types.Product, error) {
	url := fmt.Sprintf("%s/products", c.BaseURL)
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
	url := fmt.Sprintf("%s/products", c.BaseURL)
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

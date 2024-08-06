package bigcartel

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Sn1perdog/bigcartel-go-api/types"
)

// GetProducts retrieves products from BigCartel
func (c *Client) GetProducts() ([]types.Product, error) {
	url := "/products"
	respBody, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GET request to %s: %w", url, err)
	}

	var productResponse types.ProductResponse
	if err := json.Unmarshal(respBody, &productResponse); err != nil {
		log.Printf("Failed to unmarshal response: %s\n", string(respBody))
		return nil, fmt.Errorf("failed to unmarshal response from %s: %w", url, err)
	}

	products := make([]types.Product, len(productResponse.Data))
	for i, p := range productResponse.Data {
		products[i] = types.Product{
			ID:              p.ID,
			Name:            p.Attributes.Name,
			Permalink:       p.Attributes.Permalink,
			Status:          p.Attributes.Status,
			Description:     p.Attributes.Description,
			CreatedAt:       p.Attributes.CreatedAt,
			UpdatedAt:       p.Attributes.UpdatedAt,
			DefaultPrice:    p.Attributes.DefaultPrice,
			OnSale:          p.Attributes.OnSale,
			Position:        p.Attributes.Position,
			URL:             p.Attributes.URL,
			PrimaryImageURL: p.Attributes.PrimaryImageURL,
		}
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

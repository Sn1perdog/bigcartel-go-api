package types

import "time"

type Product struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	Permalink       string    `json:"permalink"`
	Status          string    `json:"status"`
	Description     string    `json:"description"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DefaultPrice    string    `json:"default_price"`
	OnSale          bool      `json:"on_sale"`
	Position        int       `json:"position"`
	URL             string    `json:"url"`
	PrimaryImageURL string    `json:"primary_image_url"`
}

type ProductResponse struct {
	Data []ProductData `json:"data"`
}

type ProductData struct {
	ID         string            `json:"id"`
	Type       string            `json:"type"`
	Attributes ProductAttributes `json:"attributes"`
}

type ProductAttributes struct {
	Name            string    `json:"name"`
	Permalink       string    `json:"permalink"`
	Status          string    `json:"status"`
	Description     string    `json:"description"`
	CategoryNames   []string  `json:"category_names"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	DefaultPrice    string    `json:"default_price"`
	OnSale          bool      `json:"on_sale"`
	Position        int       `json:"position"`
	URL             string    `json:"url"`
	HasOptionGroups bool      `json:"has_option_groups"`
	PrimaryImageURL string    `json:"primary_image_url"`
}

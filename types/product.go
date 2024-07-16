package types

import "time"

type Product struct {
	ID              string           `json:"id"`
	Name            string           `json:"name"`
	Permalink       string           `json:"permalink"`
	Status          string           `json:"status"`
	Description     string           `json:"description"`
	CreatedAt       time.Time        `json:"created_at"`
	UpdatedAt       time.Time        `json:"updated_at"`
	DefaultPrice    float64          `json:"default_price"`
	OnSale          bool             `json:"on_sale"`
	Position        int              `json:"position"`
	URL             string           `json:"url"`
	PrimaryImageURL string           `json:"primary_image_url"`
	Options         []ProductOption  `json:"options"`
	Artists         []Artist         `json:"artists"`
	Categories      []Category       `json:"categories"`
	ShippingOptions []ShippingOption `json:"shipping_options"`
	Images          []Image          `json:"images"`
}

type ProductOption struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
	Sold     int     `json:"sold"`
}

type Artist struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Permalink string `json:"permalink"`
	Position  string `json:"position"`
}

type Category struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Permalink string `json:"permalink"`
	Position  int    `json:"position"`
}

type ShippingOption struct {
	ID              int             `json:"id"`
	PriceAlone      float64         `json:"price_alone"`
	PriceWithOthers float64         `json:"price_with_others"`
	Country         ShippingCountry `json:"country"`
}

type ShippingCountry struct {
	ID string `json:"id"`
}

type Image struct {
	ID  int    `json:"id"`
	URL string `json:"url"`
}

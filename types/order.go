package types

import "time"

type Order struct {
	ID                         string            `json:"id"`
	ItemCount                  int               `json:"item_count"`
	ItemTotal                  float64           `json:"item_total"`
	DiscountTotal              float64           `json:"discount_total"`
	ShippingTotal              float64           `json:"shipping_total"`
	TaxTotal                   float64           `json:"tax_total"`
	Total                      float64           `json:"total"`
	CustomerFirstName          string            `json:"customer_first_name"`
	CustomerLastName           string            `json:"customer_last_name"`
	CustomerEmail              string            `json:"customer_email"`
	CustomerPhoneNumber        string            `json:"customer_phone_number"`
	CustomerOptedInToMarketing bool              `json:"customer_opted_in_to_marketing"`
	CustomerNote               string            `json:"customer_note"`
	ShippingAddress1           string            `json:"shipping_address_1"`
	ShippingAddress2           string            `json:"shipping_address_2"`
	ShippingCity               string            `json:"shipping_city"`
	ShippingState              string            `json:"shipping_state"`
	ShippingZip                string            `json:"shipping_zip"`
	ShippingCountryID          string            `json:"shipping_country_id"`
	ShippingCountryName        string            `json:"shipping_country_name"`
	ShippingLatitude           int               `json:"shipping_latitude"`
	ShippingLongitude          int               `json:"shipping_longitude"`
	ShippingStatus             string            `json:"shipping_status"`
	PaymentStatus              string            `json:"payment_status"`
	CreatedAt                  time.Time         `json:"created_at"`
	UpdatedAt                  time.Time         `json:"updated_at"`
	CompletedAt                time.Time         `json:"completed_at"`
	Currency                   Currency          `json:"currency"`
	Events                     []OrderEvent      `json:"events"`
	Items                      []OrderItem       `json:"items"`
	Transactions               []Transaction     `json:"transactions"`
	Adjustments                []OrderAdjustment `json:"adjustments"`
}

type Currency struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Sign   string `json:"sign"`
	Locale string `json:"locale"`
}

type OrderEvent struct {
	ID        int       `json:"id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type OrderItem struct {
	ID                int     `json:"id"`
	ProductName       string  `json:"product_name"`
	ProductOptionName string  `json:"product_option_name"`
	Quantity          int     `json:"quantity"`
	Price             float64 `json:"price"`
	Total             float64 `json:"total"`
	ImageURL          string  `json:"image_url"`
	ProductID         string  `json:"product_id"`
	ProductOptionID   string  `json:"product_option_id"`
}

type Transaction struct {
	ID           int      `json:"id"`
	Label        string   `json:"label"`
	Amount       float64  `json:"amount"`
	Processor    string   `json:"processor"`
	ProcessorID  string   `json:"processor_id"`
	ProcessorURL string   `json:"processor_url"`
	Currency     Currency `json:"currency"`
}

type OrderAdjustment struct {
	ID     int     `json:"id"`
	Amount float64 `json:"amount"`
	Label  string  `json:"label"`
}

package types

import (
	"encoding/json"
	"time"
)

// OrderResponse represents the full JSON response for orders
type OrdersResponse struct {
	Data     []OrderData    `json:"data"`
	Included []IncludedData `json:"included"`
}
type OrderResponse struct {
	Data     OrderData      `json:"data"`
	Included []IncludedData `json:"included"`
}

// OrderData represents each order entry in the response
type OrderData struct {
	ID            string             `json:"id"`
	Type          string             `json:"type"`
	Attributes    OrderAttributes    `json:"attributes"`
	Links         OrderLinks         `json:"links"`
	Relationships OrderRelationships `json:"relationships"`
}

// OrderAttributes holds the attributes of an order
type OrderAttributes struct {
	ItemCount                  int       `json:"item_count"`
	ItemTotal                  string    `json:"item_total"`
	DiscountTotal              string    `json:"discount_total"`
	ShippingTotal              string    `json:"shipping_total"`
	TaxTotal                   string    `json:"tax_total"`
	Total                      string    `json:"total"`
	TaxRemittedByBigCartel     *string   `json:"tax_remitted_by_big_cartel"`
	CustomerFirstName          string    `json:"customer_first_name"`
	CustomerLastName           string    `json:"customer_last_name"`
	CustomerEmail              string    `json:"customer_email"`
	CustomerPhoneNumber        *string   `json:"customer_phone_number"`
	CustomerOptedInToMarketing bool      `json:"customer_opted_in_to_marketing"`
	CustomerNote               string    `json:"customer_note"`
	ShippingAddress1           string    `json:"shipping_address_1"`
	ShippingAddress2           *string   `json:"shipping_address_2"`
	ShippingCity               string    `json:"shipping_city"`
	ShippingState              string    `json:"shipping_state"`
	ShippingZip                string    `json:"shipping_zip"`
	ShippingCountryID          string    `json:"shipping_country_id"`
	ShippingCountryName        string    `json:"shipping_country_name"`
	ShippingStatus             string    `json:"shipping_status"`
	ShippingLatitude           *float64  `json:"shipping_latitude"`
	ShippingLongitude          *float64  `json:"shipping_longitude"`
	CompletedAt                time.Time `json:"completed_at"`
	UpdatedAt                  time.Time `json:"updated_at"`
	CreatedAt                  time.Time `json:"created_at"`
	PaymentStatus              string    `json:"payment_status"`
}

// OrderLinks represents the links section of the order data
type OrderLinks struct {
	Self string `json:"self"`
}

// OrderRelationships captures all the relationships within an order
type OrderRelationships struct {
	Currency        OrderRelationshipData `json:"currency"`
	ShippingCountry OrderRelationshipData `json:"shipping_country"`
	Events          OrderRelationshipList `json:"events"`
	Warnings        OrderRelationshipList `json:"warnings"`
	Items           OrderRelationshipList `json:"items"`
	Transactions    OrderRelationshipList `json:"transactions"`
	Adjustments     OrderRelationshipList `json:"adjustments"`
}

// OrderRelationshipData represents a single relationship with a type and ID
type OrderRelationshipData struct {
	Data RelationshipDetail `json:"data"`
}

// OrderRelationshipList represents a list of relationships
type OrderRelationshipList struct {
	Data []RelationshipDetail `json:"data"`
}

// RelationshipDetail provides the type and ID of related entities
type RelationshipDetail struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

// IncludedData represents the included section for detailed information
type IncludedData struct {
	ID         string          `json:"id"`
	Type       string          `json:"type"`
	Attributes json.RawMessage `json:"attributes"` // Use RawMessage to defer parsing
}

// OrderLineItem represents detailed attributes of an order item
type OrderLineItem struct {
	ProductName       string   `json:"product_name"`
	ProductID         string   `json:"product_id"`
	Quantity          int      `json:"quantity"`
	Price             string   `json:"price"`
	TaxTotal          string   `json:"tax_total"`
	Total             string   `json:"total"`
	QuantityShipped   int      `json:"quantity_shipped"`
	QuantityUnshipped int      `json:"quantity_unshipped"`
	CategoryNames     []string `json:"category_names"`
}

// OrderDetail represents detailed information for an order including products
// Returned from the GetOrders function
type OrderDetail struct {
	OrderID           string          `json:"order_id"`
	CustomerFirstName string          `json:"customer_first_name"`
	CustomerLastName  string          `json:"customer_last_name"`
	CustomerEmail     string          `json:"customer_email"`
	ShippingAddress1  string          `json:"shipping_address_1"`
	ShippingAddress2  string          `json:"shipping_address_2"`
	ShippingCity      string          `json:"shipping_city"`
	ShippingState     string          `json:"shipping_state"`
	ShippingZip       string          `json:"shipping_zip"`
	ShippingCountry   string          `json:"shipping_country"`
	PaymentStatus     string          `json:"payment_status"`
	ShippingStatus    string          `json:"shipment_status"`
	CreatedAt         time.Time       `json:"created_at"`
	ProductsPurchased []ProductDetail `json:"products_purchased"`
}

// ProductDetail represents details of a product purchased in an order
// Used in the GetProducts response
type ProductDetail struct {
	ProductName string `json:"product_name"`
	ProductID   string `json:"product_id"`
	Quantity    int    `json:"quantity"`
	Price       string `json:"price"`
	Total       string `json:"total"`
}

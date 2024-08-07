package types

type OrderUpdateRequest struct {
	Data OrderUpdateData `json:"data"`
}

type OrderUpdateData struct {
	ID         string                `json:"id"`
	Type       string                `json:"type"`
	Attributes OrderUpdateAttributes `json:"attributes"`
}

type OrderUpdateAttributes struct {
	CustomerFirstName   *string `json:"customer_first_name,omitempty"`
	CustomerLastName    *string `json:"customer_last_name,omitempty"`
	CustomerEmail       *string `json:"customer_email,omitempty"`
	CustomerPhoneNumber *string `json:"customer_phone_number,omitempty"`
	CustomerNote        *string `json:"customer_note,omitempty"`
	ShippingAddress1    *string `json:"shipping_address_1,omitempty"`
	ShippingAddress2    *string `json:"shipping_address_2,omitempty"`
	ShippingCity        *string `json:"shipping_city,omitempty"`
	ShippingState       *string `json:"shipping_state,omitempty"`
	ShippingZip         *string `json:"shipping_zip,omitempty"`
	ShippingCountryID   *string `json:"shipping_country_id,omitempty"`
	ShippingStatus      *string `json:"shipping_status,omitempty"`
}

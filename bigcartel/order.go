package bigcartel

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Sn1perdog/bigcartel-go-api/types"
)

/*
GetOrders retrieves all orders from BigCartel
Includes automatically mapped product information for each order

	Search: Finds orders that match the customer name (customer_first_name + customer_last_name), customer_email, or id attributes.
		Example: `myemail@proton.com`
	Filter: Filter orders using the shipping_status, completed_at_from, completed_at_to, updated_at_from and updated_at_to.
		Example: `[shipping_status]=shipped`
	Sort: Sorts orders using the specified attribute, in ascending order by default. Prefix the value with - to specify descending order. Allowed values are completed_at, created_at, and updated_at.
		Example: `-created_at`
*/
func (c *Client) GetOrders(search, filter, sort string) ([]types.OrderDetail, error) {
	url := "/orders"

	if search != "" || filter != "" || sort != "" {
		q := url + "?"
		if search != "" {
			q += "search=" + search + "&"
		}
		if filter != "" {
			q += "filter" + filter + "&"
		}
		if sort != "" {
			q += "sort=" + sort
		}
		url = q
	}

	respBody, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GET request to %s: %w", url, err)
	}

	var orderResponse types.OrderResponse
	if err := json.Unmarshal(respBody, &orderResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal order response: %w", err)
	}

	// Process included data to extract detailed product information
	orderItems := make(map[string]types.OrderLineItem)
	for _, include := range orderResponse.Included {
		if include.Type == "order_line_items" {
			var item types.OrderLineItem
			if err := json.Unmarshal(include.Attributes, &item); err != nil {
				log.Printf("Failed to unmarshal order line item: %v", err)
				continue
			}
			orderItems[include.ID] = item
		}
	}

	var orderDetails []types.OrderDetail
	for _, order := range orderResponse.Data {
		var productsPurchased []types.ProductDetail

		for _, itemRel := range order.Relationships.Items.Data {
			if item, found := orderItems[itemRel.ID]; found {
				productDetail := types.ProductDetail{
					ProductName: item.ProductName,
					ProductID:   item.ProductID,
					Quantity:    item.Quantity,
					Price:       item.Price,
					Total:       item.Total,
				}
				productsPurchased = append(productsPurchased, productDetail)
			}
		}

		shippingAddress2 := ""
		if order.Attributes.ShippingAddress2 != nil {
			shippingAddress2 = *order.Attributes.ShippingAddress2
		}

		orderDetail := types.OrderDetail{
			OrderID:           order.ID,
			CustomerFirstName: order.Attributes.CustomerFirstName,
			CustomerLastName:  order.Attributes.CustomerLastName,
			CustomerEmail:     order.Attributes.CustomerEmail,
			ShippingAddress1:  order.Attributes.ShippingAddress1,
			ShippingAddress2:  shippingAddress2,
			ShippingCity:      order.Attributes.ShippingCity,
			ShippingState:     order.Attributes.ShippingState,
			ShippingZip:       order.Attributes.ShippingZip,
			ShippingCountry:   order.Attributes.ShippingCountryName,
			PaymentStatus:     order.Attributes.PaymentStatus,
			ShippingStatus:    order.Attributes.ShippingStatus,
			CreatedAt:         order.Attributes.CreatedAt,
			ProductsPurchased: productsPurchased,
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	return orderDetails, nil
}

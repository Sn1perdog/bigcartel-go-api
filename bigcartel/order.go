package bigcartel

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/Sn1perdog/bigcartel-go-api/types"
)

// Processes and maps order(s) data to OrderDetail
func mapOrderDetails(orderData []types.OrderData, included []types.IncludedData) ([]types.OrderDetail, error) {
	orderItems := make(map[string]types.OrderLineItem)
	for _, include := range included {
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
	for _, order := range orderData {
		productsPurchased := extractProductDetails(order.Relationships.Items.Data, orderItems)

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

// extracts product information from relationships and orderItems
func extractProductDetails(itemRelationships []types.RelationshipDetail, orderItems map[string]types.OrderLineItem) []types.ProductDetail {
	var productsPurchased []types.ProductDetail
	for _, itemRel := range itemRelationships {
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
	return productsPurchased
}

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

	var ordersResponse types.OrdersResponse
	if err := json.Unmarshal(respBody, &ordersResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal order response: %w", err)
	}

	orderDetails, err := mapOrderDetails(ordersResponse.Data, ordersResponse.Included)
	if err != nil {
		return nil, fmt.Errorf("failed to map order details: %w", err)
	}

	return orderDetails, nil
}

// UpdateOrder updates an existing order in BigCartel
func (c *Client) UpdateOrder(orderID string, updateData types.OrderUpdateAttributes) (*types.OrderDetail, error) {
	url := fmt.Sprintf("/orders/%s", orderID)

	payload := types.OrderUpdateRequest{
		Data: types.OrderUpdateData{
			ID:         orderID,
			Type:       "orders",
			Attributes: updateData,
		},
	}

	respBody, err := c.doRequest("PATCH", url, payload)
	if err != nil {
		return nil, fmt.Errorf("failed to perform PATCH request to %s: %w", url, err)
	}

	var updatedOrderResponse types.OrderResponse
	if err := json.Unmarshal(respBody, &updatedOrderResponse); err != nil {
		return nil, fmt.Errorf("failed to unmarshal updated order response: %w", err)
	}

	orderDetails, err := mapOrderDetails([]types.OrderData{updatedOrderResponse.Data}, updatedOrderResponse.Included)
	if err != nil {
		return nil, fmt.Errorf("failed to map updated order details: %w", err)
	}

	if len(orderDetails) > 0 {
		return &orderDetails[0], nil
	}

	return nil, fmt.Errorf("no order details found after update")
}

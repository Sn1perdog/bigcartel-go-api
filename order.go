package bigcartel

import (
	"encoding/json"
	"fmt"

	"bigcartel/types"
)

/*
GetOrders retrieves all orders from BigCartel

	Search: Finds orders that match the customer name (customer_first_name + customer_last_name), customer_email, or id attributes.
		Example: `myemail@proton.com`
	Filter: Filter orders using the shipping_status, completed_at_from, completed_at_to, updated_at_from and updated_at_to.
		Example: `[shipping_status]=shipped`
	Sort: Sorts orders using the specified attribute, in ascending order by default. Prefix the value with - to specify descending order. Allowed values are completed_at, created_at, and updated_at.
		Example: `-created_at`
*/
func (c *Client) GetOrders(search, filter, sort string) ([]types.Order, error) {
	url := fmt.Sprintf("%s/orders", c.BaseURL)

	if search != "" || filter != "" || sort != "" {
		q := url + "?"
		if search != "" {
			q += "search=" + search + "&"
		}
		if filter != "" {
			q += "filter=" + filter + "&"
		}
		if sort != "" {
			q += "sort=" + sort
		}
		url = q
	}

	respBody, err := c.doRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var orders []types.Order
	if err := json.Unmarshal(respBody, &orders); err != nil {
		return nil, err
	}

	return orders, nil
}

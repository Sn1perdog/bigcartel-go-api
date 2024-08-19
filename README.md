# Big Cartel Go API
This Go library provides an easy-to-use interface for interacting with the Big Cartel API, allowing you to manage products, orders, and more within your Big Cartel store programmatically.

[Bigcartel API docs](https://developers.bigcartel.com/api/v1)
# Example usage
Using .env file to store credentials information, we create a Bigcartel client through which we can perform requests and view order information.
These are required to perform the API call according to the docs.

Do not upload your credentials anywhere public.
## Example .env file
```
BIGCARTEL_ACCOUNT_NUMBER=1234567
BIGCARTEL_USER_AGENT_1=Company info
BIGCARTEL_USER_AGENT_2=Personal info
BIGCARTEL_USERNAME=username
BIGCARTEL_PASSWORD=password
```

## Example implementation
``` Go
package main

import (
	"fmt"
	"log"

	"github.com/Sn1perdog/bigcartel-go-api/bigcartel"
	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

type Config struct {
	BIGCARTEL_ACCOUNT_NUMBER string `env:"BIGCARTEL_ACCOUNT_NUMBER,required"`
	BIGCARTEL_USER_AGENT_1   string `env:"BIGCARTEL_USER_AGENT_1,required"`
	BIGCARTEL_USER_AGENT_2   string `env:"BIGCARTEL_USER_AGENT_2,required"`
	BIGCARTEL_USERNAME       string `env:"BIGCARTEL_USERNAME,required"`
	BIGCARTEL_PASSWORD       string `env:"BIGCARTEL_PASSWORD,required"`
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env file: %e", err)
	}

	cfg := Config{}       // 👈 new instance of `Config`
	err = env.Parse(&cfg) // 👈 Parse environment variables into `Config`
	if err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}

	accountNumber := cfg.BIGCARTEL_ACCOUNT_NUMBER
	userAgent1 := cfg.BIGCARTEL_USER_AGENT_1
	userAgent2 := cfg.BIGCARTEL_USER_AGENT_2
	username := cfg.BIGCARTEL_USERNAME
	password := cfg.BIGCARTEL_PASSWORD

	if accountNumber == "" || userAgent1 == "" || userAgent2 == "" || username == "" || password == "" {
		log.Fatalf("Missing required environment variables")
	}

	client := bigcartel.NewClient(accountNumber, userAgent1, userAgent2, username, password)

	// Get Products
	products, err := client.GetProducts()
	if err != nil {
		log.Fatalf("Error fetching products: %v", err)
	}

	for _, p := range products {
		fmt.Printf("Product ID: %s, Name: %s, Price: %s\n", p.ID, p.Name, p.DefaultPrice)
	}

	// Get all shipped Orders
	orders, err := client.GetOrders("", "[shipping_status]=shipped", "")
	if err != nil {
		log.Fatalf("Error fetching orders: %v", err)
	}

	for _, o := range orders {
		fmt.Printf("Order ID: %s, Customer: %s %s, Item ordered: %s, Shipment status: %s\n", o.OrderID, o.CustomerFirstName, o.CustomerLastName, o.ProductsPurchased[0].ProductName, o.ShippingStatus)
	}
 
	// order id used for update and get
	orderID := "SLFZ-875364"
	// Update Order, returns order with updated fields
	newNote := "Customer likes french fries"
	updateData := types.OrderUpdateAttributes{
		ShippingAddress1: &newNote,
	}
	updatedOrder, err := client.UpdateOrder(orderID, updateData)
	if err != nil {
		log.Fatalf("Error updating order: %v", err)
	}
	fmt.Printf("Updated Order ID: %s, New note: %s\n", updatedOrder.OrderID, updatedOrder.CustomerNote)

	// Get single Order
	order, err := client.GetOrder(orderID)
	if err != nil {
		log.Fatalf("Error fetching order: %v", err)
	}
	fmt.Printf("Order ID: %s, Customer: %s %s, Item ordered: %s, Shipment status: %s\n", order.OrderID, order.CustomerFirstName, order.CustomerLastName, order.ProductsPurchased[0].ProductName, order.ShippingStatus)
}

```

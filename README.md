# Big Cartel Go API
Still WIP, Products and Orders with their respective Products can be retrieved.
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

	// Get Orders, includes the products ordered in the []ProductsPurchased
	// See GetOrders docs for options
	orders, err := client.GetOrders("", "[shipping_status]=shipped", "")
	if err != nil {
		log.Fatalf("Error fetching orders: %v", err)
	}

	for _, o := range orders {
		fmt.Printf("Order ID: %s, Customer: %s %s, Item ordered: %s, Shipment status: %s\n", o.OrderID, o.CustomerFirstName, o.CustomerLastName, o.ProductsPurchased[0].ProductName, o.ShippingStatus)
	}
}
```

package bigcartel

import (
	"log"
	"net/http"
	"os"
)

type Client struct {
	APIToken   string
	HTTPClient *http.Client
	BaseURL    string
	UserAgent1 string
	UserAgent2 string
	Username   string
	Password   string
}

func NewClient() *Client {
	bigcartelUserAgent1 := os.Getenv("BIGCARTEL_USER_AGENT_1")
	bigcartelUserAgent2 := os.Getenv("BIGCARTEL_USER_AGENT_2")
	bigcartelAccountNumber := os.Getenv("BIGCARTEL_ACCOUNT_NUMBER")
	bigcartelUsername := os.Getenv("BIGCARTEL_USERNAME")
	bigcartelPassword := os.Getenv("BIGCARTEL_PASSWORD")

	if bigcartelUserAgent1 == "" || bigcartelUserAgent2 == "" || bigcartelAccountNumber == "" || bigcartelUsername == "" || bigcartelPassword == "" {
		log.Fatalf("Missing required environment variables")
	}

	return &Client{
		HTTPClient: &http.Client{},
		BaseURL:    "https://api.bigcartel.com/v1/accounts/" + bigcartelAccountNumber,
		UserAgent1: bigcartelUserAgent1,
		UserAgent2: bigcartelUserAgent2,
		Username:   bigcartelUsername,
		Password:   bigcartelPassword,
	}
}

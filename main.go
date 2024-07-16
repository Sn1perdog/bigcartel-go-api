package bigcartel

import (
	"fmt"
	"net/http"
	"time"
)

const apiBaseURL = "https://api.bigcartel.com/v1/accounts"

type Client struct {
	APIToken   string
	HTTPClient *http.Client
	BaseURL    string
}

func NewClient(apiToken string, storeId string) *Client {
	return &Client{
		APIToken:   apiToken,
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		BaseURL:    fmt.Sprintf("%s/%s", apiBaseURL, storeId),
	}
}

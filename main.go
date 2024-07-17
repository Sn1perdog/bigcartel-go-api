package bigcartel

import (
	"net/http"
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

func NewClient(apiToken, accountNumber, userAgent1, userAgent2, username, password string) *Client {
	return &Client{
		APIToken:   apiToken,
		HTTPClient: &http.Client{},
		BaseURL:    "https://api.bigcartel.com/v1/accounts/" + accountNumber,
		UserAgent1: userAgent1,
		UserAgent2: userAgent2,
		Username:   username,
		Password:   password,
	}
}

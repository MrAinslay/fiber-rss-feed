package api

import (
	"net/http"
	"time"
)

type Client struct {
	apiKey     string
	httpClient http.Client
}

func NewClient(timeout time.Duration) Client {
	return Client{
		apiKey: "",
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

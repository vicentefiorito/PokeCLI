// creates the client to make http requests
package pokeapi

import (
	"net/http"
	"time"
)


type Client struct{
	httpClient http.Client
}

// creates a new client
func NewClient() Client{
	return Client{
		httpClient: http.Client{
			Timeout: time.Minute,
		},
	}
}
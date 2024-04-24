// creates the client to make http requests
package pokeapi

import "net/http"


type Client struct{
	httpClient http.Client
}

// creates a new client

// creates the client to make http requests
package pokeapi

import (
	"net/http"
	"time"

	"github.com/vicentefiorito/pokeCLI/internal/pokecache"
)

type Client struct {
	cache      pokecache.Cache
	httpClient http.Client
}

// creates a new client
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

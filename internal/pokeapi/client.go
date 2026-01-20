package pokeapi

import (
	"net/http"
	"time"
	"github.com/jwoodsiii/pokedex/internal/pokecache"
)

// Client -
type Client struct {
	httpClient 	http.Client
	pokeCache	*pokecache.Cache
}

// NewClient -
func NewClient(timeout time.Duration, cache *pokecache.Cache) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: cache,
	}
}

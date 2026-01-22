package pokeapi

import (
	"net/http"
	"time"

	"github.com/jwoodsiii/pokedex/internal/pokecache"
)

// Client -
type Client struct {
	httpClient http.Client
	pokeCache  *pokecache.Cache
	Pokedex    Pokedex
}

// NewClient -
func NewClient(timeout time.Duration, cache *pokecache.Cache, pokedex Pokedex) Client {
	return Client{
		httpClient: http.Client{
			Timeout: timeout,
		},
		pokeCache: cache,
		Pokedex:   pokedex,
	}
}

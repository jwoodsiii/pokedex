package main

import (
	"time"

	"github.com/jwoodsiii/pokedex/internal/pokeapi"
	"github.com/jwoodsiii/pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(10 * time.Second)
	pokeClient := pokeapi.NewClient(5 * time.Second, cache)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}

package main

import (
	"time"

	"github.com/jwoodsiii/pokedex/internal/pokeapi"
	"github.com/jwoodsiii/pokedex/internal/pokecache"
)

func main() {
	cache := pokecache.NewCache(10 * time.Second)
	pokedex := make(pokeapi.Pokedex, 0)
	pokeClient := pokeapi.NewClient(5*time.Second, cache, pokedex)
	cfg := &config{
		pokeapiClient: pokeClient,
	}
	startRepl(cfg)
}

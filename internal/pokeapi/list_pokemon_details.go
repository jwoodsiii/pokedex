package pokeapi

import (
	"encoding/json"
	"log"
)

func (c *Client) ListPokemonDetails(pokemon string) (Pokemon, error) {
	url := baseUrl + "pokemon/" + pokemon
	var poke Pokemon
	log.Printf("checking cache for data at url: %s", url)

	if val, ok := c.pokeCache.Get(url); ok {
		log.Println("cache hit, no need for GET request")
		if err := json.Unmarshal(val, &poke); err != nil {
			return Pokemon{}, err
		}
		return poke, nil
	}

	data, err := fetch(&c.httpClient, url)
	if err != nil {
		return Pokemon{}, err
	}
	log.Printf("adding url: %s to cache", url)
	c.pokeCache.Add(url, data)
	tmp, err := unmarshal[Pokemon](data)
	if err != nil {
		return Pokemon{}, nil
	}
	poke = *tmp
	return poke, nil
}

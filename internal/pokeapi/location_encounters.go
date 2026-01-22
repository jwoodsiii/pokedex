package pokeapi

import (
	"encoding/json"
	"log"
)

func (c *Client) ExploreLocationArea(city string) (EncounterArea, error) {
	url := baseUrl + "location-area/" + city
	var encounter EncounterArea
	log.Printf("checking cache for data at url: %s", url)
	if val, ok := c.pokeCache.Get(url); ok {
		log.Println("cache hit, no need for GET request")
		if err := json.Unmarshal(val, &encounter); err != nil {
			return EncounterArea{}, err
		}
		return encounter, nil
	}

	data, err := fetch(&c.httpClient, url)
	if err != nil {
		return EncounterArea{}, err
	}
	log.Printf("adding url: %s to cache", url)
	c.pokeCache.Add(url, data)
	tmp, err := unmarshal[EncounterArea](data)
	if err != nil {
		return EncounterArea{}, nil
	}
	encounter = *tmp
	return encounter, nil
}

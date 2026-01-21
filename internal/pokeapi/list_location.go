package pokeapi

import (
	"encoding/json"
	"log"
)

// ListLocationAreas -
func (c *Client) ListLocationAreas(pageUrl *string) (LocationAreas, error) {
	url := baseUrl + "location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	var locationArea LocationAreas
	// adding cache check before making network request
	log.Printf("checking cache for data at url: %s\n", url)
	if val, ok := c.pokeCache.Get(url); ok {
		log.Println("cache hit, no need for GET request")
		if err := json.Unmarshal(val, &locationArea); err != nil {
			return LocationAreas{}, err
		}
		return locationArea, nil
	}

	data, err := fetch(&c.httpClient, url)
	if err != nil {
		return LocationAreas{}, err
	}

	log.Printf("adding url: %s to cache", url)
	c.pokeCache.Add(url, data)
	tmp, err := unmarshal[LocationAreas](data)
	if err != nil {
		return LocationAreas{}, nil
	}
	locationArea = *tmp
	// if err := json.Unmarshal(data, &locationArea); err != nil {
	// 	return LocationAreas{}, err
	// }
	return locationArea, nil
}

package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
	"log"
)

// ListLocationArea -
func (c *Client) ListLocationAreas(pageUrl *string) (LocationArea, error) {
	url := baseUrl + "location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	var locationArea LocationArea
	// adding cache check before making network request
	log.Printf("checking cache for data at url: %s\n", url)
	if val, ok := c.pokeCache.Get(url); ok {
		log.Println("cache hit, no need for GET request")
		if err := json.Unmarshal(val, &locationArea); err != nil {
			return LocationArea{}, err
		}
		return locationArea, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	// if we get here its due to a cache miss, add our resp.Body to the cache for future queries
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	log.Printf("adding url: %s to cache", url)
	c.pokeCache.Add(url, data)
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}

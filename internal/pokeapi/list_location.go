package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

// ListLocationArea -
func (c *Client) ListLocationAreas(pageUrl *string) (LocationArea, error) {
	url := baseUrl + "/location-area"
	if pageUrl != nil {
		url = *pageUrl
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return LocationArea{}, err
	}
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return LocationArea{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return LocationArea{}, err
	}
	var locationArea LocationArea
	if err := json.Unmarshal(data, &locationArea); err != nil {
		return LocationArea{}, err
	}
	return locationArea, nil
}

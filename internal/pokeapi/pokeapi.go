package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const (
	baseUrl = "https://pokeapi.co/api/v2/"
)

// Abstract fetching data from PokeAPI endpoints
func fetch(client *http.Client, url string) ([]byte, error) {
	if url == "" {
		return nil, errors.New("url cannot be empty")
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// Abstract unmarshal of data received from fetch()
func unmarshal[T any](data []byte) (*T, error) {
	if data == nil {
		return nil, errors.New("byte array cannot be empty")
	}
	var output T
	if err := json.Unmarshal(data, &output); err != nil {
		return nil, err
	}
	return &output, nil
}

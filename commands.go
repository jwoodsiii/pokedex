package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"github.com/jwoodsiii/pokedex/internal/pokeapi"
)

// type locationArea struct {
// 	id int `json:"id"`
// 	name string `json:"name"`
// 	game_index int `json:"game_index"`
// 	encounter_method_rates []struct {
// 		method struct {
// 			id int `json:"id"`
// 			name string `json:"name"`
// 		} `json:"method"`
// 		rate int `json:"rate"`
// 	} `json:"encounter_method_rates"`

// }

// Exit the Pokedex
func commandExit(c *pokeapi.Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Display help information
func commandHelp(c *pokeapi.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

// pull location areas from Pokemon world, pulls in batches of 20
func commandMap(c *pokeapi.Config) error {
	var url string
	if c.Next != "" {
		url = c.Next
	} else {
		url = "https://pokeapi.co/api/v2/location-area/"
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locationAreas pokeapi.LocationArea // placeholder, will need to create type for location area
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return err
	}
	if locationAreas.Next != "" {
		c.Next = locationAreas.Next
	}
	if locationAreas.Previous != nil {
		c.Previous = locationAreas.Previous
	}
	// c.Next = locationAreas.Next
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	return nil
}

func commandMapB(c *pokeapi.Config) error {
	var url string
	if c.Previous != nil {
		url = *c.Previous
	} else {
		fmt.Println("You're on the first page")
		return nil
	}
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locationAreas pokeapi.LocationArea // placeholder, will need to create type for location area
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return err
	}
	if locationAreas.Next != "" {
		c.Next = locationAreas.Next
	}
	if locationAreas.Previous != nil {
		c.Previous = locationAreas.Previous
	}
	// c.Next = locationAreas.Next
	for _, result := range locationAreas.Results {
		fmt.Println(result.Name)
	}
	return nil
}

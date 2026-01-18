package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// type locationArea struct {
// }

// Exit the Pokedex
func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

// Display help information
func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

// pull location areas from Pokemon world, pulls in batches of 20
func commandMap() error {
	res, err := http.Get("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		return err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return err
	}

	var locationAreas []string // placeholder, will need to create type for location area
	if err := json.Unmarshal(data, &locationAreas); err != nil {
		return err
	}
	for _, area := range locationAreas {
		fmt.Println(area)
	}
	return nil
}

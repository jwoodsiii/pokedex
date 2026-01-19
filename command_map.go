package main

import (
	"errors"
	"fmt"
)

// pull location areas from Pokemon world, pulls in batches of 20

func commandMap(cfg *config) error {
	locationAreas, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextUrl)
	if err != nil {
		return err
	}
	cfg.nextUrl = &locationAreas.Next
	cfg.prevUrl = locationAreas.Previous

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapB(cfg *config) error {
	if cfg.prevUrl == nil {
		return errors.New("you're on the first page\n")
	}

	locationAreas, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevUrl)
	if err != nil {
		return err
	}
	cfg.nextUrl = &locationAreas.Next
	cfg.prevUrl = locationAreas.Previous

	for _, loc := range locationAreas.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

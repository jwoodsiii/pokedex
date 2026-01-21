package main

import (
	"errors"
	"fmt"
)

// pull location areas from Pokemon world, pulls in batches of 20

func commandMap(cfg *config, args ...string) error {
	if len(args) != 0 {
        fmt.Println("usage: map")
        return nil
    }
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

func commandMapB(cfg *config, args ...string) error {
	if len(args) != 0 {
        fmt.Println("usage: mapb")
        return nil
    }

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

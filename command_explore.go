package main

import (
	"fmt"
)

// list all encounterable Pokemon from given location area

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		fmt.Println("usage: explore <location-area>")
		return nil
	}
	city := args[0]
	encounters, err := cfg.pokeapiClient.ExploreLocationArea(city)
	if err != nil {
		return err
	}
	for _, pokemon := range encounters.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}
	return nil
}

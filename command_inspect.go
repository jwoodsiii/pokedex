package main

import (
	"errors"
)

// inspect details of pokemon in pokedex

func commandInspect(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: inspect <pokemon>")
	}
	pokemon := args[0]
	if err := cfg.pokeapiClient.InspectPokemon(cfg.pokeapiClient.Pokedex, pokemon); err != nil {
		return err
	}
	return nil
}

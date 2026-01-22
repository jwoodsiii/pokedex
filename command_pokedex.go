package main

import (
	"errors"
	"fmt"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(args) != 0 {
		return errors.New("usage: pokedex\n")
	}
	if len(cfg.pokeapiClient.Pokedex) == 0 {
		return errors.New("empty pokedex\n")
	}
	fmt.Println("Your Pokedex:")
	for k, _ := range cfg.pokeapiClient.Pokedex {
		fmt.Printf(" - %s\n", k)
	}
	return nil
}

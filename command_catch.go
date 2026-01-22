package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("usage: catch <pokemon>\n")
	}
	pokemon := args[0]
	pokemonDetails, err := cfg.pokeapiClient.ListPokemonDetails(pokemon)
	if err != nil {
		return err
	}
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	rnd := rand.Intn(pokemonDetails.BaseExperience)
	if rnd > int(.5*float64(pokemonDetails.BaseExperience)) {
		fmt.Printf("%s was caught!\n", pokemon)
		cfg.pokeapiClient.Pokedex[pokemon] = pokemonDetails
	} else {
		fmt.Printf("%s escaped!\n", pokemon)
	}

	return nil
}

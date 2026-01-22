package pokeapi

import (
	"errors"
	"fmt"
)

// abstract parse and display of following details from pokedex
/*
 Name
 Height
 Weight
 Stats:
   -hp:
   -attack
   -defense
   -special-attack
   -special-defense
   -speed
 Types:
*/

func (c *Client) InspectPokemon(pokedex Pokedex, pokemon string) error {
	if len(pokedex) == 0 {
		return errors.New("pokedex is empty")
	}
	details, ok := pokedex[pokemon]
	if !ok {
		return errors.New("you have not caught that pokemon")
	}
	fmt.Printf("Name: %s\n", details.Name)
	fmt.Printf("Height: %d\n", details.Height)
	fmt.Printf("Weight: %d\n", details.Weight)
	fmt.Println("Stats:")
	for _, stat := range details.Stats {
		switch stat.Stat.Name {
		case "hp":
			fmt.Printf("\t-hp: %d\n", stat.BaseStat)
		case "attack":
			fmt.Printf("\t-attack: %d\n", stat.BaseStat)
		case "defense":
			fmt.Printf("\t-defense: %d\n", stat.BaseStat)
		case "special-attack":
			fmt.Printf("\t-special-attack: %d\n", stat.BaseStat)
		case "special-defense":
			fmt.Printf("\t-special-defense: %d\n", stat.BaseStat)
		case "speed":
			fmt.Printf("\t-speed: %d\n", stat.BaseStat)
		default:
			continue
		}
	}
	fmt.Println("Types:")
	for _, t := range details.Types {
		fmt.Printf("\t-%s\n", t.Type.Name)
	}
	return nil
}

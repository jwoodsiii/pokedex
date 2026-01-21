package main

import (
	"fmt"
)

// Display help information
func commandHelp(cfg *config, args ...string) error {
	if len(args) != 0 {
        fmt.Println("usage: help")
        return nil
    }
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage: ")
	for _, v := range getCommands() {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}
	return nil
}

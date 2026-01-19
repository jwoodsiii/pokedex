package main

import (
	"fmt"
	"os"
)

// Exit the Pokedex
func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

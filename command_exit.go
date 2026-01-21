package main

import (
	"fmt"
	"os"
)

// Exit the Pokedex
func commandExit(cfg *config, args ...string) error {
	if len(args) != 0 {
        fmt.Println("usage: exit")
        return nil
    }

	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

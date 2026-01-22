package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jwoodsiii/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient pokeapi.Client
	nextUrl       *string
	prevUrl       *string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func startRepl(cfg *config) {
	// create scanner for user input
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex> ")
		if !scanner.Scan() { // wait for input
			break // exit on EOF or scan error
		}
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		cleaned := cleanInput(input)
		command := cleaned[0]
		exec, ok := getCommands()[command]
		if !ok {
			fmt.Println("Unknown command")
			continue
		}
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}
		if err := exec.callback(cfg, args...); err != nil {
			fmt.Printf("Error executing callback: %s, %s", exec.name, err)
		}
	}
}

func cleanInput(text string) []string {
	// split the user's input into "words" based on whitespace.also
	// lowercase the input and trim any leading or trailing whitespace
	return strings.Fields(strings.ToLower(text))
	//return []string{}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays a map of Pokemon locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previously shown map of Pokemon locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore",
			description: "Explore a location area, listing all pokemon that inhabit the provided area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Gotta catch 'em all",
			callback:    commandCatch,
		},
	}
}

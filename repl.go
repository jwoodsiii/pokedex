package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/jwoodsiii/pokedex/internal/pokeapi"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*pokeapi.Config) error
}

// var cmdRegistry = map[string]cliCommand {
// 	"exit": {
// 		name: "exit",
// 		description: "Exit the Pokedex",
// 		callback: commandExit,
// 	},
// 	"help": {
// 		name: "help",
// 		description: "Displays a help message",
// 		callback: commandHelp,
// 	},
// }

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
			name:			"mapb",
			description:	"Displays previously shown map of Pokemon locations",
			callback:		commandMapB,
		},
	}
}

func startRepl() {
	// create scanner for user input
	var config = &pokeapi.Config {
		Next: "",
		Previous: nil,
	}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		cleaned := cleanInput(input)
		command := cleaned[0]
		exec, ok := getCommands()[command]
		if !ok {
			fmt.Println("Unknown command")
		} else {
			if err := exec.callback(config); err != nil {
				fmt.Printf("Error executing callback: %s, %s", exec.name, err)
			}
		}
		fmt.Print("Pokedex > ")
	}
}

func cleanInput(text string) []string {
	// split the user's input into "words" based on whitespace.also
	// lowercase the input and trim any leading or trailing whitespace
	return strings.Fields(strings.ToLower(text))
	//return []string{}
}

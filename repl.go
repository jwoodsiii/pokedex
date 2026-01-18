package main

import (
	"os"
	"fmt"
	"strings"
	"bufio"
)

func startRepl() {
	// create scanner for user input
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Pokedex > ")
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) == 0 {
			continue
		}
		cleaned := cleanInput(input)
		first := cleaned[0]
		fmt.Printf("Your command was: %s\n", first)
		fmt.Print("Pokedex > ")
	}
}

func cleanInput(text string) []string {
	// split the user's input into "words" based on whitespace.also
	// lowercase the input and trim any leading or trailing whitespace
	return strings.Fields(strings.ToLower(text))
	//return []string{}
}

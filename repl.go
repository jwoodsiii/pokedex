package main

import "strings"

func cleanInput(text string) []string {
	// split the user's input into "words" based on whitespace.also
	// lowercase the input and trim any leading or trailing whitespace
	return strings.Fields(strings.ToLower(text))
	//return []string{}
}

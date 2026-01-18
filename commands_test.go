package main

import (
	"os/exec"
	"testing"
	"strings"
	"bytes"
)

/*
 * Do we have to have separate tests for each command in the registry?
 * Since we already generate the registry with a function we could map that function
 * and iterate over that in a single test
 * Future commands will likely have inputs/outputs and splitting them from the functions that don't
 * will likely just be a hassle
 */


// func TestCommandExit(t *testing.T) {
// 	var buf bytes.Buffer
// 	old := os.Stdout
// 	os.Stdout = &buf
// 	defer func() {os.Stdout = old}()

// 	if err := commandExit(); err != nil {
// 		t.Fatalf("commandExit returned error %v", err)
// 	}

// 	out := buf.String()
// 	if !strings.Contains(out, "Closing the Pokedex") {
// 		t.Fatalf("expected exit to close the Pokedex, but got %q", out)
// 	}
// }

func TestHelpAndExit(t *testing.T) {
    cmd := exec.Command("go", "run", ".")
    cmd.Stdin = strings.NewReader("help\nexit\n")

    var out bytes.Buffer
    cmd.Stdout = &out
    cmd.Stderr = &out

    err := cmd.Run()
    if err != nil {
        t.Fatalf("program failed: %v", err)
    }

    s := out.String()
    if !strings.Contains(s, "Welcome to the Pokedex!") {
        t.Errorf("expected help output, got %q", s)
    }
    if !strings.Contains(s, "Closing the Pokedex... Goodbye!") {
        t.Errorf("expected exit output, got %q", s)
    }
}

// func TestCommandHelp(t *testing.T) {
//     // capture stdout
//     var buf bytes.Buffer
//     old := os.Stdout
//     os.Stdout = &buf
//     defer func() { os.Stdout = old }()

//     err := commandHelp()
//     if err != nil {
//         t.Fatalf("commandHelp returned error: %v", err)
//     }

//     out := buf.String()
//     if !strings.Contains(out, "Welcome to the Pokedex!") {
//         t.Fatalf("expected help output to contain welcome message, got %q", out)
//     }
// }

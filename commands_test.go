package main

import (
	"bytes"
	"os/exec"
	"strings"
	"testing"
)

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

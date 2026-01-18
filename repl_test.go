package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"seperators": {input: "  hello  world  ", expected: []string{"hello", "world"}},
		"simple":     {input: "black effect", expected: []string{"black", "effect"}},
		"caps":       {input: "KATT WILLIAMS IS THE FUNNIEST", expected: []string{"katt", "williams", "is", "the", "funniest"}},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(tc.input)
			if len(actual) != len(tc.expected) {
				t.Fatalf("input slice len: %d does not match expected slice len: %d", len(actual), len(tc.expected))
			}
			for i := range actual {
				word := actual[i]
				expectedWord := tc.expected[i]
				if word != expectedWord {
					t.Fatalf("word: %v does not match expected word: %v", word, expectedWord)
				}
			}
		})
	}
}

package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input		string
		expected	[]string
	}{
		"seperators":	{input: "  hello  world  ", expected: []string{"hello", "world"}},
		"simple":		{input: "black effect", expected: []string{"black", "effect"}},
		"caps":			{input: "KATT WILLIAMS IS THE FUNNIEST", expected: []string{"katt", "williams", "is", "the", "funniest"}},
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

// func TestCleanInput(t *testing.T) {
// 	cases := []struct {
// 		input    string
// 		expected []string
// 	}{
// 		{
// 			input:    "  hello  world  ",
// 			expected: []string{"hello", "world"},
// 		},
// 		{
// 			input: "black effect",
// 			expected: []string{"black", "effect"},
// 		},
// 		{
// 			input: "KATT WILLIAMS IS THE FUNNIEST",
// 			expected: []string{"katt", "williams", "is", "the", "funniest"},
// 		},
// 		// add more cases here
// 	}

// 	for _, c := range cases {
// 		actual := cleanInput(c.input)
// 		// Check the length of the actual slice against the expected slice
// 		// if they don't match, use t.Errorf to print an error message
// 		// and fail the test
// 		if len(actual) != len(c.expected) {
// 			t.Fatalf("input slice len: %d does not match expected slice len: %d", len(actual), len(c.expected))
// 		}
// 		for i := range actual {
// 			word := actual[i]
// 			expectedWord := c.expected[i]
// 			// Check each word in the slice
// 			// if they don't match, use t.Errorf to print an error message
// 			// and fail the test
// 			if word != expectedWord {
// 				t.Fatalf("word: %v does not match expected word: %v", word, expectedWord)
// 			}
// 		}
// 	}
// }

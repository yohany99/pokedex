package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello WORLD",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Hello, WoRlD!  ",
			expected: []string{"hello,", "world!"},
		},
		{
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "     ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		t.Run(c.input, func(t *testing.T) {
			actual := cleanInput(c.input)
			if len(actual) != len(c.expected) {
				t.Fatal("lengths do not match")
			}
			for i := range actual {
				if actual[i] != c.expected[i] {
					t.Error("word mismatch")
				}
			}
		})
	}
}

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
			input:    "hello",
			expected: []string{"hello"},
		},
		{
			input:    "   leading spaces",
			expected: []string{"leading", "spaces"},
		},
		{
			input:    "trailing spaces   ",
			expected: []string{"trailing", "spaces"},
		},
		{
			input:    "multiple     spaces     between",
			expected: []string{"multiple", "spaces", "between"},
		},
		{
			input:    "mixed spaces\ttabs  and\nnewlines",
			expected: []string{"mixed", "spaces", "tabs", "and", "newlines"},
		},
		{
			input:    "   ", // Only spaces
			expected: []string{},
		},
		{
			input:    "\t\n  \t", // Only whitespace characters
			expected: []string{},
		},
		{
			input:    "word1,word2,word3",
			expected: []string{"word1,word2,word3"}, // No spaces to split on
		},
		{
			input:    "hello\nworld",
			expected: []string{"hello", "world"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		// Check the length of the actual slice against the expected slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
		if len(actual) != len(c.expected){
			t.Errorf("lengths don't match: '%v' vs '%v'", actual, c.expected)
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			// Check each word in the slice
			// if they don't match, use t.Errorf to print an error message
			// and fail the test
			if word != expectedWord {
				t.Errorf("cleanInput(%v) == %v, expected %v", c.input, actual, c.expected)
			}
		}
	}
	
}
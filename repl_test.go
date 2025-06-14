package main

import (
	"testing"
	"strings"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input		string
		expected	[]string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input: " I choose you ",
			expected: []string{"I", "choose", "you"},
		},
		// add more cases here
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("slices don't match")
		}

		for i := range actual {
			word := actual[i]
			expectedWord := strings.ToLower(c.expected[i])
			if word != expectedWord {
				t.Errorf("words don't match %s, %s", word, expectedWord)
			}
		}
	}
}
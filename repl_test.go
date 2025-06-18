package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	//Testing clean input function.
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "    hello world    ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Charmander BULBASAUR    pikachu",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		{
			input:    "I am sp3cia1 []",
			expected: []string{"i", "am", "sp3cia1", "[]"},
		},
	}

	for i, c := range cases {
		actual := cleanInput(c.input)
		t.Logf("Expected: %v\n", c.expected)
		t.Logf("Actual: %v\n", actual)

		if len(actual) != len(c.expected) {
			t.Errorf("length of actual does not match length expected: %d vs %d", len(actual), len(c.expected))
			t.FailNow()
		}

		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]

			if word != expectedWord {
				t.Errorf("Expected: %s, Actual: %s", expectedWord, word)
				t.Fail()
			}
		}

		t.Logf("Test %d - Pass\n", i+1)
	}

}

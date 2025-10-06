package main

import "testing"

func TestCleanInput(t *testing.T) {
	type testCase struct {
		input string
		expected []string
	}

	cases := []testCase {
		{input: "hello world", expected: []string{"hello", "world"}},
		{input: "this is a test case", expected: []string{"this", "is", "a", "test", "case"}},
		{input: "Charmander Bulbasaur PIKACHU", expected: []string{"charmander", "bulbasaur", "pikachu"}},
		{input: "This  is a  test of \t multiple  spaces \n and  other  different chars", expected: []string{"this", "is", "a", "test", "of", "multiple", "spaces", "and", "other", "different", "chars"}},
	}

for _, c := range cases {
	actual := cleanInput(c.input)
	// expected := c.expected
	len_actual := len(actual)
	len_expected := len(c.expected)
	if len_actual != len_expected {
		t.Errorf("length mismatch' Expected %s to have length %d; got %d", c.expected, len_expected, len_actual);
		continue
	}
	// Check the length of the actual slice against the expected slice
	// if they don't match, use t.Errorf to print an error message
	// and fail the test
	for i := range actual {
		word := actual[i]
		expectedWord := c.expected[i]
		if word != expectedWord {
			t.Errorf("word mismatch; Expected %s; got %s", expectedWord, word)
		}
		// Check each word in the slice
		// if they don't match, use t.Errorf to print an error message
		// and fail the test
	}
}
}
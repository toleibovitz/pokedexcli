package main

import "strings"

// split words by white space
func cleanInput(text string) []string {
	processed := []string{}
	split_words := strings.Fields(text)
	for _, w := range split_words {
		lowered := strings.ToLower(w);
		processed = append(processed, lowered)
	}
	
	return processed
}
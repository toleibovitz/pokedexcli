package main

import (
	"strings"
	"bufio"
	"os"
	"fmt"
)





func startRepl() {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}

		commandWord := words[0]
		cmdExit := commands["exit"]
		if commandWord == cmdExit.name {
			commandExit()
		}
	}
}

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
package repl

import (
	"strings"
	"bufio"
	"os"
	"fmt"
	"github.com/toleibovitz/pokedexcli/internal/commands"
)

func StartRepl() {
	reader := bufio.NewScanner(os.Stdin)
	// cfg := &commands.Config{}
	var cfg *commands.Config
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		words := cleanInput(reader.Text())
		if len(words) == 0 {
			continue
		}
		cmd := words[0]
		commands.GetCommands()[cmd].Callback(cfg)
		
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




package commands

import (
	"fmt"
	
)
func Help(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!\nUsage:")
	fmt.Println()
	cmds := GetCommands()
	for _, c := range cmds {
		fmt.Printf("%s: %s\n", c.name, c.description)
	}

	return nil
}
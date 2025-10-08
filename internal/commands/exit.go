package commands

import (
	"os"
	"fmt"
)

func Exit(cfg *Config) error{
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
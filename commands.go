package main

import "os"

type cliCommand struct {
	name string
	description string
	callback func() error
}

var commands = map[string]cliCommand {
	"exit": {
		name: "exit",
		description: "Exit the Pokedex",
		callback: commandExit,
	},
}

func commandExit() error{
	os.Exit(0)
	err := error.Error("Some error")
	return err
}
package commands

type cliCommand struct {
	name string
	description string
	cfg *Config
	Callback func(cfg *Config) error
	
}

func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			Callback:    Exit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			Callback:    Help,
		},
		"map": {
			name: "map",
			description: "Fetches location-areas 20 at a time, advances on addtional calls",
			Callback: Map,
		},
		"mapb": {
			name: "mapb",
			description: "shows previous results",
			Callback: Mapb,
		},
	}
}
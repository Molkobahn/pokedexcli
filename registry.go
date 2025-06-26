package main

type cliCommand struct {
	name		string
	description	string
	callback	func(*config) error
}


func GetCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:			"exit",
			description:	"Exit the Pokedex",
			callback:		commandExit,
		},
		"help": {
			name:			"help",
			description:	"Displays a help message",
			callback:		commandHelp,
		},
		"map": {
			name:			"map",
			description:	"Displays next list of locations",
			callback:		commandMap,
		},
		"mapb": {
			name:			"mapb",
			description:	"Diplays previous list of locations",
			callback:		commandMapb,
		},
		"explore": {
			name:			"explore",
			description:	"Diplays list of Pokemon in location",
			callback:		commandExplore,
		},
		"catch": {
			name:			"catch",
			description:	"Attempts to catch a pokemon",
			callback:		commandCatch,
		},
		"inspect": {
			name:			"inspect",
			description:	"Displays information about pokemon",
			callback:		commandInspect,
		},
		"pokedex": {
			name:			"pokedex",
			description:	"Displays a list of caught pokemon",
			callback:		commandPokedex,
		},
	}
}
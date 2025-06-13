package main

type cliCommand struct {
	name		string
	description	string
	callback	func() error
}


func getCommands() map[string]cliCommand {
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
	}
}
package main

import (
	"github.com/molkobahn/pokedexcli/internal/pokecache"
	"time"
)

type config struct {
	pokecache	*pokecache.Cache
	Next		string
	Previous	string
}

func main() {
	Cache := pokecache.NewCache(5*time.Second)
	
	cfg := &config{
		pokecache:	Cache,
		Next: 		"https://pokeapi.co/api/v2/location-area/",
		Previous:	"",
	}
	commands := map[string]cliCommand{
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
	startRepl(commands, cfg)
}

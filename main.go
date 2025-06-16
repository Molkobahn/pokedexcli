package main

import (
	"github.com/molkobahn/pokedexcli/internal/pokecache"
	"time"
)

type config struct {
	pokecache	*pokecache.Cache
	Next		string
	Previous	string
	Parameter	string
}

func main() {
	Cache := pokecache.NewCache(5*time.Second)
	
	cfg := &config{
		pokecache:	Cache,
		Next: 		"https://pokeapi.co/api/v2/location-area/",
		Previous:	"",
		Parameter:	"",
	}
	commands := GetCommands()

	startRepl(commands, cfg)
}

package main

import (
	"github.com/molkobahn/pokedexcli/internal/pokecache"
	"time"
	"github.com/molkobahn/pokedexcli/internal/pokeapi"
)

type config struct {
	pokecache	*pokecache.Cache
	Next		string
	Previous	string
	Parameter	string
	Pokedex		map[string]pokeapi.RespPokemon
}

func main() {
	Cache := pokecache.NewCache(5*time.Second)
	
	cfg := &config{
		pokecache:	Cache,
		Next: 		"https://pokeapi.co/api/v2/location-area/",
		Previous:	"",
		Parameter:	"",
		Pokedex:	make(map[string]pokeapi.RespPokemon),
	}
	commands := GetCommands()

	startRepl(commands, cfg)
}

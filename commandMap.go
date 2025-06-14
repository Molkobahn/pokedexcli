package main

import (
	"fmt"
	"github.com/molkobahn/pokedexcli/internal/pokeapi"
)


func commandMap(cfg *config) error{
	maps, err := pokeapi.GetRequest(cfg.Next, cfg.pokecache)
	if err != nil {
		return err
	}

	for _, pokeMaps := range maps.Results {
		fmt.Printf("%s\n", pokeMaps.Name)
	}
	
	cfg.Previous = maps.Previous
	cfg.Next = maps.Next

	return nil
}

func commandMapb(cfg *config) error {
	
	if cfg.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	maps, err := pokeapi.GetRequest(cfg.Previous, cfg.pokecache)
	if err != nil {
		return err
	}
	
	for _, pokeMaps := range maps.Results {
		fmt.Printf("%s\n", pokeMaps.Name)
	}

	cfg.Previous = maps.Previous
	cfg.Next = maps.Next
	
	return nil

}
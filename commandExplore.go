package main

import (
	"fmt"
	"github.com/molkobahn/pokedexcli/internal/pokeapi"
	"errors"
)

func commandExplore(cfg *config) error {
	if cfg.Parameter == "" {
		return errors.New("You must provide a location name")
	}
	
	url := "https://pokeapi.co/api/v2/location-area/" + cfg.Parameter

	pokeMap, err := pokeapi.LocationDetails(url, cfg.pokecache)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s\n", pokeMap.Location.Name)
	fmt.Println("Found Pokemon:")

	for _, encounters := range pokeMap.PokemonEncounters {
		fmt.Printf("%s\n", encounters.Pokemon.Name)
		}
	
	return nil
}
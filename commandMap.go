package main

import (
	"fmt"
	"github.com/molkobahn/pokedexcli/internal/pokeapi"
)
var maps = pokeapi.GetRequest("https://pokeapi.co/api/v2/location-area/")

func commandMap() error{
	
	for _, pokeMaps := range maps.Results {
		fmt.Printf("%s\n", pokeMaps.Name)
	}
	
	pokeapi.MapConfig.Previous = maps.Previous
	pokeapi.MapConfig.Next = maps.Next
	
	maps = pokeapi.GetRequest(pokeapi.MapConfig.Next)

	return nil
}

func commandMapb() error {
	
	if pokeapi.MapConfig.Previous == "" {
		fmt.Println("You're on the first page")
		return nil
	}

	maps = pokeapi.GetRequest(pokeapi.MapConfig.Previous)

	for _, pokeMaps := range maps.Results {
		fmt.Printf("%s\n", pokeMaps.Name)
	}

	pokeapi.MapConfig.Previous = maps.Previous
	pokeapi.MapConfig.Next = maps.Next
	

	return nil

}
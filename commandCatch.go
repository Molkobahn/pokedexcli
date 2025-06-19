package main

import(
	"fmt"
	"math/rand"
	"errors"
	"github.com/molkobahn/pokedexcli/internal/pokeapi"
)


func commandCatch(cfg *config) error {
	if cfg.Parameter == "" {
		errors.New("You must provide a Pokemon name")
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", cfg.Parameter)

	url := "https://pokeapi.co/api/v2/pokemon/" + cfg.Parameter

	pokemon, err := pokeapi.GetPokemon(url, cfg.pokecache)
	if err != nil {
		return err
	}

	num := rand.Intn(pokemon.BaseExperience)
	if num >= pokemon.BaseExperience / 20 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.Pokedex[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}
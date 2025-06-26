package main

import(
	"fmt"
)

func commandPokedex(cfg *config) error {
	fmt.Println("Your Pokedex:")
	for _, pokemon := range cfg.Pokedex {
		fmt.Printf("    - %v\n", pokemon.Name)
	}
	return nil
}
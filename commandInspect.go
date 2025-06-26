package main

import(
	"fmt"
)

func commandInspect(cfg *config) error {
	pokemon, exists := cfg.Pokedex[cfg.Parameter]
	if !exists {
		fmt.Println("You have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)
	fmt.Printf("Stats:\n")
	for _, stat := range pokemon.Stats {
		fmt.Printf("    -%v: %v\n", stat.Stat.Name, stat.BaseStat)
	}
	fmt.Println("Types:")
	for _, types := range pokemon.Types {
		fmt.Printf("    -%v\n", types.Type.Name)
	}
	return nil
}
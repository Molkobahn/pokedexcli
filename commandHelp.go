package main

import(
	"fmt"
)

func commandHelp(cfg *config) error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("exit:	Exit the Pokedex")
	fmt.Println("help:	Displays a help message")
	fmt.Println("map:	Displays next list of locations")
	fmt.Println("mapb:	Displays previous list of locations")
	fmt.Println()
	return nil
}
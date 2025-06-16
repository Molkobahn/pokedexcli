package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

func startRepl(commands map[string]cliCommand, cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		input := scanner.Text()
		words := cleanInput(input)
		
		if len(words) > 1 {
			cfg.Parameter = words[1]
		}
		
		command, exists := commands[words[0]]
		if exists {
			err := command.callback(cfg)
			if err != nil {
					fmt.Print(err)
				}
				continue
		} else {
				fmt.Print("Unknown command")
		}

	}
}


func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}
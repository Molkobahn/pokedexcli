package main

import (
	"strings"
	"fmt"
	"bufio"
	"os"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		
		input := scanner.Text()
		words := cleanInput(input)
		
		command, exists := getCommands()[words[0]]
		if exists {
			err := command.callback()
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
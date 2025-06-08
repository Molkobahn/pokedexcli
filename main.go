package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	finalSlice := make([]string, 0)
	slice := strings.Fields(text)

	for i := range slice {
		finalSlice = append(finalSlice, slice[i])
	}

	return finalSlice
}
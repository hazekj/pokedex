package main

import "fmt"

func commandHelp() error {
	fmt.Println()
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for _, com := range getCommands() {
		fmt.Println(com.name, ": ", com.description)
	}
	fmt.Println()
	return nil
}

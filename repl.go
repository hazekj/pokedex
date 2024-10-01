package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/hazekj/pokedex/internal/pokeapi"
)

type config struct {
	pokeClient  pokeapi.Client
	NextURL     string
	PreviousURL string
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

func repl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println()
		fmt.Print("Pokedex>")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Command not found")
			fmt.Println()
		} else {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

}

func cleanInput(text string) []string {
	lowerText := strings.ToLower(text)
	wordsText := strings.Fields(lowerText)
	return wordsText
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "List map regions",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "List map regions (backwards)",
			callback:    commandMapBack,
		},
	}
}

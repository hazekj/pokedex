package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex>")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		commandName := input[0]
		command, ok := getCommands()[commandName]
		if !ok {
			fmt.Println("Command not found")
		} else {
			command.callback()
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
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

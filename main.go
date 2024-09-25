package main

import (
	"bufio"
	"fmt"
	"os"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	fmt.Println("Pokedex>")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())
}

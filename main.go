package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Pokedex>")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fmt.Println(scanner.Text())
}

package main

import (
	"time"

	"github.com/hazekj/pokedex/internal/pokeapi"
)

func main() {
	pokeClient := pokeapi.NewClient(5 * time.Second)
	cfg := &config{pokeClient: pokeClient}
	repl(cfg)
}

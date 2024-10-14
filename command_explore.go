package main

import "fmt"

func commandExplore(cfg *config, area string) error {
	pokeEncounters, err := cfg.pokeClient.ListPokemon(area)
	if err != nil {
		return err
	}

	fmt.Println("Exploring: ", area)
	fmt.Println("Pokemon found: ")
	for _, encounter := range pokeEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}
	return nil
}

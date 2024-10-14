package main

import (
	"fmt"
)

func commandMap(cfg *config, _ string) error {
	locations, err := cfg.pokeClient.ListLocations(cfg.NextURL)
	if err != nil {
		return err
	}
	cfg.PreviousURL = locations.Previous
	cfg.NextURL = locations.Next
	for i, loc := range locations.Results {
		fmt.Printf("(%d) %s \n", i, loc.Name)
	}

	return nil

}

func commandMapBack(cfg *config, _ string) error {
	locations, err := cfg.pokeClient.ListLocations(cfg.PreviousURL)
	if err != nil {
		return err
	}
	cfg.PreviousURL = locations.Previous
	cfg.NextURL = locations.Next
	for i, loc := range locations.Results {
		fmt.Printf("(%d) %s \n", i, loc.Name)
	}
	return nil
}

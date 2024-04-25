package main

import (
	"errors"
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	// if the user doesn't include the location area
	if len(args) != 1 {
		return errors.New("please include a location name")
	}

	locationName := args[0]

	location, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Printf("Exploring %v...\n", location.Name)
	fmt.Println("Found Pokemon: ")

	for _, pokemon := range location.PokemonEncounters {
		fmt.Printf("- %v\n", pokemon.Pokemon.Name)
	}
	fmt.Println()
	return nil
}

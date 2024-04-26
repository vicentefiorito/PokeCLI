package main

import (
	"errors"
	"fmt"
)

func commandCatch(cfg *config, args ...string) error {
	// if the user doesn't input the pokemon name
	if len(args) != 1 {
		return errors.New("please include a pokemon name")
	}

	// gets the pokemon name from args
	pokemonName := args[0]

	// gets the data from the API call
	pokemon, err := cfg.pokeapiClient.GetPokemon(pokemonName)
	if err != nil {
		return err
	}

	fmt.Println(pokemon)
	return nil
}

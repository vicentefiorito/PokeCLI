package main

import (
	"errors"
	"fmt"
	"math/rand"
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

	// generates the catch based on based experience
	catchRate := rand.Intn(pokemon.BaseExperience)

	fmt.Printf("Throwing a Pokeball at %s\n", pokemon.Name)
	fmt.Println(catchRate)

	// if catch rate is over than the treshold, then the catch
	// fails
	if catchRate > 40 {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		return nil
	}

	fmt.Printf("%s was caught!\n", pokemon.Name)

	// if catch successfull, add it into the user pokedex
	cfg.pokedex[pokemon.Name] = pokemon

	return nil
}

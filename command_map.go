package main

import (
	"errors"
	"fmt"
	"log"
)

// displays the 20 forward locations
func commandMapf(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)

	if err != nil {
		log.Fatal(err)
	}

	// getting the next page
	cfg.nextLocationAreaURL = resp.Next
	// getting the previous page
	cfg.prevLocationAreaURL = resp.Previous

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	fmt.Println()

	return nil
}

// displays the previous 20 locations
func commandMapb(cfg *config) error {
	// if we are on the first page
	// return an error
	if cfg.prevLocationAreaURL == nil {
		return errors.New("you're on the first page")
	}

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocationAreaURL)

	if err != nil {
		log.Fatal(err)
	}

	// getting the next page
	cfg.nextLocationAreaURL = resp.Next
	// getting the previous page
	cfg.prevLocationAreaURL = resp.Previous

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	fmt.Println()

	return nil

}

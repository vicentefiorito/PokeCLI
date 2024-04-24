package main

import (
	"fmt"
	"log"
)

// displays the 20 forward locations
func commandMapf(cfg *config) error {

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)

	if err != nil {
		log.Fatal(err)
	}

	for _, area := range resp.Results {
		fmt.Println(area.Name)
	}
	fmt.Println()

	// getting the next page
	cfg.nextLocationAreaURL = resp.Next
	// getting the previous page
	cfg.prevLocationAreaURL = resp.Previous
	return nil
}

// // displays the previous 20 locations
// func commandMapb() error {

// }

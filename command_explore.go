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

	resp, err := cfg.pokeapiClient.GetLocation(locationName)
	if err != nil {
		return err
	}

	fmt.Println(resp)
	return nil
}

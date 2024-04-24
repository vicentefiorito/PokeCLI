package main

import (
	"fmt"
	"log"

	"github.com/vicentefiorito/pokeCLI/internal/pokeapi"
)


func main() {
	// testing the api call
	pokeapiClient := pokeapi.NewClient()

	resp, err := pokeapiClient.ListLocationAreas()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(resp)

	// startRepl()
}
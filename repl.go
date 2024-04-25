package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/vicentefiorito/pokeCLI/internal/pokeapi"
)

// config type that keeps state
// to keep track of pagination
type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLocationAreaURL *string
}

// initializes the repl to always be active and listening
// to the command line
func startRepl(cfg *config) {
	for {
		// this reads from the standard input in the console
		reader := bufio.NewScanner(os.Stdin)
		// the repl is always active unless the exit command is given
		fmt.Print("Pokedex > ")
		reader.Scan()

		// gets the user input from the scanner and
		// stores it in a variable
		input := cleanInput(reader.Text())

		// if there is no input from the user
		// go to the next iteration
		if len(input) == 0 {
			continue
		}

		commandName := input[0]

		// using the command parsed from the user
		command, exists := getCommands()[commandName]
		if exists {
			err := command.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			fmt.Println()
			continue
		}
	}

}

// helper function that cleans up the user input
func cleanInput(text string) []string {
	// puts the command into lower case
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

// types to architect the cliCommands
type cliCommands struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

// this function returns all the commands that are currently
// available to use
func getCommands() map[string]cliCommands {

	return map[string]cliCommands{
		"help": {
			name:        "help",
			description: "displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "displays the names of 20 locations",
			callback:    commandMapf,
		},
		"mapb": {
			name:        "mapb",
			description: "displays the previous 20 locations",
			callback:    commandMapb,
		},
		"exit": {
			name:        "exit",
			description: "exit the pokedex",
			callback:    commandExit,
		},
	}
}

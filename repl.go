package main

import (
	"bufio"
	"fmt"
	"os"
)

// initializes the repl to always be active and listening
// to the command line
func startRepl() {
	// this reads from the standard input in the console
	reader := bufio.NewScanner(os.Stdin)
	// the repl is always active unless the exit command is given
	fmt.Println("Pokedex > ")
	reader.Scan()
	fmt.Println("You entered:", reader.Text())

}

// types to architect the cliCommands
type cliCommands struct {
	name        string
	description string
	callback    func() error
}

// this function returns all the commands that are currently
// available to use
// func getCommands() map[string]cliCommands {

// 	return map[string]cliCommands{
// 		"help": {
// 			name: "help",
// 			description: "displays a help message",
// 			callback: commandHelp,
// 		},
// 		"exit": {
// 			name: "exit",
// 			description: "exit the pokedex",
// 			callback: commandExit,
// 		},

// 	}
// }

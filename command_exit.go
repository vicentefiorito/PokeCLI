package main

import (
	"os"
)

// all this command does is return from the prompt

func commandExit() error {
	os.Exit(0)
	return nil
}

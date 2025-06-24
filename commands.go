package main

import (
	"fmt"
	"os"
)

var commands map[string]cliCommand

func initCommands() {
	commands = map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the names of next 20 location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the names of previous 20 location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <area_name>",
			description: "Display the names of pokemon encounters in an area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Attempts to catch a pokemon",
			callback:    commandCatch,
		},
	}
}

func commandExit(cfg *Config, target string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config, target string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return nil
}

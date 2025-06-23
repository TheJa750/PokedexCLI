package main

import (
	"fmt"
	"os"

	"github.com/TheJa750/pokedexcli/internal/pokeapi"
	"github.com/TheJa750/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config) error
}

type Config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	Cache         *pokecache.Cache
}

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
	}
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return nil
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/TheJa750/pokedexcli/internal/pokeapi"
)

func cleanInput(text string) []string {
	var first, strs []string
	trim := strings.Trim(text, " ")
	lower := strings.ToLower(trim)
	first = strings.Split(lower, " ")

	for i, s := range first {
		if s != "" {
			strs = append(strs, first[i])
		}
	}

	return strs
}

func main() {
	initCommands()
	scanner := bufio.NewScanner(os.Stdin)
	cfg := &Config{
		Next:     nil,
		Previous: nil,
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		clean := cleanInput(scanner.Text())
		//fmt.Printf("Debug: clean = %v, len = %d\n", clean, len(clean))
		if len(clean) > 0 {
			command, ok := commands[clean[0]]
			if !ok {
				fmt.Println("Unknown command")
			} else {
				command.callback(cfg)
			}
		}
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

func commandMap(cfg *Config) error {
	// call map function in api package "https://pokeapi.co/api/v2/location"
	url := ""

	if cfg.Next == nil {
		url = "https://pokeapi.co/api/v2/location-area"
	} else {
		url = *cfg.Next
	}

	loc, err := pokeapi.GetMapData(url)
	if err != nil {
		return err
	}
	cfg.Next = &loc.Next
	cfg.Previous = loc.Previous

	//fmt.Printf("Debug: Next URL: %v, Previous URL: %v\n", loc.Next, loc.Previous)

	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	url := *cfg.Previous

	loc, err := pokeapi.GetMapData(url)
	if err != nil {
		return err
	}
	cfg.Next = &loc.Next
	cfg.Previous = loc.Previous

	//fmt.Printf("Debug: Next URL: %v, Previous URL: %v\n", loc.Next, loc.Previous)

	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}

	return nil
}

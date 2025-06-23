package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *Config) error {
	// call map function in api package "https://pokeapi.co/api/v2/location-area"

	loc, err := cfg.pokeapiClient.GetMapData(cfg.Next)
	if err != nil {
		return err
	}
	cfg.Next = loc.Next
	cfg.Previous = loc.Previous

	//fmt.Printf("Debug: Next URL: %v, Previous URL: %v\n", loc.Next, loc.Previous)

	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}

	loc, err := cfg.pokeapiClient.GetMapData(cfg.Previous)
	if err != nil {
		return err
	}
	cfg.Next = loc.Next
	cfg.Previous = loc.Previous

	//fmt.Printf("Debug: Next URL: %v, Previous URL: %v\n", loc.Next, loc.Previous)

	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}

	return nil
}

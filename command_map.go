package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/TheJa750/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *Config, area string) error {
	// call map function in api package "https://pokeapi.co/api/v2/location-area"
	//fmt.Println("DEBUG: commandMap started")
	//fmt.Printf("DEBUG: cfg.Next = %v\n", cfg.Next)
	callAPI := true
	var loc pokeapi.RespShallowLocations
	var err error

	if cfg.Next != nil {
		//fmt.Println("DEBUG: Checking cache...")
		data, exists := cfg.Cache.Get(*cfg.Next)
		if exists {
			err := json.Unmarshal(data, &loc)
			if err == nil {
				callAPI = false
			}
		}
	}

	if callAPI {
		//fmt.Println("DEBUG: Calling API...")
		loc, err = cfg.pokeapiClient.GetMapData(cfg.Next)
		if err != nil {
			//fmt.Printf("DEBUG: API error: %v\n", err)
			return err
		}
		//fmt.Printf("DEBUG: Got %d results from API\n", len(loc.Results))
		data, err := json.Marshal(loc)
		if err != nil {

			return err
		}
		if cfg.Next != nil {
			cfg.Cache.Add(*cfg.Next, data)
		}
	}

	cfg.Next = loc.Next
	cfg.Previous = loc.Previous

	//fmt.Printf("Debug: Next URL: %v, Previous URL: %v\n", loc.Next, loc.Previous)

	//fmt.Printf("DEBUG: About to print %d locations\n", len(loc.Results))

	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandMapb(cfg *Config, area string) error {
	if cfg.Previous == nil {
		return errors.New("you're on the first page")
	}
	callAPI := true
	var loc pokeapi.RespShallowLocations
	var err error

	data, exists := cfg.Cache.Get(*cfg.Previous)
	if exists {
		err := json.Unmarshal(data, &loc)
		if err == nil {
			callAPI = false
		}
	}

	if callAPI {
		loc, err = cfg.pokeapiClient.GetMapData(cfg.Previous)
		if err != nil {
			return err
		}
		data, err := json.Marshal(loc)
		if err != nil {
			return err
		}
		cfg.Cache.Add(*cfg.Previous, data)
	}
	cfg.Next = loc.Next
	cfg.Previous = loc.Previous

	//fmt.Printf("Debug: Next URL: %v, Previous URL: %v\n", loc.Next, loc.Previous)

	for _, area := range loc.Results {
		fmt.Println(area.Name)
	}

	return nil
}

func commandExplore(cfg *Config, area string) error {
	if area == "" {
		return errors.New("invalid location")
	}
	callAPI := true
	var pokemon pokeapi.AreaPokemon
	var err error

	data, exists := cfg.Cache.Get(area)
	if exists {
		err = json.Unmarshal(data, &pokemon)
		if err == nil {
			callAPI = false
		}
	}

	if callAPI {
		pokemon, err = cfg.pokeapiClient.ExploreArea(area)
		if err != nil {
			return err
		}
		data, err := json.Marshal(pokemon)
		if err != nil {
			return err
		}
		cfg.Cache.Add(area, data)
	}

	fmt.Printf("Exploring %s...\n", area)
	fmt.Println("Found Pokemon:")

	for _, mon := range pokemon.PokemonEncounters {
		fmt.Printf("- %s\n", mon.Pokemon.Name)
	}

	return nil
}

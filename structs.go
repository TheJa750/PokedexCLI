package main

import (
	"github.com/TheJa750/pokedexcli/internal/pokeapi"
	"github.com/TheJa750/pokedexcli/internal/pokecache"
)

type cliCommand struct {
	name        string
	description string
	callback    func(cfg *Config, area string) error
}

type Config struct {
	pokeapiClient pokeapi.Client
	Next          *string
	Previous      *string
	Cache         *pokecache.Cache
	Pokedex       map[string]Pokemon
}

type PokeInfo struct {
	Height int
	Weight int
	Stats  struct {
		hp      int
		attack  int
		defense int
		specAtk int
		specDef int
		speed   int
	}
	Typing []string
}

type Pokemon struct {
	Name   string
	Seen   bool
	Caught bool
	Info   PokeInfo
}

package main

import (
	"time"

	"github.com/TheJa750/pokedexcli/internal/pokeapi"
	"github.com/TheJa750/pokedexcli/internal/pokecache"
)

func main() {
	initCommands()
	cfg := &Config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
		Next:          nil,
		Previous:      nil,
		Cache:         pokecache.NewCache(15 * time.Minute),
		Pokedex:       make(map[string]Pokemon),
	}

	startRepl(cfg)
}

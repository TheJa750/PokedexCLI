package main

import (
	"time"

	"github.com/TheJa750/pokedexcli/internal/pokeapi"
)

func main() {
	initCommands()
	cfg := &Config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
		Next:          nil,
		Previous:      nil,
	}

	startRepl(cfg)
}

package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"

	"github.com/TheJa750/pokedexcli/internal/pokeapi"
)

func commandCatch(cfg *Config, target string) error {
	if target == "" {
		return errors.New("invalid pokemon")
	}

	if cfg.SeeBeforeCatch {
		_, exists := cfg.Pokedex[target]
		if !exists {
			return fmt.Errorf("you have not found a %s anywhere", target)
		}
	}

	callAPI := true
	var pokemonInfo pokeapi.CatchPokemonInfo
	var err error
	var info PokeInfo

	data, exists := cfg.Cache.Get(target)
	if exists {
		err := json.Unmarshal(data, &pokemonInfo)
		if err == nil {
			callAPI = false
		}
	}

	if callAPI {
		pokemonInfo, err = cfg.pokeapiClient.GetPokemonInfo(target)
		if err != nil {
			return err
		}

		data, err := json.Marshal(pokemonInfo)
		if err != nil {
			return err
		}
		cfg.Cache.Add(target, data)
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", target)

	exp := pokemonInfo.BaseExperience
	chance := rand.Intn(exp)
	var req int

	if exp < 75 {
		req = 25
	} else if exp < 150 {
		req = 35
	} else if exp < 250 {
		req = 55
	} else {
		req = 75
	}

	if chance < req {
		writePokeInfo(&info, pokemonInfo)
		fmt.Printf("%s was caught!\n", target)
		addToDex(cfg, target, info)
	} else {
		fmt.Printf("%s escaped!\n", target)
	}

	return nil
}

func addToDex(cfg *Config, name string, info PokeInfo) {
	updated := Pokemon{
		Name:   name,
		Seen:   true,
		Caught: true,
		Info:   info,
	}

	cfg.Pokedex[name] = updated
	fmt.Printf("%s was added to the Pokedex\n", name)
}

func writePokeInfo(writeInfo *PokeInfo, readInfo pokeapi.CatchPokemonInfo) {
	writeInfo.Height = readInfo.Height
	writeInfo.Weight = readInfo.Weight

	for _, stat := range readInfo.Stats {
		switch stat.Stat.Name {
		case "attack":
			writeInfo.Stats.attack = stat.BaseStat
		case "hp":
			writeInfo.Stats.hp = stat.BaseStat
		case "defense":
			writeInfo.Stats.defense = stat.BaseStat
		case "special-attack":
			writeInfo.Stats.specAtk = stat.BaseStat
		case "special-defense":
			writeInfo.Stats.specDef = stat.BaseStat
		case "speed":
			writeInfo.Stats.speed = stat.BaseStat
		}
	}

	for _, elem := range readInfo.Types {
		writeInfo.Typing = append(writeInfo.Typing, elem.Type.Name)
	}
}

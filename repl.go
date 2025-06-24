package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *Config) {
	reader := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		reader.Scan()
		clean := cleanInput(reader.Text())
		//fmt.Printf("Debug: clean = %v, len = %d\n", clean, len(clean))
		if len(clean) > 0 {
			command, ok := commands[clean[0]]
			if !ok {
				fmt.Println("Unknown command")
			} else {
				multiArgs := false
				if clean[0] == "explore" && len(clean) >= 2 {
					multiArgs = true
				}
				if clean[0] == "catch" && len(clean) >= 2 {
					multiArgs = true
				}

				if multiArgs {
					//fmt.Printf("DEBUG: Calling commandExplore with %s area", clean[1])
					err := command.callback(cfg, clean[1])
					if err != nil {
						fmt.Println(err)
					}
				} else {
					err := command.callback(cfg, "")
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}
	}
}

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

	/*
		//Debugging input
		for _, str := range strs {
			println(str)
		}
	*/

	return strs
}

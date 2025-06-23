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
				command.callback(cfg)
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

	return strs
}

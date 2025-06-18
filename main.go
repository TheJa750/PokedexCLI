package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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
				command.callback()
			}
		}
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")

	for _, v := range commands {
		fmt.Printf("%s: %s\n", v.name, v.description)
	}

	return nil
}

func commandMap() error {
	// call map function in api package

	return nil
}

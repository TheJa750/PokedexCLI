package main

import "fmt"

func commandInspect(cfg *Config, target string) error {
	mon, exists := cfg.Pokedex[target]
	if !exists {
		return fmt.Errorf("you have not caught a %s", target)
	}

	fmt.Printf("Name: %s\n", mon.Name)
	fmt.Printf("Height: %d\n", mon.Info.Height)
	fmt.Printf("Weight: %d\n", mon.Info.Weight)

	fmt.Println("Stats:")
	fmt.Printf("  - hp: %d\n", mon.Info.Stats.hp)
	fmt.Printf("  - attack: %d\n", mon.Info.Stats.attack)
	fmt.Printf("  - defense: %d\n", mon.Info.Stats.defense)
	fmt.Printf("  - special-attack: %d\n", mon.Info.Stats.specAtk)
	fmt.Printf("  - special-defense: %d\n", mon.Info.Stats.specDef)
	fmt.Printf("  - speed: %d\n", mon.Info.Stats.speed)

	fmt.Println("Types:")
	for _, elem := range mon.Info.Typing {
		fmt.Printf("  - %s\n", elem)
	}

	return nil
}

func commandPokedex(cfg *Config, target string) error {
	if len(cfg.Pokedex) == 0 {
		return fmt.Errorf("your pokedex is empty, go catch some pokemon")
	}

	count := 0
	fmt.Println("Your Pokedex:")

	for name, mon := range cfg.Pokedex {
		if mon.Caught {
			count++
			fmt.Printf("Caught - %s\n", name)
		} else {
			fmt.Printf("Seen - %s\n", name)
		}

	}

	fmt.Printf("Total seen: %d\n", len(cfg.Pokedex))
	fmt.Printf("Total caught: %d\n", count)

	return nil
}

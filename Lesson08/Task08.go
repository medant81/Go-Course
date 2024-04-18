package main

import (
	"fmt"
)

func main() {

	// Task 8.1
	mapAnimals := make(map[string]struct{}, 4)
	mapAnimals["слон"] = struct{}{}
	mapAnimals["бегемот"] = struct{}{}
	mapAnimals["носорог"] = struct{}{}
	mapAnimals["лев"] = struct{}{}

	fmt.Println(mapAnimals)

	// Task 8.2
	var animalsCount map[string]int
	animalsCount = map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}

	fmt.Println(animalsCount)

	c, ok := animalsCount["слон"]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", "слон", c, ok)
	c, ok = animalsCount["бегемот"]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", "бегемот", c, ok)
	c, ok = animalsCount["осьминог"]
	fmt.Printf("Животное: %s, количество: %d (есть в карте: %v)\n", "осьминог", c, ok)

	// Task 8.3
	var animalsEmptyStruct map[string]struct{}
	animalsEmptyStruct = map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	delete(animalsEmptyStruct, "бегемот")
	fmt.Println(animalsEmptyStruct)

	// Task 8.4
	bigAnimals := make(map[string]struct{}, 5)
	bigAnimals = map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	bigAnimals["выдра"] = struct{}{}
	fmt.Println(bigAnimals)

	// Task 8.5
	bigAnimalsCount := make(map[string]int, 4)
	bigAnimalsCount = map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}
	bigAnimalsCount["бегемот"] = 2
	fmt.Println(bigAnimalsCount)
}

package main

import "fmt"

func fruitMarket(nameFruit string) int {

	mapFruit := make(map[string]int, 4)
	mapFruit = map[string]int{
		"апельсин": 5,
		"яблоки":   3,
		"сливы":    1,
		"груши":    0,
	}

	f, ok := mapFruit[nameFruit]

	if ok {
		return f
	}

	fmt.Printf("not found fruit %s\n", nameFruit)
	return 0
}

func checkFood(nameFood string) {

	fruits := make(map[string]struct{}, 3)
	fruits = map[string]struct{}{
		"груша":    {},
		"яблоко":   {},
		"апельсин": {},
	}

	vegetables := make(map[string]struct{}, 3)
	vegetables = map[string]struct{}{
		"тыква":   {},
		"огурец":  {},
		"помидор": {},
	}

	_, isFruit := fruits[nameFood]
	_, isVegetable := vegetables[nameFood]

	switch {
	case isFruit:
		fmt.Println("Это фрукт")
	case isVegetable:
		fmt.Println("Это овощ")
	default:
		fmt.Println("Что-то странное...")
	}
}

func main() {

	// Task 9.1
	fmt.Println(fruitMarket("Слон"))
	fmt.Println(fruitMarket("яблоки"))

	// Task 9.2
	numbers := []int{1, 2, 3}

	for _, r1 := range numbers {
		fmt.Printf("v1:%d\n", r1)
	OUT:
		for _, r2 := range numbers {
			fmt.Printf("\tv2:%d\n", r2)
			for _, r3 := range numbers {
				fmt.Printf("\t\tv3:%d\n", r3)
				for i, r4 := range numbers {
					fmt.Printf("\t\t\tv4:%d\n", r4)
					if i == 1 {
						break OUT
					}
				}
			}
		}
	}

	// Task 9.3
	checkFood("яблоко")
	checkFood("карась")
	checkFood("тыква")
}

package main

import (
	"fmt"
)

func main() {

	// Task 7.1
	s := [5]string{"H", "e", "l", "l", "o"}

	fmt.Println(s)

	// Task 7.2
	fruits := [4]string{"яблоко", "груша", "слива", "абрикос"}

	fmt.Println(fruits)

	// Task 7.3
	food := [4]string{"яблоко", "груша", "помидор", "абрикос"}
	food[2] = "персик"
	fmt.Println(food)

	// Task 7.3
	food[2] = "помидор" // Вернул значение в начальное состояние
	fmt.Println(food)
	f := func(mas [4]string, find string) int {
		for i, element := range mas {
			if find == element {
				return i
			}
		}
		return len(mas)
	}
	food[f(food, "помидор")] = "персик"
	fmt.Println(food)

	// Task 7.4
	numbers := []int{5, 2, 8, 3, 1, 9}
	fmt.Println(numbers)

	// Task 7.5
	slice10 := make([]int, 0, 10)
	fmt.Println(slice10)

	// Task 7.6
	slice10 = append(slice10, 4, 1, 8, 9)
	fmt.Println(slice10)

	// Task 7.7
	sliceFirst := []int{1, 2, 3}
	sliceSecond := []int{4, 5, 6}
	bothSlices := make([]int, 0, len(sliceFirst)+len(sliceSecond))
	bothSlices = append(bothSlices, sliceFirst...)
	bothSlices = append(bothSlices, sliceSecond...)
	fmt.Println(bothSlices)

	// Task 7.8
	delSliceElement := func(s []int, index int) []int {
		return append(s[:index], s[index+1:]...)
	}
	sl6 := []int{1, 2, 3, 4, 5, 6}
	sl6 = delSliceElement(sl6, 3)
	fmt.Println(sl6)
}

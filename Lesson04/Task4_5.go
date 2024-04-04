package main

import "fmt"

func main() {

	hello()

}

func hello() {

	defer fmt.Println("Завершение работы")
	fmt.Println("Hello, Go!")

}

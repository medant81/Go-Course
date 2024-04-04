package main

import "fmt"

const sizeFibonacci = 23

func main() {

	printFibonacci(1, 1, 1)

}

func printFibonacci(size int, prevFib int, curFib int) {

	if size > sizeFibonacci-1 {
		fmt.Printf("Fibonacci number %d: %d", size, prevFib)
		return
	}

	fmt.Printf("Fibonacci number %d: %d\n", size, prevFib)
	printFibonacci(size+1, curFib, prevFib+curFib)

}

package main

import "fmt"

type funcType func()

func main() {

	f := hello()
	f()

}

func hello() funcType {

	f := func() {
		fmt.Println("Hello, Go!")
	}
	return f

}

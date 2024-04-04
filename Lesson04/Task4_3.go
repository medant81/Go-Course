package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("Hello, Go!")
	}

	hello(f)

}

func hello(f func()) {
	f()
}

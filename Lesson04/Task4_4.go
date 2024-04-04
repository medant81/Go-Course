package main

type funcType func()

func main() {

	f := hello()
	f()

}

func hello() funcType {

	f := func() {
		println("Hello, Go!")
	}
	return f

}

package main

func main() {

	f := func() {
		println("Hello, Go!")
	}

	hello(f)

}

func hello(f func()) {
	f()
}

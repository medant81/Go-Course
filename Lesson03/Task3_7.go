package main

import "fmt"

const (
	const1 = 1 << iota
	const2
	const3
	const4
	const5
)

func main() {

	fmt.Println(const1, const2, const3, const4, const5)

}

package main

import "fmt"

const myConst = "I'm a global constant "

func main() {

	const myConst = "i'm a local constant"

	fmt.Println(myConst)
}

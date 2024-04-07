package main

import "fmt"

func main() {

	// Task 5.1
	var s *string

	fmt.Println(s, &s)

	// Task 5.2
	var i int

	fmt.Println(i, &i)

	// Task 5.3
	//m := 5
	//p := &m
	//*p = 10

	// Task 5.4
	var b1, b2, b3, b4 byte = 10, 15, 20, 25
	var i1, i2, i3, i4 int64 = 1001, 1002, 1003, 5000
	fmt.Printf("Byte pointers: %p %p %p %p\nByte values: %d %d %d %d\n", &b1, &b2, &b3, &b4, b1, b2, b3, b4)
	fmt.Printf("Int64 pointers: %p %p %p %p\nInt64 values: %d %d %d %d\n", &i1, &i2, &i3, &i4, i1, i2, i3, i4)

}

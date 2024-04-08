package main

import (
	"fmt"
	"unsafe"
)

type phoneCheck struct {
	Name    string
	isPhone bool
}

func main() {

	// Task 5.1
	var s *string = new(string)

	fmt.Println(s, &s)

	// Task 5.2
	var i int

	fmt.Println(i, &i)

	// Task 5.3
	m := 5
	p := &m
	*p = 10

	// Task 5.4
	//var b1, b2, b3, b4 byte = 10, 15, 20, 25
	var b1, b2, b3 byte = 10, 15, 20
	//var i1, i2, i3, i4 int64 = 1001, 1002, 1003, 5000
	var i1, i2, i3 int64 = 1001, 1002, 1003
	var b4 byte = 25
	var i4 int64 = 5000
	fmt.Printf("Byte pointers: %p %p %p %p\nByte values: %d %d %d %d\n", &b1, &b2, &b3, &b4, b1, b2, b3, b4)
	fmt.Printf("Int64 pointers: %p %p %p %p\nInt64 values: %d %d %d %d\n", &i1, &i2, &i3, &i4, i1, i2, i3, i4)
	fmt.Printf("Total memory last byte: %T  %d\n", b4, unsafe.Sizeof(b4))
	var s1 *string = new(string)
	*s1 = "Hello string!"
	fmt.Printf("String pointers: %p %p\nString values: %s %s\n", s, s1, *s, *s1)
	var check1, check2, check3 phoneCheck
	check1.Name = "Sam"
	check1.isPhone = true
	check2.Name = "Bob"
	check2.isPhone = false
	check3.Name = "Jon"
	check3.isPhone = false
	fmt.Printf("Struckt pointers: %p %p %p\n", &check1, &check2, &check3)
	fmt.Printf("Total memory strackt: %T  %d\n", check1, unsafe.Sizeof(check1))

}

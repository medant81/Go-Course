package main

import (
	"fmt"
	"unsafe"
)

type phoneCheck struct {
	Name    string
	isPhone bool
}

type square int

func main() {

	// Task 5.1
	var str *string = new(string)

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
	var str1 *string = new(string)
	*str1 = "Hello string!"
	fmt.Printf("String pointers: %p %p\nString values: %s %s\n", str, str1, *str, *str1)
	var check1, check2, check3 phoneCheck
	check1.Name = "Sam"
	check1.isPhone = true
	check2.Name = "Bob"
	check2.isPhone = false
	check3.Name = "Jon"
	check3.isPhone = false
	fmt.Printf("Struckt pointers: %p %p %p\n", &check1, &check2, &check3)
	fmt.Printf("Total memory strackt: %T  %d\n", check1, unsafe.Sizeof(check1))

	// Task 5.5
	v := 5
	fmt.Println(v) // 5
	change(&v)
	fmt.Println(v) // 10

	// Task 5.6
	sq := square(25)
	fmt.Println(sq)

	// Task 5.7
	s := square(30)
	s += 15
	fmt.Println(s)

	// Task 5.8
	s1 := square(34)
	s1 += 10
	// Немного забежал вперед, но мне с методами понравилось больше
	fmt.Printf("%s\n", s1.print())
	fmt.Printf("%s\n", printSquare(s1))
}

func change(v *int) {
	*v = 10
}

func printSquare(s square) string {
	return fmt.Sprintf("%d м\u00b2", s)
}

func (s square) print() string {
	return fmt.Sprintf("%d м\u00b2", s)
}

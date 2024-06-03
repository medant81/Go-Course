package main

import "fmt"

func main() {

	a := 1
	do(a)

}

func do(v any) {

	a := 10
	/*
		Здесь нужно увеличить значение a на v.
		В случае невозможности приведения к int
		необходимо сообщить об этом и немедленно
		завершить полнение программы.
	*/
	if _, ok := v.(int); !ok {
		panic("Не удалось преобразовать к int")
	}
	a += v.(int)
	fmt.Println(a)

}
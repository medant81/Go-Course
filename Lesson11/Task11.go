package main

import (
	"errors"
	"fmt"
)

type myErr1 string
type myErr2 string
type myErr3 string

type myFirstError struct {
	message string
}

func (e myErr1) Error() string {
	return string(e)
}
func (e myErr2) Error() string {
	return string(e)
}
func (e myErr3) Error() string {
	return string(e)
}

func (s myFirstError) Error() string {
	return s.message
}

func run1(i int) (int, error) {
	if i == 1 {
		return i, nil
	} else {
		return i, myErr1("Ошибка 1:")
	}
}

func run2(i int) (int, error) {
	_, err := run1(i)
	if i == 2 {
		return i, err
	} else {
		if err != nil {
			return i, fmt.Errorf("%w%w", myErr2("Ошибка 2:"), err)
		}
		return i, myErr2("Ошибка 2:")
	}
}

func run3(i int) (int, error) {
	_, err := run2(i)
	if i == 3 {
		return i, err
	} else {
		if err != nil {
			return i, fmt.Errorf("%w%w", myErr3("Ошибка 3:"), err)
		}
		return i, myErr3("Ошибка 3:")
	}
}

func level1() (int, error) {
	return 0, errors.New("Ошибка 1")
}

func level2() (int, error) {
	_, err := level1()
	return 0, fmt.Errorf("Ошибка 2:%w", err)
}

func level3() (int, error) {
	_, err := level2()
	return 0, fmt.Errorf("Ошибка 3:%w", err)
}

func main() {
	fmt.Println("H")
	// Task 11.1
	// Without func used
	err := fmt.Errorf("Ошибка 3:%w", fmt.Errorf("Ошибка 2:%w", errors.New("Ошибка 1")))
	fmt.Println(err)
	// With func used
	_, err = level3()
	fmt.Println(err)

	// Task 11.2
	err = errors.Unwrap(err)
	fmt.Println(err)

	// Task 11.3
	// Сделал вложенные функции run3 -> run2 -> run1
	// В каждой функцие создается ошибка с соответствующим пользовательским типом myErr1, myErr2, myErr3
	//
	// Если в run3 передавать числа 1,2,3 то исключается одна ошибка соответствующая числу
	// Посмотрел код Errorf - разобрал чем отличается 'err' от 'errs'
	_, err = run3(0)
	fmt.Println(err)
	fmt.Println("Есть \"Ошибка 1:\": ", errors.As(err, new(myErr1)))
	_, err = run3(1)
	fmt.Println(err)
	fmt.Println("А теперь нет \"Ошибка 1:\": ", errors.As(err, new(myErr1)))

	// Task 11.3 fix
	err1 := errors.New("Ошибка 1:")
	err2 := fmt.Errorf("Ошибка 2:%w", err1)
	err3 := fmt.Errorf("Ошибка 3:%w", err2)
	fmt.Println("Есть ошибка 1: ", errors.Is(err3, err1))

	// Task 11.4
	// Из предыдущей задачи взял три ошибки
	_, err = run3(0)
	fmt.Println(err)
	fmt.Println("Нет ошибки \"myFirstError\": ", errors.As(err, new(myFirstError)))

	// Task 11.4 fix
	myErr := myFirstError{"Моя ошибка"}
	err4 := fmt.Errorf("%w%w", err, myErr)
	fmt.Println("Нет ошибки \"myFirstError\": ", errors.Is(err4, myErr))
}

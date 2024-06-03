package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Bird interface {
	Sing() string
}

type Duck struct {
	voice string
}

func (d *Duck) Sing() string {
	// Исправление №3: в методе Sing проверим на пустое значение nil
	if d == nil {
		return ""
	}
	return d.voice
}

func main() {
	//var d *Duck = new(Duck) //#1
	// Ошибку выдает, т.к. в строке ниже объявили указатель с типом Duck и выполнили вызов Sing
	// Функция Sing принимает параметр b с типом интерфейса Bird
	// параметру b будет назначен динамический тип Duck и пустое динамическое значение = nil
	// Значение интерфеса будет = nil когда и динамический тип и динамическое значение = nil
	// Поэтому условие b != nil выполняется и происходит вызов Sing() с попыткой получить значение из пустого указателя
	var d *Duck
	// Исправление №1: инициализируем указатель d что бы заполнилось значение по умолчанию (для строки это "")
	d = new(Duck)

	song, err := Sing(d)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(song)
}

func Sing(b Bird) (string, error) {
	// Исправление №2: используя рефлексию проверим что динамическое значение не nil
	/* Было
	if b != nil {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
	*/
	// Стало
	// С помощью библиотеки reflect проверяем значение на nil
	if b != nil && !reflect.ValueOf(b).IsNil() {
		return b.Sing(), nil
	}
	// Исправление №4: приведение к типу Duck и проверка на nil
	// Мне это решение не нравится, т.к. привязано к конкретному типу
	if b != nil && b != (*Duck)(nil) {
		return b.Sing(), nil
	}
	return "", errors.New("Ошибка пения!")
}

/*
Необходимо создать канал с буфером 4, записать в него 2 значения:
«Привет», «буферизованный канал!». Далее необходимо прочитать все значения из канала и вывести в stdout
*/
package main

import "fmt"

func main() {

	ch := make(chan string, 4)

	go func() {
		ch <- "Привет"
		ch <- "буферизованный канал!"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch)

}

/*
Необходимо запустить 5 горутин.
Неиспользуя time.Sleep нужно обеспечить вывод в консоль каждой горутиной своего уникального сообщения.

Например:
горутина: 1
горутина: 2 ...
*/
package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		j := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Горутина: ", j)
		}()
	}
	wg.Wait()
}

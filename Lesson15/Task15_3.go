/*
Используя Mutex, необходимо реализовать счётчик, с которым
параллельно могут работать несколько горутин
*/
package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup

var counter int

func severalGoroutines(c int) {

	for i := 0; i < c; i++ {
		n := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			fmt.Printf("Горутина %d счетчик = %d\n", n, counter)
			mu.Unlock()
		}()
	}

}

func main() {

	wg = sync.WaitGroup{}
	severalGoroutines(10)
	wg.Wait()
	fmt.Println("Итоговое значение счетчика: ", counter)
}

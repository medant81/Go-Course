/*
Используя пакет atomic, необходимо реализовать счётчик, с которым параллельно могут работать несколько горутин.
*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var j int32

func severalGoroutines(c int) {

	for i := 0; i < c; i++ {
		n := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&j, 1)
			fmt.Printf("Горутина %d счетчик = %d\n", n, j)
		}()
	}

}

func main() {

	wg = sync.WaitGroup{}
	severalGoroutines(15)
	wg.Wait()
}

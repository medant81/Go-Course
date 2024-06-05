/*
Необходимо создать функцию start, которая в консоль выводит
некоторое сообщение. Необходимо запустить 10 горутин,
которые будут запускать функцию start и выводить в консоль факт своего запуска,
причём необходимо обеспечить однократный запуск функции start
*/
package main

import (
	"fmt"
	"sync"
)

var loadOnce sync.Once

func start() {
	fmt.Println("Я вывожу некоторое сообщение!")
}

func main() {

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		j := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println("Я горутина: ", j)
			loadOnce.Do(start)
		}()
	}
	wg.Wait()

}

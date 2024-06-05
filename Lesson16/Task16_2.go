/*
Нужно запустить 5 горутин и остановить через 2 секунды
*/
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		j := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Printf("Я горутина %d запущена\n", j)
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Остановилась горутина: ", j)
					return
				default:
					continue
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("А я главная горутина, жду 2 сек.")
		time.Sleep(2 * time.Second)
		fmt.Println("и останавливаю все горутины")
		cancel()
	}()

	wg.Wait()

}

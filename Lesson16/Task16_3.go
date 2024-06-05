/*
Нужно запустить 5 горутин и остановить в некоторое время, которое рассчитывается по формуле: текущий момент + 2 секунды
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
	d := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, d)
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
					fmt.Printf("Ой мне %d пора\n", j)
					return
				default:
					continue
				}
			}
		}()
	}

	wg.Wait()

}

/*
Нужно запустить 5 горутин и остановить, используя контекст с
отменой
*/
package main

import (
	"context"
	"fmt"
	"sync"
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
		fmt.Println("Останавливаю все горутины")
		cancel()
	}()

	wg.Wait()

}

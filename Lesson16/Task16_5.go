/*
Не используя context и буферизованные каналы необходимо написать программу,
которая будет запускать 10 рабочих горутин и одну капризную управляющую горутину.
Каждая рабочая горутина с задержкой в 1 секунду должна выводить в stdout сообщение «сложные вычисления горутины: 1»,
где 1 - порядковый номер горутины.
Управляющая горутина через 3 секунды после своего запуска должна в stdout вывести «ой, всё!»,
после чего рабочие горутины должны в stdout вывести «stop горутина: 1»,
где 1 - порядковый номер горутины, и завершить своё выполнение. В консоли должны увидеть что-то подобное:

сложные вычисления горутины: 3
сложные вычисления горутины: 8
сложные вычисления горутины: 4
сложные вычисления горутины: 9
сложные вычисления горутины: 7
сложные вычисления горутины: 2
сложные вычисления горутины: 0
сложные вычисления горутины: 1
сложные вычисления горутины: 5

сложные вычисления горутины: 6
ой, всё!
сложные вычисления горутины: 3
stop горутина: 3
сложные вычисления горутины: 9
stop горутина: 9
сложные вычисления горутины: 7
stop горутина: 7
сложные вычисления горутины: 4
stop горутина: 4
сложные вычисления горутины: 8
stop горутина: 8
сложные вычисления горутины: 6
stop горутина: 6
сложные вычисления горутины: 2
stop горутина: 2
сложные вычисления горутины: 0
stop горутина: 0
сложные вычисления горутины: 1
stop горутина: 1
сложные вычисления горутины: 5
stop горутина: 5
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := make(chan struct{})

	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		j := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ch:
					fmt.Println("stop горутина: ", j)
					return
				default:
					fmt.Println("сложные вычисления горутина: ", j)
					time.Sleep(time.Second)
					continue
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("Ой, все!")
		close(ch)
	}()

	wg.Wait()
}

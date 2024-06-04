/*
которая в консоль выводит счёт. Горутина настолько утомительна, что в коде поставлено ограничение по времени выполнения программы.
Не меняя логику дочерних горутин, необходимо дописать логику так, чтобы «болтушка» не смогла вывести в консоль свой утомительный счёт.
Вместо пустой болтовни должны увидеть:
завершение работы горутины_1
завершение работы горутины_2
завершение работы главной горутины

Причём логика главной горутины должна отработать в полном объёме, без изменения длительности временных задержек.
*/
/*
Ответ: Если мы объявим пустой канал ch
То 1-й горутине будет нечего читать из канала ch
А 2-я так и не отправит в ch ни чего, т.к. есть проверка ch == nil
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	// Было
	//ch := make(chan int)
	// Стало
	var ch chan int

	stop := make(chan struct{}, 2)

	go func() {
	OUT:
		for {
			select {
			case <-stop:
				break OUT
			case v, ok := <-ch:
				if !ok {
					break OUT
				}
				fmt.Println(v)
			default:
				continue
			}
		}
		fmt.Println("завершение работы горутины_1")
	}()

	go func() {
		var i int
	OUT:
		for {
			i++
			select {
			case <-stop:
				break OUT
			default:
				time.Sleep(time.Second)
				if ch == nil {
					continue
				}
				ch <- i
			}
		}
		fmt.Println("завершение работы горутины_2")
	}()

	time.Sleep(5 * time.Second)
	stop <- struct{}{}
	stop <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("завершение работы главной горутины")
}

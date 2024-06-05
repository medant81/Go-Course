/*
Необходимо реализовать интерфейс

type Meteo interface {
ReadTemp() string
ChangeTemp(v string)
}

Речь про температуру окружающей среды. ReadTemp читает
температуру, ChangeTemp изменяет температуру. Код должен быть потокобезопасным,
т.е. при работе с температурой нескольких параллельных горутин не должно возникать состояние гонки.
*/
package main

import (
	"fmt"
	"sync"
)

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type currentTemp struct {
	valueTemp string
	mu        sync.RWMutex
}

func (t *currentTemp) ReadTemp() string {

	t.mu.RLock()
	defer t.mu.RUnlock()

	return fmt.Sprintf("%sC", t.valueTemp)
}

func (t *currentTemp) ChangeTemp(v string) {

	t.mu.Lock()
	defer t.mu.Unlock()

	t.valueTemp = v

}

func main() {

	var m Meteo = new(currentTemp)

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		m.ChangeTemp("10")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Сейчас температура: ", m.ReadTemp())
	}()

	wg.Wait()
}

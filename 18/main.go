// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
)

type counter struct {
	value int
	m sync.Mutex
}

func(cnt *counter) inc(w *sync.WaitGroup){
	defer w.Done()
	cnt.m.Lock()
	cnt.value++
	cnt.m.Unlock()
}

func main() {
	var (
		w sync.WaitGroup
		cnt counter
	)
	for i := 0; i < 500; i++ {
		w.Add(1)
		go cnt.inc(&w)
	}
	w.Wait()

	fmt.Printf("Cnt value = %v\n", cnt.value)
}
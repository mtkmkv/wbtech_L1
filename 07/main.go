/*Реализовать конкурентную запись данных в map.*/
package main

import (
	"fmt"
	"sync"
)

// Использование Map и sync.Mutex для конкурентной записи
func MapMutex() {
	var (
		w   sync.WaitGroup
		mux sync.Mutex
		m   = make(map[int]int)
		N   = 10
	)

	for i := 0; i < N; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			mux.Lock()
			m[i] = i
			mux.Unlock()
			fmt.Printf("%v\t%v\n", i, m[i])
		}(i)
	}
	w.Wait()
}

// Использование sync.Map для конкурентной записи
func SyncMap() {
	var (
		m sync.Map
		w sync.WaitGroup
		N int = 10
	)

	for i := 0; i < N; i++ {
		w.Add(1)
		go func(i int) {
			defer w.Done()
			m.Store(i, i*i)
			v, _ := m.Load(i)
			fmt.Printf("%v\t%v\n", i, v)
		}(i)
	}
	w.Wait()
}

func main() {
	fmt.Printf("Map & Mutex:\n")
	MapMutex()

	fmt.Printf("SyncMap:\n")
	SyncMap()
}

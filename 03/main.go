/*Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(2^2+3^2+4^2….) с использованием конкурентных вычислений.*/
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func methodWaitGroup(mass []int) (sum int) {

	var (
		w sync.WaitGroup
		m sync.Mutex
	)

	for _, v := range mass {
		w.Add(1)
		go func(v int) {
			defer w.Done()
			m.Lock()
			sum += v * v
			m.Unlock()
		}(v)
	}
	w.Wait()

	return
}

func methodSleep(mass []int) (sum int) {

	var (
		m sync.Mutex
	)

	for _, v := range mass {
		go func(v int) {
			m.Lock()
			sum += v * v
			m.Unlock()
		}(v)
	}
	time.Sleep(100 * time.Millisecond)
	return
}

func methodBufChannel(mass []int) (sum int) {

	ch := make(chan int, len(mass))
	defer close(ch)

	for _, v := range mass {
		go func(v int) {
			ch <- v * v
		}(v)
		sum += <-ch
	}

	return
}

func methodWaitGroupAtomic(mass []int) int {

	var (
		sum int64
		w   sync.WaitGroup
	)

	for _, v := range mass {
		w.Add(1)
		go func(v int) {
			defer w.Done()
			atomic.AddInt64(&sum, int64(v*v))
		}(v)
	}
	w.Wait()

	return int(sum)
}

func methodSleepAtomic(mass []int) int {
	var (
		sum int64
	)
	for _, v := range mass {
		go func(v int) {
			atomic.AddInt64(&sum, int64(v*v))
		}(v)
	}
	time.Sleep(100 * time.Millisecond)
	return int(sum)
}
func main() {
	mass := [...]int{2, 4, 6, 8, 10}
	slice := mass[:]

	start := time.Now()
	fmt.Printf("WaitGroup & Mutex method:\t%v\t%v\n", methodWaitGroup(slice), time.Now().Sub(start).Seconds())
	start = time.Now()
	fmt.Printf("Sleep & Mutex method:\t\t%v\t%v\n", methodSleep(slice), time.Now().Sub(start).Seconds())
	start = time.Now()
	fmt.Printf("Buf Channel method:\t\t%v\t%v\n", methodBufChannel(slice), time.Now().Sub(start).Seconds())
	start = time.Now()
	fmt.Printf("WaitGroup & Atomic method:\t%v\t%v\n", methodWaitGroupAtomic(slice), time.Now().Sub(start).Seconds())
	start = time.Now()
	fmt.Printf("Sleep & Atomic method:\t\t%v\t%v\n", methodSleepAtomic(slice), time.Now().Sub(start).Seconds())
}

/*Написать программу, которая конкурентно рассчитает значение квадратов чисел,
взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func methodWaitGroup(slice []int) {
	var w sync.WaitGroup
	for _, v := range slice {
		w.Add(1)
		go func(v int) {
			defer w.Done()
			fmt.Printf("%v ", v*v)
		}(v)
	}
	w.Wait()
	fmt.Println()
}

func methodSleep(slice []int) {
	for _, v := range slice {
		go func(v int) {
			fmt.Printf("%v ", v*v)
		}(v)
	}
	time.Sleep(100 * time.Millisecond)
	fmt.Println()
}

func methodBufChannel(slice []int) {
	ch := make(chan int)
	defer close(ch)

	for _, v := range slice {
		go func(v int) {
			ch <- v * v
		}(v)
	}
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%v ", <-ch)
	}
}

func main() {
	mass := [...]int{2, 4, 6, 8, 10}
	slice := mass[:]

	fmt.Println("WaitGroup:")
	methodWaitGroup(slice)

	fmt.Println("Sleep 100ms:")
	methodSleep(slice)

	fmt.Println("Buf channel:")
	methodBufChannel(slice)
}

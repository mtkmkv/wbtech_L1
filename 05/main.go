/*Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать.
По истечению N секунд программа должна завершаться.*/

package main

import (
	"fmt"
	"time"
)

func ReadChannel(ch chan int) {
	for v := range ch {
		fmt.Printf("%v ", v)
	}
}

func main() {
	var (
		N  int = 10
		ch     = make(chan int)
	)
	defer close(ch)

	go ReadChannel(ch)

	go func(ch chan int) {
		for i := 0; ; i++ {
			ch <- i * i
			time.Sleep(time.Second)
		}
	}(ch)
	time.Sleep(time.Duration(N) * time.Second)
}

/*Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива,
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout.*/
package main

import "fmt"

func WriteToFirstChannel(ch chan int, slice []int) {
	defer close(ch)

	for _, v := range slice {
		ch <- v
	}
}

func WriteToSecondChannel(ch1, ch2 chan int) {
	defer close(ch2)

	for v := range ch1 {
		ch2 <- 2 * v
	}
}
func main() {
	var (
		ch1   = make(chan int)
		ch2   = make(chan int)
		arr   = [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		slice = arr[:]
	)

	go WriteToFirstChannel(ch1, slice)

	go WriteToSecondChannel(ch1, ch2)

	for v := range ch2 {
		fmt.Printf("%v ", v)
	}
}

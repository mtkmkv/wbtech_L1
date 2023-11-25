/*Реализовать собственную функцию sleep.*/
package main

import (
	"fmt"
	"time"
)

func Sleep(seconds int) {
	//time.After создаёт канал, который отправляет текущее время после заданного интервала
	//(в данном случае, после указанного количества секунд)
	//Оператор <- ожидает сигнал от канала, что приводит к задержке выполнения на указанное время.
	<-time.After(time.Duration(seconds) * time.Second)
}
func main() {
	seconds := 5 
	fmt.Printf("Start. Sleep %v sec.\n", seconds)
	start := time.Now()
	Sleep(seconds)
	fmt.Printf("Time since start: %v sec.\n", time.Since(start))
}

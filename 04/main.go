/*Реализовать постоянную запись данных в канал (главный поток).
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
Необходима возможность выбора количества воркеров при старте.*/

package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"time"
)

func worker(channel chan int, number int, ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("\nРаботник №%v остановлен.", number)
			return
		case data := <-channel:
			fmt.Printf("Работник №%v выполнил: %v\n", number, data)
		}
	}
}

func main() {
	var ch = make(chan int)
	var count int
	fmt.Print("Введите число работников: ")
	_, err := fmt.Scan(&count)
	if err != nil {
		fmt.Print("Ошибка ввода! Попробуйте еще раз: ")
	}

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	for i := 1; i <= count; i++ {
		go worker(ch, i, ctx)
	}

	i := 0
	for {
		select {
		case <-ctx.Done():
			close(ch)
			return
		default:
			ch <- i
			i++
			time.Sleep(500 * time.Millisecond)
		}
	}
}

/*Реализовать бинарный поиск встроенными методами языка.*/
package main

import (
	"fmt"
	"sort"
)

func BinarySearch(slice []int, x int) int {
	var (
		low  int = 0
		high int = len(slice) - 1
	)

	for low <= high {
		mid := (low + high) / 2
		guess := slice[mid]
		if guess == x {
			return mid
		}
		if guess > x {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func main() {
	var (
		slice     = []int{1, 4, 6, 8, 9, 5, 0, 10, 34, 12, 76, 45, 98}
		x     int = 5
	)

	sort.Slice(slice, func(i, j int) bool {
		return slice[i] < slice[j]
	})
	fmt.Printf("Отсортированный входной слайс - %v\n", slice)
	res := BinarySearch(slice, x)

	if res == -1 {
		fmt.Printf("Слайс не содержит элемента %v", x)
	} else {
		fmt.Printf("Индекс элемента %v в слайсе - %v", x, res)
	}
}

/*16.	Реализовать быструю сортировку массива (quicksort) встроенными методами языка.*/
package main

import "fmt"

func partOfQuickSort(slice []int, left, right int) int {
	pivot := slice[(left+right)/2]
	for left <= right {
		for slice[left] < pivot {
			left++
		}

		for slice[right] > pivot {
			right--
		}

		if left <= right {
			slice[left], slice[right] = slice[right], slice[left]
			left++
			right--
		}
	}
	return left
}
func quickSort(slice []int, start, end int) {
	if start >= end {
		return
	}
	rightStart := partOfQuickSort(slice, start, end)

	quickSort(slice, start, rightStart-1)
	quickSort(slice, rightStart, end)
}

func main() {
	slice := []int{4, 6, 7, 3, 1, 4, 99, 8, 7, 0, -1}
	fmt.Printf("Исходный слайс: %v\n", slice)
	quickSort(slice, 0, len(slice)-1)
	fmt.Printf("Отсортированный: %v\n", slice)
}

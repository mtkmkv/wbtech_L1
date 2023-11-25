/*Удалить i-ый элемент из слайса.*/
package main

import "fmt"

// Удаление элемента с изменением порядка
func DeleteElemChangeOrder(slice []int, i int) []int {
	slice[i] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

// Удаление элемента без изменения порядка
func DeleteElemNoChangeOrder(slice []int, i int) []int {
	return append(slice[:i], slice[i+1:]...)
}

func main() {
	slice := []int{1, 3, 5, 4, 6, 7}
	indDel := 2 // Удаляем второй элемент из слайса

	fmt.Printf("Изначальный слайс: %v\n", slice)
	slice = DeleteElemChangeOrder(slice, indDel)
	fmt.Printf("Удаление элемента с индексов %v с изменением порялка:\t%v\n", indDel, slice)

	slice = []int{1, 3, 5, 4, 6, 7}
	slice = DeleteElemNoChangeOrder(slice, indDel)
	fmt.Printf("Удаление элемента с индексов %v с без изменения порялка:\t%v\n", indDel, slice)

}

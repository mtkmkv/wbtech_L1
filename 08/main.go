/*8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.*/
package main

import "fmt"

func SetBit(num int64, numBit uint, valueBit int64) int64 {
	if valueBit == 1 {
		// Устанавливаем бит в 1
		return num | (1 << numBit)
	}
	// Устанавливаем бит в 0
	return num & ^(1 << numBit)
}

func PrintResult(num, res int64) {
	fmt.Printf("Число: %v\tв дв. системе: %b\nОтвет: %v\tв дв. системе: %b\n", num, num, res, res)
}
func main() {
	var (
		num1 int64 = 1024
		bit1 uint  = 4

		num2 int64 = 2023
		bit2 uint  = 6
	)

	PrintResult(num1, SetBit(num1, bit1, 1))
	PrintResult(num2, SetBit(num2, bit2, 0))
}

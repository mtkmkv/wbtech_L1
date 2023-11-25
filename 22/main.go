/*22.	Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20 = 1048576.*/
package main

import (
	"fmt"
	"math/big"
)

// Сложение
func sum(first, second *big.Int) *big.Int {
	return new(big.Int).Add(first, second)
}

// Вычитание
func sub(first, second *big.Int) *big.Int {
	return new(big.Int).Sub(first, second)
}

// Умножение
func mul(first, second *big.Int) *big.Int {
	return new(big.Int).Mul(first, second)
}

// Деление
func div(first, second *big.Int) *big.Int {
	return new(big.Int).Div(first, second)
}

func main() {
	firstBigInt := new(big.Int)
	firstBigInt.SetString("7777777777", 10)

	secondBigInt := new(big.Int)
	secondBigInt.SetString("1111111111", 10)

	fmt.Printf("%v + %v = %v\n", firstBigInt, secondBigInt, sum(firstBigInt, secondBigInt))
	fmt.Printf("%v - %v = %v\n", firstBigInt, secondBigInt, sub(firstBigInt, secondBigInt))
	fmt.Printf("%v - %v = %v\n", secondBigInt, firstBigInt, sub(secondBigInt, firstBigInt))
	fmt.Printf("%v * %v = %v\n", firstBigInt, secondBigInt, mul(firstBigInt, secondBigInt))
	fmt.Printf("%v / %v = %v\n", firstBigInt, secondBigInt, div(firstBigInt, secondBigInt))

}

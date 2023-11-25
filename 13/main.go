/*Поменять местами два числа без создания временной переменной*/
package main

import "fmt"

func main() {
	var (
		a int = 10
		b int = 15
	)
	fmt.Printf("a = %v; b = %v.\n", a, b)
	a, b = b, a
	fmt.Printf("a = %v; b = %v.\n", a, b)
}

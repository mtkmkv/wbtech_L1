/*
К каким негативным последствиям может привести данный фрагмент кода, и как это исправить?

var justString string

	func someFunc() {
	  v := createHugeString(1 << 10)
	  justString = v[:100]
	}

	func main() {
	  someFunc()
	}

Приведите корректный пример реализации.
*/
package main

import (
	"fmt"
	"strconv"
)

func createHugeString(value int) (str string) {

	for i := 0; i < value; i++ {
		if i%10 == 0 {
			str += "|" + strconv.Itoa(i%10)
		} else {
			str += strconv.Itoa(i % 10)
		}

	}
	return
}

var justString string

// так как строка - это массив байтов, то и срез будет идти не посимвольно, а побайтно
// поэтому перед тем, как брать срез, преобразуем строку в срез рун
func someFunc() {
	v := createHugeString(1 << 10)
	justString = string([]rune(v)[:100])
	fmt.Printf("\"%v\"", justString)
}

func main() {
	someFunc()
}

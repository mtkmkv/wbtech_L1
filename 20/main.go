/*
Разработать программу, которая переворачивает слова в строке.
Пример: «snow dog sun — sun dog snow».
*/
package main

import (
	"fmt"
	"strings"
)

func ReverseWords(str string) string {
	var (
		strSlice = strings.Split(str, " ")
		space    string
		resStr   string
	)
	for i := len(strSlice) - 1; i >= 0; i-- {
		resStr += space + strSlice[i]
		space = " "
	}

	return resStr
}

func main() {
	str := "snow dog sun"
	fmt.Printf("Исходная строка:\t%v\nПеревернутая строка:\t%v\n", str, ReverseWords(str))
}

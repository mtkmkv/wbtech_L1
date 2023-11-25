/*
 26. Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc).
    Функция проверки должна быть регистронезависимой.
    Например:
    abcd — true
    abCdefAaf — false
    aabcd — false
*/
package main

import (
	"fmt"
	"strings"
)

func IsSymbolsAreUnique(str string) bool {
	str = strings.ToLower(str)
	m := make(map[rune]struct{})

	for _, v := range str {
		if _, ok := m[v]; ok {
			return false
		}	else {
			m[v] = struct{}{}
		}
	}
	return true
}
func main() {
	str1 := "abcd"
	fmt.Printf("%v - %v\n", str1, IsSymbolsAreUnique(str1))
	str2 := "abCdefAaf"
	fmt.Printf("%v - %v\n", str2, IsSymbolsAreUnique(str2))
	str3 := "aabcd"
	fmt.Printf("%v - %v\n", str2, IsSymbolsAreUnique(str3))
}

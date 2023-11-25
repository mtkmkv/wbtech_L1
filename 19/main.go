/*19.	Разработать программу, которая переворачивает подаваемую на ход строку
(например: «главрыба — абырвалг»). Символы могут быть unicode.*/
package main

import "fmt"

func ReverseStr(str string) string {
	runeStr := []rune(str)
	runeResult := []rune{}
	for _, v := range runeStr {
		runeResult = append([]rune{v}, runeResult...)
	}
	return string(runeResult)
}
func main() {
	str := "qwerty"
	fmt.Printf("%s - %s", str, ReverseStr(str))
}

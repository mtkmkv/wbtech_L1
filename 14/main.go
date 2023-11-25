/*
Разработать программу, которая в рантайме способна определить тип переменной:
int, string, bool, channel из переменной типа interface{}.
*/
package main

import (
	"fmt"
	"reflect"
)

func usualPrint(v interface{}) {
	fmt.Printf("%v\t-\t%T\n", v, v)
}

func reflectPrint(v interface{}) {
	fmt.Printf("%v\t-\t%v\n", v, reflect.TypeOf(v))
}

func main() {
	arr := []interface{}{1, int64(1), "Hello", true, 2.4, 1 + 2i, make(chan int), make(chan string), struct{}{}}

	fmt.Println("Вывод с использованием возможностей форматирования через %Т:")
	for _, v := range arr {
		usualPrint(v)
	}

	fmt.Println("Вывод с reflect.TypeOf():")
	for _, v := range arr {
		reflectPrint(v)
	}
}

/*Дана структура Human (с произвольным набором полей и методов).
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования)*/

package main

import (
	"fmt"
)

// Родительская структура
type Human struct {
	Name string
	Age  uint8
}

// Метод Human, доступный объектам типа Action
func (h Human) NameFunc() {
	fmt.Printf("My name is %s.\n", h.Name)
}

// Метод Human
func (h Human) Hello() {
	fmt.Printf("Hello, my name is %s. My age is %d.\n", h.Name, h.Age)
}

// Дочерняя структура
type Action struct {
	Human
	Job string
}

// Переопределенный метод Human
func (a Action) Hello() {
	fmt.Printf("Hello, my name is %s. My age is %d. My job is %s.\n", a.Name, a.Age, a.Job)
}

// Метод Action
func (a Action) MyJob() {
	fmt.Printf("My name is %s. I am a %s.\n", a.Name, a.Job)
}
func main() {
	human := Human{"Jame", 19}
	action := Action{Human{"Robert", 21}, "baker"}

	human.NameFunc() // Метод Human
	human.Hello()    // Метод Human
	//	human.Action() 	// Ошибка! Метод дочерней структуры не доступен родительской

	action.NameFunc() // Метод Human, доступный объектам типа Action
	action.Hello()    // Переопределенный метод Huma
	action.MyJob()    // Метод Action
}

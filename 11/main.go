// Реализовать пересечение двух неупорядоченных множеств.

package main

import "fmt"

type myStruct struct{}

func main() {
	var (
		s1 = map[int]myStruct{
			1: {},
			2: {},
			3: {},
		}
		s2 = map[int]myStruct{
			3: {},
			4: {},
			5: {},
		}
		intersection = make(map[int]myStruct)
	)
	for k, _ := range s1 {
		if _, ok := s2[k]; ok {
			intersection[k] = myStruct{}
		}
	}
	fmt.Printf("Первое множество: %v\nВторое множество: %v\nИх пересечение: %v", s1, s2, intersection)
}

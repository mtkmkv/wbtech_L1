/*12.	Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.*/
package main

import "fmt"

func main() {
	list := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]int)

	for _, item := range list {
		if _, ok := m[item]; !ok {
			m[item] = 1
		} else {
			m[item]++
		}
	}
	fmt.Printf("Множество: %v\n", m)
}

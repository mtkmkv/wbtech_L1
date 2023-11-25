/*
24.	Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.
*/
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (p *Point) set(x, y float64) {
	p.x, p.y = x, y
}
func findDistance(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow((p2.x-p1.x), 2) + math.Pow((p2.y-p1.y), 2))
}

func printResult(p1, p2, res interface{}) {
	fmt.Printf("Расстояние между %v и %v равно %v\n", p1, p2, res)
}
func main() {
	var p1, p2 Point
	p1.set(0.0, 0.0)
	p2.set(3.0, 3.0)

	printResult(p1, p2, findDistance(p1, p2))
}

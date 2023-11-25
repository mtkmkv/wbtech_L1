/*Дана последовательность температурных колебаний: -25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5.
Объединить данные значения в группы с шагом в 10 градусов. Последовательность в подмножноствах не важна.
Пример: -20:{-25.0, -27.0, -21.0}, 10:{13.0, 19.0, 15.5}, 20: {24.5}, etc.*/

package main

import (
	"fmt"
	"strconv"
)

func groupValues(value float64) string {
	if value < 0 && value > -10 {
		return "-0"
	} else {
		return strconv.Itoa(int(value/10) * 10)
	}
}
func main() {
	m := make(map[string][]float64)
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5, -5.0, -7.9, -10.5, -10.0, 23.9, 6.7, 8.1}
	for _, v := range temperatures {
		group := groupValues(v)

		if value, ok := m[group]; ok {
			value = append(value, v)
			m[group] = value
		} else {
			m[group] = []float64{v}
		}
	}
	fmt.Println(m)
}

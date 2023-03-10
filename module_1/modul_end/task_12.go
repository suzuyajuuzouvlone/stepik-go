// По данному числу N распечатайте все целые значения степени двойки, не превосходящие N, в порядке возрастания.
//
// # Входные данные
//
// Вводится натуральное число.
//
// # Выходные данные
//
// Выведите ответ на задачу.
// код с gita ((
package main

import (
	"fmt"
	"math"
)

func main() {
	var a int
	fmt.Scan(&a)

	p := float64(0)
	for {
		if math.Pow(2, p) <= float64(a) {
			fmt.Printf("%v ", math.Pow(2, p))
			p++
		} else {
			break
		}
	}
}

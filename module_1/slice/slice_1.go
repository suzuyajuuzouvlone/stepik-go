// Напишите программу, принимающая на вход число N (N \geq 4)N(N≥4), а затем NN целых чисел-элементов среза. На вывод нужно подать 4-ый (3 по индексу) элемент данного среза.

package main

import "fmt"

func main() {
	var inputNumber int

	_, _ = fmt.Scan(&inputNumber)

	inputArray := make([]int, inputNumber)

	_, _ = fmt.Scan(&inputArray)
	for i := 0; i < inputNumber; i++ {
		_, _ = fmt.Scan(&inputArray[i])
	}

	fmt.Println(inputArray[3])
}

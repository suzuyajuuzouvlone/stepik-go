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

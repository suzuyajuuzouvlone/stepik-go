// Из натурального числа удалить заданную цифру.
//
// # Входные данные
//
// Вводятся натуральное число и цифра, которую нужно удалить.
//
// # Выходные данные
//
// Вывести число без заданных цифр.
//
// Sample Input:
//
// 38012732
// 3
// Sample Output:
//
// 801272
package main

import "fmt"

func main() {
	var number, del int
	fmt.Scan(&number, &del)
	result, rank := 0, 1
	for number != 0 {
		reminder := number % 10
		if reminder != del {
			result += reminder * rank
			rank *= 10
		}
		number /= 10
	}
	fmt.Println(result)
}

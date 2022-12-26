// На вход подается целое число. Необходимо возвести в квадрат каждую цифру числа и вывести получившееся число.
//
//Например, у нас есть число 9119. Первая цифра - 9. 9 в квадрате - 81. Дальше 1. Единица в квадрате - 1. В итоге получаем 811181
//
//Sample Input:
//
//9119
//Sample Output:
//
//811181
package main

import "fmt"

func main() {
	var num, result int
	fmt.Scan(&num)

	rank := 1
	for num != 0 {
		reminder := num % 10
		square := reminder * reminder
		result += square * rank
		if square > 9 {
			rank *= 100
		} else {
			rank *= 10
		}
		num /= 10
	}
	fmt.Println(result)
}

// Дана строка, содержащая только арабские цифры. Найти и вывести наибольшую цифру.
//
// # Входные данные
//
// Вводится строка ненулевой длины. Известно также, что длина строки не превышает 1000 знаков и строка содержит только арабские цифры.
//
// # Выходные данные
//
// Выведите максимальную цифру, которая встречается во введенной строке.
//
// Sample Input:
//
// 1112221112
// Sample Output:
//
// 2
package main

import "fmt"

func main() {
	var number string
	fmt.Scan(&number)
	runes := []rune(number)
	max := string(runes[0])
	for i := 0; i < len(runes); i++ {
		if max < string(runes[i]) {
			max = string(runes[i])
		}
	}
	fmt.Println(max)
}

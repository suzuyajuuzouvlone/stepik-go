// Дано трехзначное число. Найдите сумму его цифр.
//
//Формат входных данных
//На вход дается трехзначное число.
//
//Формат выходных данных
//Выведите одно целое число - сумму цифр введенного числа.
//
//Sample Input:
//
//745
//Sample Output:
//
//16
package main

import "fmt"

func main() {
	var a, sum int
	fmt.Scan(&a)
	if a >= 100 && a <= 999 {
		first_int := a / 100
		second_int := a / 10 % 10
		third_int := a % 10
		sum = first_int + second_int + third_int
	}
	fmt.Println(sum)
}

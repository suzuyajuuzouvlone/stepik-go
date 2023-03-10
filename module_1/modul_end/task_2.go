// Дано трехзначное число. Переверните его, а затем выведите.
//
//Формат входных данных
//На вход дается трехзначное число, не оканчивающееся на ноль.
//
//Формат выходных данных
//Выведите перевернутое число.
//
//Sample Input:
//
//843
//Sample Output:
//
//348
package main

import "fmt"

func main() {
	var a, rev int
	fmt.Scan(&a)
	if a >= 100 && a <= 999 {
		f := a / 100
		s := a / 10 % 10
		t := a % 10
		rev = t*100 + s*10 + f
	}
	fmt.Println(rev)
}

//	На вход дается строка, из нее нужно сделать другую строку, оставив только нечетные символы (считая с нуля)
//
// Sample Input:
//
// ihgewlqlkot
// Sample Output:
//
// hello
package main

import "fmt"

func main() {
	var s, even string
	fmt.Scan(&s)
	sRunes := []rune(s)
	length := len(sRunes)
	for i := 0; i < length; i++ {
		if i%2 == 1 {
			even += string(sRunes[i])
		}
	}
	fmt.Println(even)
}

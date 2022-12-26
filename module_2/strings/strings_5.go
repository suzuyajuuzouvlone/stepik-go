// Дается строка. Нужно удалить все символы, которые встречаются более одного раза и вывести получившуюся строку
//
// Sample Input:
//
// zaabcbd
// Sample Output:
//
// zcd
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s, zxc string
	fmt.Scan(&s)
	sRunes := []rune(s)
	for i := range sRunes {
		if strings.Count(s, string(sRunes[i])) > 1 {
			continue
		} else {
			zxc += string(sRunes[i])
		}
	}
	fmt.Println(zxc)
}

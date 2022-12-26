// Поменяйте местами значения переменных на которые ссылаются указатели. После этого переменные нужно вывести.
// ВАЖНО: Считайте что пакет main уже объявлен, а также функция main() уже есть.
//
// Sample Input:
//
// 2 4
// Sample Output:
//
// 4 2
package main

import "fmt"

func test(x1 *int, x2 *int) {
	s := *x1
	*x1 = *x2
	*x2 = s
	fmt.Println(*x1, *x2)
}

func main() {
	var x1, x2 int
	fmt.Scan(&x1, &x2)
	test(&x1, &x2)
}

// Напишите функцию sumInt, принимающую переменное количество аргументов типа int, и возвращающую количество полученных функцией аргументов и их сумму. Пакет "fmt" уже импортирован, функция и пакет main объявлены.
package main

import "fmt"

func sumInt(nums ...int) (count, sum int) {
	for _, value := range nums {
		sum += value
		count++
	}
	return count, sum
}

func main() {
	a, b := sumInt(1, 0, 3)
	fmt.Println(a, b)
}

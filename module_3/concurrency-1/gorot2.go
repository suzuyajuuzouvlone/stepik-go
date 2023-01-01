// Напишите функцию которая принимает канал и строку. Необходимо отправить эту же строку в заданный канал 5 раз, добавив к ней пробел.
package main

import (
	"fmt"
)

func task2(ch chan string, s string) {
	for i := 1; i <= 5; i++ {
		ch <- s + " "
	}
}

func main() {
	ch := make(chan string)

	go task2(ch, "test")

	fmt.Printf("Output: ")
	for i := 0; i < 5; i++ {
		fmt.Printf("%#v", <-ch)
	}
	fmt.Printf("\n")

}

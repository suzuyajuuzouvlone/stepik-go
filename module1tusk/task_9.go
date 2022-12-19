package main

import "fmt"

func main() {
	var n, ck int
	fmt.Scan(&n)

	if n > 10 {
		ck = 1 + ((n - 1) % 9)
	}
	fmt.Println(ck)
}

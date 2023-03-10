// Ваша задача сделать проверку подходит ли пароль вводимый пользователем под заданные требования. Длина пароля должна быть не менее 5 символов, он должен содержать только арабские цифры и буквы латинского алфавита. На вход подается строка-пароль. Если пароль соответствует требованиям - вывести "Ok", иначе вывести "Wrong password"
//
// Sample Input:
//
// fdsghdfgjsdDD1
// Sample Output:
//
// Ok
package main

import (
	"fmt"
	"unicode"
)

func IsValid(p string) bool {
	runes := []rune(p)
	length := len(runes)
	var isValid bool
	if length < 5 {
		isValid = false
	} else {
		for _, r := range runes {
			if !(unicode.Is(unicode.Latin, r) || unicode.Is(unicode.Digit, r)) {
				isValid = false
				break
			} else {
				isValid = true
			}
		}
	}
	return isValid
}

func main() {
	var password string
	fmt.Scan(&password)
	if IsValid(password) {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}

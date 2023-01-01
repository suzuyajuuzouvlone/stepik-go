// Представьте что вы работаете в большой компании где используется модульная архитектура. Ваш коллега написал модуль с какой-то логикой (вы не знаете) и передает в вашу программу какие-то данные. Вы же пишете функцию которая считывает две переменных типа string ,  а возвращает число типа int64 которое нужно получить сложением этих строк.
//
// Но не было бы так все просто, ведь ваш коллега не пишет на Go, и он зол из-за того, что гоферам платят больше. Поэтому он решил подшутить над вами и подсунул вам свинью. Он придумал вставлять мусор в строки перед тем как вызывать вашу функцию.
package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func CleanFromTrash(s string) int64 {
	sArr := []rune(s)
	for _, r := range sArr {
		if !unicode.IsDigit(r) {
			s = strings.ReplaceAll(s, string(r), "")
		}
	}
	result, _ := strconv.ParseInt(s, 10, 64)
	return result
}

func adding(s1, s2 string) int64 {
	first := CleanFromTrash(s1)
	second := CleanFromTrash(s2)

	return first + second
}

func main() {
	var a, b string
	fmt.Scan(&a, &b)

	fmt.Println(adding(a, b))
}

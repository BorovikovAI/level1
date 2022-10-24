// Разработать программу, которая проверяет, что все символы в строке уникальные
// (true — если уникальные, false etc). Функция проверки должна быть регистронезависимой.

// Например:
// abcd — true
// abCdefAaf — false
// aabcd — false

package main

import (
	"fmt"
	"strings"
)

func StringCheck(str string) bool {
	exist := make(map[string]bool)

	str = strings.ToLower(str)

	for _, v := range str {
		if exist[string(v)] == true {
			return false
		}

		exist[string(v)] = true
	}

	return true
}

func main() {
	srt1 := "abcd"
	str2 := "abCdefAaf"
	str3 := "aabcd"
	fmt.Println("abcd -", StringCheck(srt1))
	fmt.Println("abCdefAaf -", StringCheck(str2))
	fmt.Println("aabcd", StringCheck(str3))
}

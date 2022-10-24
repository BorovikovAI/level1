// Разработать программу, которая переворачивает
// подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

package main

import (
	"fmt"
)

func main() {
	s := "s''123;asввыфвфы1231323э123эd;aweqd"
	fmt.Println(ReverseStr(s))
}

func ReverseStr(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

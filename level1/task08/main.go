//Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

package main

import (
	"fmt"
)

func main() {
	// 10010
	var num int64 = 18
	fmt.Printf("num: %b\n", num)
	newNum := SetBitToOne(num, 0)
	fmt.Printf("Set i=%d bit to 1: %b \n", 0, newNum)
	fmt.Printf("Set i=%d bit to 0: %b \n", 1, SetBitToZero(newNum, 1))
}

func SetBitToOne(num int64, i int) int64 {
	num |= (1 << i)
	return num
}

func SetBitToZero(num int64, i int) int64 {
	num &^= (1 << i)
	return num
}

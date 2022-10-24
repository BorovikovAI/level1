//Поменять местами два числа без создания временной переменной.

package main

import (
	"fmt"
)

func main() {
	x := 45
	y := 67
	fmt.Println(x, y)
	x, y = y, x
	fmt.Println(x, y)
}

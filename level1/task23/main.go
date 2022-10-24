// Удалить i-ый элемент из слайса.

package main

import (
	"fmt"
)

func RemoveElementFromSlice(slice []int, i int) []int {
	return append(slice[0:i], slice[i+1:]...) //при помощи append
}

func main() {
	slice := []int{23, 45, 78, 56, 453}
	fmt.Println(slice)
	fmt.Println(RemoveElementFromSlice(slice, 3))
}

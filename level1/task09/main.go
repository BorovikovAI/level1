// Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, во второй —
// результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.

package main

import (
	"fmt"
)

func main() {
	chanIn := make(chan int)
	chanOut := make(chan int)

	arr := []int{12, 23, 34, 45, 56, 67, 78, 8989}

	go func() {
		defer close(chanIn)
		for i := 0; i < len(arr); i++ {
			chanIn <- arr[i]
		}
	}()

	go func() {
		defer close(chanOut)
		var n int
		for i := 0; i < len(arr); i++ {
			n = <-chanIn
			chanOut <- n * 2
		}
	}()

	for v := range chanOut {
		fmt.Println(v)
	}
}

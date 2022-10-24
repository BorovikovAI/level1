//Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых
//из массива (2,4,6,8,10) и выведет их квадраты в stdout.

package main

import (
	"fmt"
	"sync"
)

func main() {
	nums := []int{2, 4, 6, 8, 10}
	GetAllSquares(nums...)
}

func GetAllSquares(nums ...int) {
	//for waiting all goroutines done
	var wg sync.WaitGroup
	//adding the number of goroutines
	wg.Add(len(nums))

	for _, num := range nums {
		go func(num int) {
			//decrementing the number of goroutines
			defer wg.Done()
			fmt.Println(Square(num))
		}(num)
	}
	//waiting til all goroutines are done
	wg.Wait()
}

func Square(num int) int {
	return num * num
}

//Дана последовательность чисел: 2,4,6,8,10
//Найти сумму их квадратов (2^2 + 3^2 + 4^2 ...) с использованием конкурентных вычислений

package main

import (
	"fmt"
	"sync"
)

func main() {
	//rcvCh := make(chan int)
	nums := []int{2, 4, 6, 8, 10}
	GetAllSquares(nums...)
}

func GetAllSquares(nums ...int) {
	//channel for squares
	rcvCh := make(chan int)
	defer close(rcvCh)
	//for waiting all goroutines done
	var wg sync.WaitGroup
	//adding the number of goroutines
	wg.Add(len(nums))

	var sum int

	for _, num := range nums {
		go func(num int) {
			//decrementing the number of goroutines
			defer wg.Done()
			fmt.Println(Square(num))
			//sending result to channel
			rcvCh <- Square(num)
		}(num)
		//getting data from channel
		sum += <-rcvCh
	}
	//waiting til all goroutines are done
	wg.Wait()
	fmt.Println("sum:", sum)
}

func Square(num int) int {
	return num * num
}

// Реализовать постоянную запись данных в канал (главный поток).
// Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.
// Необходима возможность выбора количества воркеров при старте.

// Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров.

package main

import (
	"fmt"
	"runtime"
)

func main() {
	workers, someData := 0, 0
	//get workers num from stdin
	fmt.Scanln(&workers)
	//channel to workers
	ch := make(chan int)
	defer close(ch)

	for n := 0; n < workers; n++ {
		//starting N-workers goroutines
		go func(n int, ch chan int) {
			//reading data from channel
			for v := range ch {
				fmt.Println("Worker:", n, "\tData from channel:", v)
				//to allow other goroutines run
				runtime.Gosched()
			}
		}(n, ch)
	}

	for {
		if someData >= 20 {
			break
		}
		//sending some data into the channel
		someData++
		ch <- someData
	}
	//waiting all goroutines are done
	fmt.Scanln()
}

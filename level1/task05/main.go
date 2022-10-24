//Разработать программу, которая будет последовательно отправлять значения в канал,
//а с другой стороны канала — читать. По истечению N секунд программа должна завершаться.

package main

import (
	"fmt"
	"time"
)

func main() {
	var N, someData int = 0, 0
	//get N seconds
	fmt.Scanln(&N)

	ch := make(chan int)
	defer close(ch)

	//goroutine to read data from the channel
	go func(ch chan int) {
		for {
			//getting data from channel
			for v := range ch {
				fmt.Println(v)
			}
		}
	}(ch)

	for {
		if someData >= 20 {
			break
		}
		//sending some data into the channel
		someData++
		ch <- someData
	}

	//sleep N seconds
	time.Sleep(time.Duration(N) * time.Second)
}

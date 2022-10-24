//Реализовать все возможные способы остановки выполнения горутины.

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	RecChanAndQuitChan()
	fmt.Println("")
	BasicWayToCloseTheChannel()
	fmt.Println("")
	UsingTheContext()
}

func RecChanAndQuitChan() {
	quit := make(chan struct{})
	ch := make(chan int)

	go func() {

		for {
			//select{} for listening "ch" & "quit"
			select {
			case flag := <-quit:
				//closening both channels
				close(ch)
				close(quit)
				fmt.Println("quit status:", flag)
				//return if "quit" was used to stop goroutine
				return
			case n := <-ch:
				fmt.Println(n)
			default:
				fmt.Println("...")
			}
		}
	}()

	var data int = 0

	//permanently sending data to the channel
	for {
		data++
		ch <- data
		if data >= 10 {
			quit <- struct{}{}
			fmt.Println("end")
			return
		}
	}
}

func BasicWayToCloseTheChannel() {
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan int)

	go func() {
		for {
			v, ok := <-ch
			//if the channel is closed
			if !ok {
				//decrementing wait group
				wg.Done()
				//printing channel status
				fmt.Println("channel status:", ok)
				return
			}
			fmt.Println(v)
		}
	}()
	//sending data to the channel
	ch <- 678
	//closening the channel
	close(ch)
	//waiting til goroutine stoped
	wg.Wait()
	fmt.Println("end")
}

func UsingTheContext() {
	ch := make(chan int)
	//context for canceling the channel
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for {
			select {
			//when channel is canceled
			case <-ctx.Done():
				return
			//when some data was sent to the channel
			case n := <-ch:
				//printing recieved data
				fmt.Println(n)
			default:
				fmt.Println("...")
			}
		}
	}(ctx)

	go func() {
		time.Sleep(50 * time.Millisecond)
		//canceling the channel
		cancel()
	}()

	var someData int = 0

	for {
		if someData >= 10 {
			break
		}
		//sending some data into the channel
		someData++
		ch <- someData
	}

	fmt.Println("end")
}

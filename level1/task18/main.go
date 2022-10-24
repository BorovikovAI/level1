// Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде.
// По завершению программа должна выводить итоговое значение счетчика.

package main

import (
	"fmt"
	"sync"
	"time"
)

type Counter struct {
	Incr int
	sync.Mutex
}

func (c *Counter) Incrementation() {
	c.Mutex.Lock()
	defer c.Mutex.Unlock()
	c.Incr++
}

func main() {
	incr := Counter{Incr: 0}

	var wg sync.WaitGroup
	var N, Msec int = 100, 50

	wg.Add(N)

	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(Msec) * time.Millisecond)
			incr.Incrementation()
			time.Sleep(time.Duration(Msec) * time.Millisecond)
			incr.Incrementation()
			time.Sleep(time.Duration(Msec) * time.Millisecond)
			incr.Incrementation()
		}()
	}

	wg.Wait()
	fmt.Println(incr.Incr)
}

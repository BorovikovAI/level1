//Реализовать собственную функцию sleep.

package main

import (
	"fmt"
	"time"
)

func MySleep(sec int) {
	<-time.After(time.Second * time.Duration(sec))
}

func main() {
	fmt.Println("start:", time.Now())
	MySleep(10)
	fmt.Println("done:", time.Now())
}

// Реализовать конкурентную запись данных в map

package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	myMap := make(map[int]string)

	rwMtx := &sync.Mutex{}

	wg := &sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go WriteInMap(i, rwMtx, wg, myMap)
	}

	wg.Wait()
	fmt.Println(myMap)
}

func WriteInMap(i int, rwMtx *sync.Mutex, wg *sync.WaitGroup, myMap map[int]string) {

	defer wg.Done()
	rwMtx.Lock()
	defer rwMtx.Unlock()
	myMap[i] = "string_" + strconv.Itoa(i)
	fmt.Println(myMap[i])
}

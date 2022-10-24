package main

import (
	"fmt"
	"sync"
)

// "p" in this func is not "p" in main
func update(p *int) *int {
	b := 2
	p = &b
	return p
}

func main10() {
	var (
		a = 1
		p = &a
	)
	fmt.Println(*p)
	// return updated "p"
	newp := update(p)
	fmt.Println(*p)
	// print new "p"
	fmt.Println(*newp)
}

func main11() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		// нужно передавать адрес, иначе происходит копирование объекта
		go func(wg *sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(&wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

func main12() {
	n := 0
	if true {
		// в этом блоке была создана новая переменная, которая была удалена после его завершения
		n = 1
		n++
	}
	fmt.Println(n)
}

func someAction(v []int8, b int8) {
	//меняет значение в ячейке с индексом 0
	v[0] = 100
	//добавляем новый элемент в слайс, тем самым создаем новый локальный массив, на который указывает слайс
	v = append(v, b)
}

func main13() {
	var a = []int8{1, 2, 3, 4, 5}
	someAction(a, 6)
	fmt.Println(a)
}

func main() {
	slice := []string{"a", "a"}
	func(slice []string) {
		slice = append(slice, "a") //создается новый массив на который ссылается копия слайса
		slice[0] = "b"
		slice[1] = "b"
		fmt.Println(slice)
	}(slice)
	fmt.Println(slice)
}

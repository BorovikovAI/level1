// Дана структура Human (с произвольным набором полей и методов).
// Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования).
package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
}

func (h Human) ReturnAge() int {
	return h.Age
}

func (h Human) ReturnName() string {
	return h.Name
}

func (a Action) GetAge() int {
	return a.Human.ReturnAge()
}

func (a Action) GetName() string {
	return a.Human.ReturnName()
}

func main() {
	me := Action{
		Human: Human{"Alexander Borovikov", 24},
	}
	fmt.Println(me.GetName(), me.GetAge())
}

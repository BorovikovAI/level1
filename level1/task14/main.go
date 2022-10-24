// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.

package main

import (
	"fmt"
	"reflect"
)

// Package reflect implements run-time reflection,
// allowing a program to manipulate objects with arbitrary types.

// The typical use is to take a value with static type interface{}
// and extract its dynamic type information by calling TypeOf, which returns a Type.

func main() {
	myMap := map[string]interface{}{
		"int":     4,
		"string":  "defrt",
		"bool":    true,
		"channel": make(chan struct{}),
	}

	UsingReflectTypeOf(myMap)
}

func UsingReflectTypeOf(myMap map[string]interface{}) {
	for k, v := range myMap {
		fmt.Println(k, v, reflect.TypeOf(v))
	}
}

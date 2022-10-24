// Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

package main

import (
	"fmt"
)

func main() {
	slice := []string{"cat", "cat", "dog", "cat", "tree"}

	fmt.Println(GetNewSlice(slice))
}

func GetNewSlice(slice []string) []string {
	resultExist := make(map[string]bool)
	result := make([]string, 0)

	for _, v := range slice {
		if !resultExist[v] {
			result = append(result, v)
			resultExist[v] = true
		}
	}

	return result
}

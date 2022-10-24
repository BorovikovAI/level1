// Реализовать пересечение двух неупорядоченных множеств.

package main

import (
	"fmt"
)

func main() {
	slice1 := []string{"df", "a", "df", "edr"}
	slice2 := []string{"gh4", "df", "eee", "a", "aa", "a"}

	fmt.Println(SliceInter(slice1, slice2))
}

func SliceInter(slice1, slice2 []string) []string {
	result := make([]string, 0)
	existMap := make(map[string]bool)
	resultExist := make(map[string]bool)

	for _, v := range slice1 {
		existMap[v] = true
	}

	for _, v := range slice2 {
		if existMap[v] && !resultExist[v] {
			result = append(result, v)
			resultExist[v] = true
		}
	}

	return result
}

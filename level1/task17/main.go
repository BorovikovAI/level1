//Реализовать бинарный поиск встроенными методами языка.

package main

import (
	"errors"
	"fmt"
)

func binarySearch(array []int, item int) (int, error) {
	lowKey := 0
	highKey := len(array) - 1
	i := 0
	for lowKey <= highKey {
		fmt.Println("\nIteration:", i, ":")
		i++
		mid := (lowKey + highKey) / 2
		guess := array[mid]
		fmt.Println("mid:", mid, ", guess:", guess)
		if guess == item {
			return mid, nil
		} else if guess > item {
			highKey = mid - 1
			fmt.Println("new highKey:", highKey)
		} else {
			lowKey = mid + 1
			fmt.Println("new lowKey:", lowKey)
		}
	}

	return 0, errors.New("No number in array")
}

func main() {
	//array must be sorted
	arr := []int{-5, -1, 0, 2, 6, 10, 16, 19, 56, 77}

	result, err := binarySearch(arr, 19)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nPosition of item:", result)
}

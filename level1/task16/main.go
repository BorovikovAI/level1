//Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

// Разделяем массив на левую и правую части с помощью опорной точки,
// проходим массив, чтобы элементы, которые меньше опорной точки оказались слева от нее,
// а которые больше - справа.
// Дальше берем часть массива до опорной точки и вторую часть после опортной точки,
// повторяем на них сортировку.
// Продолжаем до того момента, как сортируемая часть массива будет пустой или состоять из одного элемента.

package main

import (
	"fmt"
)

func main() {
	myArr := []int{12, 1, -55, 0, -9}
	fmt.Println("Original array:", myArr)

	QuickSort(myArr)
	fmt.Println("Result array:", myArr)
}

func Change(a []int, i, j int) {
	fmt.Println("changed")
	c := a[i]
	a[i] = a[j]
	a[j] = c
}

func QuickSort(myArr []int) {
	if len(myArr) <= 1 {
		return
	}

	sortPoint := SliceTheArray(myArr)
	fmt.Println("\nSortPoint:", sortPoint)

	QuickSort(myArr[:sortPoint])
	fmt.Println("1:", myArr)
	QuickSort(myArr[sortPoint:])
	fmt.Println("2:", myArr)
}

func SliceTheArray(myArr []int) int {
	sp := myArr[len(myArr)/2]
	fmt.Println("len/2:", len(myArr)/2, ", sp:", sp)

	left := 0
	right := len(myArr) - 1

	for {
		for ; myArr[left] < sp; left++ {
			fmt.Println(myArr[left], "<", sp)
		}

		for ; myArr[right] > sp; right-- {
			fmt.Println(myArr[right], ">", sp)
		}

		if left >= right {
			fmt.Println("Left:", left, ", Right:", right)
			return right
		}

		Change(myArr, left, right)
	}
}

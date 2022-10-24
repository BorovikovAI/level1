//Разработать программу нахождения расстояния между двумя точками,
//которые представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func CreatePoint(x float64, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func Distance(p1 *Point, p2 *Point) float64 {
	return math.Sqrt(math.Pow((p1.x-p2.x), 2) + math.Pow((p1.y-p2.y), 2))
}

func main() {
	a := CreatePoint(20, 0)
	b := CreatePoint(10, -30)
	fmt.Println(Distance(a, b))
}

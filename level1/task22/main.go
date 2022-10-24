//Разработать программу,
//которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

package main

import (
	"fmt"
	"math/big"
)

func main() {
	num := new(big.Int)
	num2 := new(big.Int)
	res := new(big.Int)
	num.SetString("500000000000000000000000", 10)
	num2.SetString("200000000000000000000000", 10)

	fmt.Println("a+b:", res.Add(num, num2))
	fmt.Println("a-b:", res.Sub(num, num2))
	fmt.Println("a*b:", res.Mul(num, num2))
	fmt.Println("a/b:", res.Div(num, num2))
}

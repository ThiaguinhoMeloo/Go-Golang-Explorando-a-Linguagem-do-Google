package fundamentos

import (
	"fmt"
	"math"
)

func Arithmetics() {
	a := 3
	b := 2

	// bitwise
	fmt.Println("E => ", a&b)   // 11 & 10 = 10
	fmt.Println("Ou => ", a|b)  // 11 | 10 = 11
	fmt.Println("Xor => ", a^b) // 11 ^ 10 = 01

	c := 3.0
	d := 2.0

	// outras operações usando math
	fmt.Println("Maior => ", math.Max(float64(a), float64(b)))
	fmt.Println("Menor => ", math.Min(c, d))
	fmt.Println("Exponenciação => ", math.Pow(c, d))
}

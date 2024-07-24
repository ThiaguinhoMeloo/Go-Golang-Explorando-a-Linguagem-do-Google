package main

import (
	"fmt"
	"inicio/functions"
)

func main() {
	x, y := functions.ToReplace(2, 3)
	fmt.Println(x, y)

	second, first := functions.ToReplace(7, 1)
	fmt.Println(second, first)
}

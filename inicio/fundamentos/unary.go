package fundamentos

import (
	"fmt"
)

func Unary() {
	x := 1
	y := 2

	// apenas postfix
	x++ // x += 1 ou x = x + 1
	fmt.Println(x)

	y-- // y -= 1 ou x = y - 1
	fmt.Println(y)
}

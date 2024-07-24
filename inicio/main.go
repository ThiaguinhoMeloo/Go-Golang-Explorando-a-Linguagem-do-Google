package main

import (
	"fmt"
	"inicio/functions"
)

func main() {
	result := functions.Exec(functions.Multiplication, 3, 4)
	fmt.Println(result)
}

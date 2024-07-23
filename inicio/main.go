package main

import (
	"fmt"
	"inicio/controlstructures"
	"time"
)

// "inicio/conversoes"

func main() {
	fmt.Println(controlstructures.Type(2.3))
	fmt.Println(controlstructures.Type(1))
	fmt.Println(controlstructures.Type("Teste"))
	fmt.Println(controlstructures.Type(func() {}))
	fmt.Println(controlstructures.Type(time.Now()))
}

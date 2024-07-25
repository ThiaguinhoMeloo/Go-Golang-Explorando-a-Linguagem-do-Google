package main

import (
	"fmt"
	"inicio/concorrencia"
)

func main() {
	c := make(chan int, 60)
	go concorrencia.Cousins(cap(c), c)
	for cousin := range c {
		fmt.Printf("%d ", cousin)
	}
	fmt.Println("Fim!")
}

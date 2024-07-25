package main

import (
	"fmt"
	"inicio/concorrencia"
)

func main() {
	c := make(chan int)
	go concorrencia.TwoThreeFourTimes(2, c)
	fmt.Println("A")

	a, b := <-c, <-c //recebendo dados do canal
	fmt.Println("B")
	fmt.Println(a, b)

	fmt.Println(<-c)
}

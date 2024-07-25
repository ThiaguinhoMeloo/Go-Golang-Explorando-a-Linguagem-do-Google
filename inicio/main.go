package main

import (
	"fmt"
	"inicio/concorrencia"
)

func main() {
	c := make(chan int) // canal sem buffer
	go concorrencia.Routine(c)
	fmt.Println(<-c) // operação bloqueante
	fmt.Println("Foi lido")
	fmt.Println(<-c) // deadlock
	fmt.Println("Fim")
}

package concorrencia

import (
	"fmt"
	"time"
)

func Routine(c chan int) {
	time.Sleep(time.Second)
	c <- 1 // operação bloqueante.
	fmt.Println("Só depois que o foi lido")
}

// func main() {
// 	c := make(chan int) // canal sem buffer
// 	go concorrencia.Routine(c)
// 	fmt.Println(<-c) // operação bloqueante
// 	fmt.Println("Foi lido")
// 	fmt.Println(<-c) // deadlock
// 	fmt.Println("Fim")
// }

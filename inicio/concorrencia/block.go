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

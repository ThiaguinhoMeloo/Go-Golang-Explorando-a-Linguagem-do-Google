package main

import (
	"fmt"
	"inicio/concorrencia"
)

func main() {
	c := concorrencia.JoinMult(concorrencia.SpeakMult("João"), concorrencia.SpeakMult("Maria"))
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
	fmt.Println(<-c, <-c)
}

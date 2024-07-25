package main

import (
	"fmt"
	"inicio/concorrencia"
)

func main() {
	t1 := concorrencia.Title("https://www.cod3r.com.br", "https://www.google.com")
	t2 := concorrencia.Title("https://www.amazon.com", "https://www.youtube.com")
	fmt.Println("Primeiros:", <-t1, "|", <-t2)
	fmt.Println("Segundos:", <-t1, "|", <-t2)
}

package concorrencia

import "fmt"

func Channel1() {
	ch := make(chan int, 1)

	ch <- 1 // enviando dados para o canal (escrita)
	<-ch    // recebendo dados do canal (leitura)

	ch <- 2
	fmt.Println(<-ch)

	// o channel é um mecanismo de sincronismo.
	// Goroutine é a forma independente.
}

// func main() {
// 	concorrencia.Channel1()
// }

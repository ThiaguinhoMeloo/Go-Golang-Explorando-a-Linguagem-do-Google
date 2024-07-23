package controlstructures

import (
	"fmt"
	"time"
)

func For() {
	i := 1
	for i <= 10 { // não tem while em Go
		fmt.Println(i)
		i++
	}

	fmt.Println("--------------------------------------------------------------------------------------------")

	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d ", i)
	}

	fmt.Println("\n--------------------------------------------------------------------------------------------")

	fmt.Println("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Printf("Par ")
		} else {
			fmt.Printf("Impar ")
		}
	}

	fmt.Println("\n--------------------------------------------------------------------------------------------")

	stop := time.After(10 * time.Second)
	ticker := time.NewTicker(1 * time.Second)
	for {
		// laço infinito
		select {
		case <-stop:
			fmt.Println("Tempo acabou!!")
			return
		case <-ticker.C:
			fmt.Println("Para sempre...")
		}
	}
}

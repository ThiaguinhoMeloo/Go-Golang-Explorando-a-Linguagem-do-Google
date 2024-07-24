package functions

func Average(numbers ...float64) float64 {
	total := 0.0
	for _, num := range numbers {
		total += num
	}

	return total / float64(len(numbers))
	//len serve para ver o tamanho do array.

	// fmt.Printf("Média: %.2f", functions.Average(7.7, 8.1, 5.9, 9.9))

	// funções variaticas são basicamente uma função que pode aceitar um número variável de argumentos do mesmo tipo.
}

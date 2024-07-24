package functions

import "fmt"

func GetProvedValue(number int) int {
	defer fmt.Println("Fim!")

	if number > 5000 {
		fmt.Println("Valor alto...")
		return number
	} else {
		fmt.Println("Valor baixo...")
		return number
	}

	// fmt.Println(functions.GetProvedValue(6000))
	// fmt.Println(functions.GetProvedValue(3000))

	// Defer em Go é usada para adiar a execução de uma função até que a função que a contém retorne.
}

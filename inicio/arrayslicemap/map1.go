package arrayslicemap

import "fmt"

func Map() {
	// var approved map[int]string
	// mapas devem ser inicializados
	approved := make(map[int]string)

	approved[12345678978] = "Maria"
	approved[98765432100] = "João"
	approved[92398293871] = "Ana"

	fmt.Println(approved)

	for cpf, nome := range approved {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(approved[92398293871])
	delete(approved, 92398293871)
	fmt.Println(approved[92398293871])
}

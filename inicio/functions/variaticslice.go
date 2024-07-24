package functions

import "fmt"

func PrintApproved(approveds ...string) {
	fmt.Println("Lista de aprovados: ")
	for i, approved := range approveds {
		fmt.Printf("%d) %s\n", i+1, approved)
	}

	// approveds := []string{"Maria", "Pedro", "Guilherme", "Allan"} // slice
	// functions.PrintApproved(approveds...)
}

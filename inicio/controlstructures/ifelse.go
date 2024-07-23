package controlstructures

import (
	"fmt"
)

func IfElse(note float64) {
	if note >= 7 {
		fmt.Println("Aprovado com nota: ", note)
	} else {
		fmt.Println("Reprovado com nota: ", note)
	}

	// note1 := 7.3
	// note2 := 5.1
	// controlstructures.IfElse(note1)
	// controlstructures.IfElse(note2)
}

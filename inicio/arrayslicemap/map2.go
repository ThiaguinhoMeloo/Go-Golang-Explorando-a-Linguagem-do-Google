package arrayslicemap

import "fmt"

func Map2() {
	funcsEWage := map[string]float64{
		"José João":      11325.45,
		"Gabriela Silva": 15456.78,
		"Pedro Júnior":   1200.0,
	}

	funcsEWage["Rafael Filho"] = 1350.0

	delete(funcsEWage, "Inexistente")

	for name, wage := range funcsEWage {
		fmt.Println(name, wage)
	}
}

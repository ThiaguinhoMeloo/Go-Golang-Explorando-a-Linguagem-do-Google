package arrayslicemap

import "fmt"

func MapAninhado() {
	funcsByLetter := map[string]map[string]float64{
		"G": {
			"Gabriel Silva": 15456.78,
			"Guga Pereira":  8456.78,
		},
		"J": {
			"José João": 11325.45,
		},
		"P": {
			"Pedro Júnior": 1200.0,
		},
	}

	delete(funcsByLetter, "P")

	for latter, funcs := range funcsByLetter {
		fmt.Printf(latter, funcs)
	}
}

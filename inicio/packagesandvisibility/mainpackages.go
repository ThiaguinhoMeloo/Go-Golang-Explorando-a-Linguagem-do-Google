package packagesandvisibility

import "fmt"

func MainPackages() {
	p1 := Ponto{
		X: 2.0,
		Y: 2.0,
	}

	p2 := Ponto{
		X: 2.0,
		Y: 4.0,
	}

	fmt.Println(catetos(p1, p2))
	fmt.Println(Distancia(p1, p2))
}

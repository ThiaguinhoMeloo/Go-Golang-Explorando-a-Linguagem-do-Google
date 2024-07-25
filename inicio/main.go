package main

import (
	"fmt"
	"inicio/packagesandvisibility"
)

func main() {
	p1 := packagesandvisibility.Ponto{
		X: 2.0,
		Y: 2.0,
	}

	p2 := packagesandvisibility.Ponto{
		X: 2.0,
		Y: 4.0,
	}

	fmt.Println(packagesandvisibility.CalcularCatetos(p1, p2))
	fmt.Println(packagesandvisibility.Distancia(p1, p2))
}

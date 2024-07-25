package packagesandvisibility

import "math"

// Iniciando com letra maiúscula é PÚBLICO (visibilidade dentro e fora do pacote)!
// Iniciando com letra minúscula é PRIVADO (visivel apenas dentro do pacote)!

// Por exemplo...
// Ponto - gerará algo público.
// ponto ou _Ponto - gerará algo privado.

// Ponto representa uma coordenada no plano cartesiano.
type Ponto struct {
	X float64
	Y float64
}

func catetos(a, b Ponto) (cx, cy float64) {
	cx = math.Abs(b.X - a.X)
	cy = math.Abs(b.Y - a.Y)
	return
}

// Distancia é responsável por calcular a distância linear entre dois pontos.
func Distancia(a, b Ponto) float64 {
	cx, cy := catetos(a, b)
	return math.Sqrt(math.Pow(cx, 2) + math.Pow(cy, 2))
}

func CalcularCatetos(a, b Ponto) (float64, float64) {
	return catetos(a, b)
}

// func main() {
// 	p1 := packagesandvisibility.Ponto{
// 		X: 2.0,
// 		Y: 2.0,
// 	}

// 	p2 := packagesandvisibility.Ponto{
// 		X: 2.0,
// 		Y: 4.0,
// 	}

// 	fmt.Println(packagesandvisibility.CalcularCatetos(p1, p2))
// 	fmt.Println(packagesandvisibility.Distancia(p1, p2))
// }

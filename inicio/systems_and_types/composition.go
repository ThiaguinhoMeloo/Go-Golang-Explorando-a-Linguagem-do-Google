package systemsandtypes

import "fmt"

type Sports2 interface {
	TurboOn()
}

type Luxurious interface {
	MakeBalisa()
}

type SportLuxurious interface {
	Sports2
	Luxurious
	// pode adicionar mais métodos
}

type Bmw7 struct{}

func (b Bmw7) TurboOn() {
	fmt.Println("Turbo...")
}

func (b Bmw7) MakeBalisa() {
	fmt.Println("Baliza...")
}

// func main() {
// 	var b systemsandtypes.SportLuxurious = systemsandtypes.Bmw7{}
// 	b.TurboOn()
// 	b.MakeBalisa()
// }

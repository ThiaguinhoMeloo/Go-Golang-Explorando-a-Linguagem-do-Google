package systemsandtypes

import "fmt"

type Printable interface {
	ToString() string
}

type People2 struct {
	Name     string
	LastName string
}

type Product2 struct {
	Name  string
	Price float64
}

// interfaces são implementadas implicitamente.

func (p People2) ToString() string {
	return p.Name + " " + p.LastName
}

func (p Product2) ToString() string {
	return fmt.Sprintf("%s - R$ %2.f", p.Name, p.Price)
}

func Print(x Printable) {
	fmt.Println(x.ToString())
}

// func main() {
// 	var coisa systemsandtypes.Printable = systemsandtypes.People2{
// 		Name:     "Roberto",
// 		LastName: "Silva",
// 	}
// 	fmt.Println(coisa.ToString())
// 	systemsandtypes.Print(coisa)

// 	coisa = systemsandtypes.Product2{
// 		Name:  "Calça Jeans",
// 		Price: 79.90,
// 	}
// 	fmt.Println(coisa.ToString())
// 	systemsandtypes.Print(coisa)

// 	p2 := systemsandtypes.Product2{
// 		Name:  "Calça Jeans",
// 		Price: 179.90,
// 	}
// 	systemsandtypes.Print(p2)

// }

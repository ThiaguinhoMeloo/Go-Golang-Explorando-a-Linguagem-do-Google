package systemsandtypes

type Product struct {
	Name     string
	Price    float64
	Discount float64
}

// Método: função com receiver (receptor)

func (p Product) DiscountedProduct() float64 {
	return p.Price * (1 - p.Discount)
}

// Struct é um agrupamento de dados.
// var product1 systemsandtypes.Product
// product1 = systemsandtypes.Product{
// 	Name:     "Lapis",
// 	Price:    1.79,
// 	Discount: 0.05,
// }
// fmt.Println(product1, product1.DiscountedProduct())

// product2 := systemsandtypes.Product{
// 	"Notebook",
// 	1989.90,
// 	0.10,
// }

// fmt.Println(product2.DiscountedProduct())

package main

import (
	"fmt"
	systemsandtypes "inicio/systems_and_types"
)

func main() {
	var product1 systemsandtypes.Product
	product1 = systemsandtypes.Product{
		Name:     "Lapis",
		Price:    1.79,
		Discount: 0.05,
	}
	fmt.Println(product1, product1.DiscountedProduct())

	product2 := systemsandtypes.Product{
		"Notebook",
		1989.90,
		0.10,
	}

	fmt.Println(product2.DiscountedProduct())
}

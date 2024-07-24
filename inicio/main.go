package main

import (
	"fmt"
	systemsandtypes "inicio/systems_and_types"
)

func main() {
	order := systemsandtypes.Order{
		UserId: 1,
		Items: []systemsandtypes.Item{
			systemsandtypes.Item{
				ProductId: 1,
				Qtde:      2,
				Price:     12.10,
			},
			systemsandtypes.Item{
				ProductId: 2,
				Qtde:      1,
				Price:     23.49,
			},
			systemsandtypes.Item{
				ProductId: 11,
				Qtde:      100,
				Price:     3.13,
			},
		},
	}

	fmt.Printf("Valor total do pedido é R$ %.2f", order.Amount())
}

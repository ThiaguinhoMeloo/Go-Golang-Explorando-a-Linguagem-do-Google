package systemsandtypes

type Item struct {
	ProductId int
	Qtde      int
	Price     float64
}

type Order struct {
	UserId int
	Items  []Item
}

func (p Order) Amount() float64 {
	total := 0.0
	for _, item := range p.Items {
		total += item.Price * float64(item.Qtde)
	}
	return total
}

// order := systemsandtypes.Order{
// 	UserId: 1,
// 	Items: []systemsandtypes.Item{
// 		systemsandtypes.Item{
// 			ProductId: 1,
// 			Qtde:      2,
// 			Price:     12.10,
// 		},
// 		systemsandtypes.Item{
// 			ProductId: 2,
// 			Qtde:      1,
// 			Price:     23.49,
// 		},
// 		systemsandtypes.Item{
// 			ProductId: 11,
// 			Qtde:      100,
// 			Price:     3.13,
// 		},
// 	},
// }

// fmt.Printf("Valor total do pedido é R$ %.2f", order.Amount())

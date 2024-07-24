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

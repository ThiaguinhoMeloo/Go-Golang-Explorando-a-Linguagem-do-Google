package functions

func ToReplace(p1, p2 int) (second int, first int) {
	second = p2
	first = p1
	return second, first // retorno limpo

	// x, y := functions.ToReplace(2, 3)
	// fmt.Println(x, y)

	// second, first := functions.ToReplace(7, 1)
	// fmt.Println(second, first)
}

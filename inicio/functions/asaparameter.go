package functions

func Multiplication(a, b int) int {
	return a * b
}

func Exec(funcs func(int, int) int, p1, p2 int) int {
	return funcs(p1, p2)
}

// result := functions.Exec(functions.Multiplication, 3, 4)
// 	fmt.Println(result)

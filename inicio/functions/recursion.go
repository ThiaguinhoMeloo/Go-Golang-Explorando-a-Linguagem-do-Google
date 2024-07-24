package functions

import "fmt"

func Factorial(n int) (int, error) {
	switch {
	case n < 0:
		return -1, fmt.Errorf("número inválido: %d", n)
	case n == 0:
		return 1, nil
	default:
		previousFactorial, _ := Factorial(n - 1)
		return n * previousFactorial, nil
	}

	// result, _ := functions.Factorial(5)
	// fmt.Println(result)

	// _, err := functions.Factorial(-4)
	// if err != nil {
	// 	fmt.Println(err)
	// }
}

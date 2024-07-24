package functions

import "fmt"

func Closure() func() {
	x := 10
	var funcao = func() {
		fmt.Println(x)
	}

	return funcao

	// x := 20
	// fmt.Println(x)

	// printX := functions.Closure()
	// printX()

	// Closure no Go ele tem como função não esquecer a suas raizes ou seja, ela tem ideia do scopo mais abrangente do que foi definido
	// ela trás com sigo as informações de onde ela foi definida para ela ler de forma correta a variavel.
}

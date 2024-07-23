package fundamentos

import (
	"fmt"
)

func Pointer() {
	i := 1

	var p *int = nil

	p = &i // pegando o endereço da variavel
	*p++   // desreferenciando (pegando o vaor)
	i++

	// Go não tem aritmética de ponteiros
	// p++

	fmt.Println(p, *p, i, &i)
}

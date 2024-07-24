package functions

func Inc1(n int) {
	n++
}

// um ponteiro é representado por um *.

func Inc2(n *int) {
	// * é usado para desferenciar, ou seja, ter acesso ao valor no qual o ponteiro aponta.

	*n++
}

// n := 1

// functions.Inc1(n) // por valor
// fmt.Println(n)

// // & é usado para obter o endereço da variavel.

// functions.Inc2(&n)
// fmt.Println(n)

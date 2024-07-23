package fundamentos

func Shopping(trab1, trab2 bool) (bool, bool, bool) {
	tv50 := trab1 && trab2
	tv32 := trab1 != trab2 // ou exclusivo
	iceCream := trab1 || trab2

	// tv50, tv32, iceCream := fundamentos.Shopping(true, true)
	// fmt.Printf("Tv50: %t, Tv32: %t, Sorvete: %t, Saudável: %t", tv50, tv32, iceCream, !iceCream)
	// fundamentos.Unary()

	return tv50, tv32, iceCream
}

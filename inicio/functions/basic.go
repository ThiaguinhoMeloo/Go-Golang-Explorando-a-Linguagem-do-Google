package functions

import (
	"fmt"
)

func F1() {
	fmt.Println("F1")
}

func F2(p1 string, p2 string) {
	fmt.Printf("F2: %s, %s\n", p1, p2)
}

func F3() string {
	return "F3"
}

func F4(p1, p2 string) string {
	return fmt.Sprintf("F4 : %s, %s", p1, p2)
}

func F5() (string, string) {
	return "Retorno 1", "Retorno 2"
}

// functions.F1()
// functions.F2("Param 1", "Param 2")

// r3, r4 := functions.F3(), functions.F4("Param 1", "Param 2")
// fmt.Println(r3)
// fmt.Println(r4)

// r51, r52 := functions.F5()
// fmt.Println("F5: ", r51, r52)

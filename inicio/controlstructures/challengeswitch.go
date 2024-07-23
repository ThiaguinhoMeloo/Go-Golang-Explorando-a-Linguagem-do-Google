package controlstructures

import "fmt"

func ChallengeSwitch(n float64) string {
	fmt.Println("Metodo Refatorada do IfElseIf:")
	switch true {
	case n >= 9 && n <= 10:
		return "A"
	case n >= 8 && n < 9:
		return "B"
	case n >= 6 && n < 8:
		return "C"
	case n >= 3 && n < 5:
		return "D"
	default:
		return "E"
	}

	// fmt.Println(controlstructures.ChallengeSwitch(9.8))
	// fmt.Println(controlstructures.ChallengeSwitch(6.9))
	// fmt.Println(controlstructures.ChallengeSwitch(2.1))
}

package main

import (
	"fmt"
	systemsandtypes "inicio/systems_and_types"
)

func main() {
	car1 := systemsandtypes.Ferrari2{
		Model:        "F40",
		TurboOn:      false,
		CurrentSpeed: 0,
	}
	car1.TurnOnTurbo()

	var car2 systemsandtypes.Sports = &systemsandtypes.Ferrari2{
		Model:        "F40",
		TurboOn:      false,
		CurrentSpeed: 0,
	}
	car2.TurnOnTurbo()

	fmt.Println(car1, car2)
}

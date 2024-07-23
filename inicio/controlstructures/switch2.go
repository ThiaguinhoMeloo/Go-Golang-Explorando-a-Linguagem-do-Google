package controlstructures

import (
	"fmt"
	"time"
)

func Switch2() {
	t := time.Now()
	switch true { // switch true
	case t.Hour() < 12:
		fmt.Println("Bom dia!")
	case t.Hour() < 18:
		fmt.Println("Boa Tarde!")
	default:
		fmt.Println("Boa noite!")
	}
}

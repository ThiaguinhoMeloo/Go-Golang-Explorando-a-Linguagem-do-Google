package controlstructures

import (
	"math/rand"
	"time"
)

func RandomNumber() int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(10)

	// if i := controlstructures.RandomNumber(); i > 5 { // também suportado no switch
	// 	fmt.Println("Ganhou!!!")
	// } else {
	// 	fmt.Println("Perdeu!")
	// }
}

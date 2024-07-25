package concorrencia

import "fmt"

func RoutineBuffer(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	fmt.Println("Executou!")
	ch <- 5
	ch <- 6
}

// func main() {
// 	ch := make(chan int, 3)
// 	go concorrencia.RoutineBuffer(ch)

// 	time.Sleep(time.Second)
// 	fmt.Println(<-ch)
// }

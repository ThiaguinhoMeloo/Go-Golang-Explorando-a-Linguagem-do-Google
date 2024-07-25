package concorrencia

import "time"

func IsCousin(num int) bool {
	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	return true
}

func Cousins(n int, c chan int) {
	start := 2
	for i := 0; i < n; i++ {
		for cousin := start; ; cousin++ {
			if IsCousin(cousin) {
				c <- cousin
				start = cousin + 1
				time.Sleep(time.Millisecond * 100)
				break
			}
		}
	}
	close(c) // para encerrar a função for, para que assim não gere um Deadlock
}

// func main() {
// 	c := make(chan int, 60)
// 	go concorrencia.Cousins(cap(c), c)
// 	for cousin := range c {
// 		fmt.Printf("%d ", cousin)
// 	}
// 	fmt.Println("Fim!")
// }

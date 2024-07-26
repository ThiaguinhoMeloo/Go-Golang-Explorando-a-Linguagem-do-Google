package concorrencia

import (
	"fmt"
	"time"
)

func SpeakMult(people string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 3; i++ {
			time.Sleep(time.Second)
			c <- fmt.Sprintf("%s falando: %d", people, i)
		}
	}()
	return c
}

func JoinMult(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case s := <-input1:
				c <- s
			case s := <-input2:
				c <- s
			}
		}
	}()
	return c
}

// "github.com/ThiaguinhoMeloo/html"
// func main() {
// 	c := concorrencia.JoinMult(concorrencia.SpeakMult("João"), concorrencia.SpeakMult("Maria"))
// 	fmt.Println(<-c, <-c)
// 	fmt.Println(<-c, <-c)
// 	fmt.Println(<-c, <-c)
// c := html.Title("https://google.com")
// fmt.Println(<-c)
// }

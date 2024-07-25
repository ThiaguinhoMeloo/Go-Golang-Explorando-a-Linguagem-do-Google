package concorrencia

import (
	"fmt"
	"time"
)

func Speak(people, text string, qtde int) {
	for i := 0; i < qtde; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%s: %s (iteração %d)\n", people, text, i+1)
	}
}

// func main() {
// 	// concorrencia.Speak("Maira", "Por que não fala comigo?", 3)
// 	// concorrencia.Speak("João", "Só posso falar depois de você!", 1)

// 	// go concorrencia.Speak("Maria", "Ei...", 500)
// 	// go concorrencia.Speak("João", "Ou...", 500)

// 	// time.Sleep(time.Second * 5)
// 	// fmt.Println("Fim!")

// 	go concorrencia.Speak("Maria", "Entendi!!!", 10)
// 	concorrencia.Speak("João", "Parabéns!", 5)
// }

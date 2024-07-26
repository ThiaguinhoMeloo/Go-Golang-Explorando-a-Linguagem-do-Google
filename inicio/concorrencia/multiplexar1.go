package concorrencia

func ToForward(origin <-chan string, destiny chan string) {
	for {
		destiny <- <-origin
	}
}

// Multiplexar - misturar (mensages) num canal

func Join(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go ToForward(input1, c)
	go ToForward(input2, c)

	return c
}

// func main() {
// 	c := concorrencia.Join(
// 		concorrencia.Title("https://www.cod3r.com.br", "https://www.google.com"),
// 		concorrencia.Title("https://github.com", "https://www.youtube.com"),
// 	)
// 	fmt.Println(<-c, "|", <-c)
// 	fmt.Println(<-c, "|", <-c)
// }

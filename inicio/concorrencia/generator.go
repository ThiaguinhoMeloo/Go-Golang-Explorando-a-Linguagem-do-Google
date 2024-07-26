package concorrencia

import (
	"io"
	"net/http"
	"regexp"
)

// padro~es apresentados na aula - Google I/0 2012 - Go Concurrency Patterns

// <-chan - canal somente-leitura

func Title(urls ...string) <-chan string {
	c := make(chan string)
	for _, url := range urls {
		go func(url string) {
			resp, _ := http.Get(url)
			html, _ := io.ReadAll(resp.Body)
			r, _ := regexp.Compile("<title>(.*?)<\\/title>")
			c <- r.FindStringSubmatch(string(html))[1]
		}(url)
	}

	return c
}

// func main() {
// 	t1 := concorrencia.Title("https://www.cod3r.com.br", "https://www.google.com")
// 	t2 := concorrencia.Title("https://github.com", "https://www.youtube.com")
// 	fmt.Println("Primeiros:", <-t1, "|", <-t2)
// 	fmt.Println("Segundos:", <-t1, "|", <-t2)
// }

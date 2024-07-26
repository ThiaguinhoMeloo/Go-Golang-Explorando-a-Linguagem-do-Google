package concorrencia

import "time"

func TheFastest(url1, url2, url3 string) string {
	c1 := Title(url1)
	c2 := Title(url2)
	c3 := Title(url3)

	// estrutura de controle específica para concorrência
	select {
	case t1 := <-c1:
		return t1
	case t2 := <-c2:
		return t2
	case t3 := <-c3:
		return t3
	case <-time.After(1000 * time.Millisecond):
		return "Todos perderam!"
		// default:
		// 	return "Sem resposta ainda!"
	}
}

// func main() {
// 	champion := concorrencia.TheFastest(
// 		"https://www.cod3r.com.br",
// 		"https://www.google.com",
// 		"https://github.com",
// 	)
// 	fmt.Println(champion)
// }

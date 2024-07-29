package http

import (
	"fmt"
	"net/http"
	"time"
)

func RightTime(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Hora certa: %s<h1>", s)

}

// import (
// 	response "inicio/http"
// 	"log"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	http.HandleFunc("/horaCerta", response.RightTime)
// 	log.Println("Executando...")
// 	log.Fatal(http.ListenAndServe(":3000", nil))
// }

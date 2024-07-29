package http

import (
	"log"
	"net/http"
)

func Static() {
	fs := http.FileServer(http.Dir("http/public"))
	http.Handle("/", fs)

	log.Println("Executando...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}

// package main

// import (
// 	"inicio/http"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	http.Static()
// }

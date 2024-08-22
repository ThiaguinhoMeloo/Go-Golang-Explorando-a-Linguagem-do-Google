package main

import (
	request "inicio/http"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/", request.CurrentTime)
	http.ListenAndServe(":3000", nil)
}

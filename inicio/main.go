package main

import (
	"inicio/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.Static()
}

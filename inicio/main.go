package main

import (
	request "inicio/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	request.Static()
}

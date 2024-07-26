package main

import (
	"fmt"
	"inicio/estrutura"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := estrutura.ConnectDbMyDql()
	if err != nil {
		fmt.Println(err.Error())
	}

	estrutura.SelectMySql(db)
}

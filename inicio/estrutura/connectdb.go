package estrutura

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDbMyDql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:Mamutao@2melo@tcp(localhost:3306)/cursogo")
	if err != nil {
		panic(err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

package estrutura

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type user struct {
	ID   int
	Name string
}

func SelectMySql(db *sql.DB) {
	rows, _ := db.Query("select id, name from usuarios where id > ?", 5)
	defer rows.Close()

	for rows.Next() {
		var u user
		rows.Scan(&u.ID, &u.Name)
		fmt.Println(u)
	}
}

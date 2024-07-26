package estrutura

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Exec(db *sql.DB, sql string) sql.Result {
	result, err := db.Exec(sql)
	if err != nil {
		panic(err)
	}

	return result
}

// func main() {
// 	db, err := sql.Open("mysql", "root:Mamutao@2melo@/")
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer db.Close()

// 	estrutura.Exec(db, "create database if not exists cursogo")
// 	estrutura.Exec(db, "use cursogo")
// 	estrutura.Exec(db, "drop table if exists usuarios")
// 	estrutura.Exec(db, `create table usuarios (
// 		id integer auto_increment,
// 		name varchar(80),
// 		PRIMARY KEY(id)
// 	)`)
// }

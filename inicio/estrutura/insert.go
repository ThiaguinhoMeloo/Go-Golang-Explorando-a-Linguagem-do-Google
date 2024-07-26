package estrutura

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertMySql(db *sql.DB) {
	stmt, err := db.Prepare("insert into usuarios(name) values(?)")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()

	stmt.Exec("Maria")
	stmt.Exec("João")

	res, _ := stmt.Exec("Pedro")

	id, _ := res.LastInsertId()
	fmt.Println(id)

	lines, _ := res.RowsAffected()
	fmt.Println(lines)
}

// func main() {
// 	db, err := estrutura.ConnectDbMyDql()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// 	estrutura.InsertMySql(db)
// }

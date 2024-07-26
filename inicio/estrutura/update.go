package estrutura

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func UpdateMySql(db *sql.DB) {
	stmt, _ := db.Prepare("update usuarios set name = ? where id = ?")

	stmt.Exec("Thiago Melo", 1)
	stmt.Exec("Evelyn Pereira", 2)
}

// func main() {
// 	db, err := estrutura.ConnectDbMyDql()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	estrutura.UpdateMySql(db)
// }

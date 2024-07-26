package estrutura

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func Delete(db *sql.DB) {
	stmt, _ := db.Prepare("delete from usuarios where id = ?")
	stmt.Exec(2000)
}

// func main() {
// 	db, err := estrutura.ConnectDbMyDql()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	estrutura.Delete(db)
// }

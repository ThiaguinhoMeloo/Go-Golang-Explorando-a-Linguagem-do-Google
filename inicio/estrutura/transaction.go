package estrutura

import (
	"database/sql"
)

func Transaction(db *sql.DB) {
	tx, _ := db.Begin()
	stmt, _ := tx.Prepare("insert into usuarios(id, name) values(?,?)")

	stmt.Exec(2000, "Bia")
	stmt.Exec(2001, "Carlos")

	// _, err := stmt.Exec(1, "Thiago") // chave duplicada

	// if err != nil {
	// 	tx.Rollback()
	// 	log.Fatal(err)
	// }

	tx.Commit()
}

// func main() {
// 	db, err := estrutura.ConnectDbMyDql()
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}

// 	estrutura.Transaction(db)
// }

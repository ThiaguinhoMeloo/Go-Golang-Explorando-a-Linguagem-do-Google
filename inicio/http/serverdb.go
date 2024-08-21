package http

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"inicio/estrutura"
	"log"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// Usuário
type User struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
}

// UserHandler analisa o request e delega para a função adequada.
func UserHandler(w http.ResponseWriter, r *http.Request) {
	db, err := estrutura.ConnectDbMyDql()
	if err != nil {
		log.Fatal(err.Error())
	}

	sid := strings.TrimPrefix(r.URL.Path, "/usuarios/")
	id, _ := strconv.Atoi(sid)

	switch {
	case r.Method == "GET" && id > 0:
		userById(w, r, id, db)
	case r.Method == "GET":
		allUsers(w, r, db)
	default:
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Desculpa... :(")
	}
}

func userById(w http.ResponseWriter, r *http.Request, id int, db *sql.DB) {
	var u User
	db.QueryRow("select id, name from usuarios where id = ?", id).Scan(&u.Id, &u.Name)

	json, _ := json.Marshal(u)

	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, string(json))
}

func allUsers(w http.ResponseWriter, r *http.Request, db *sql.DB) {
	rows, _ := db.Query("select * from usuarios")
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		rows.Scan(&user.Id, &user.Name)
		users = append(users, user)
	}

	json, _ := json.Marshal(users)

	w.Header().Set("Content-type", "application/json")
	fmt.Fprint(w, string(json))
}

// func main() {

// 	http.HandleFunc("/usuarios/", resquest.UserHandler)
// 	log.Println("Executando...")
// 	log.Fatal(http.ListenAndServe(":3000", nil))
// }

package http

import (
	"net/http"
	"text/template"
	"time"
)

var tpl *template.Template

type Hora struct {
	Hora string
}

func CurrentTime(w http.ResponseWriter, r *http.Request) {
	// http.ServeFile(w, r, "http/hora/index.html")
	tpl, _ = template.ParseFiles("http/hora/index.html")
	loc, _ := time.LoadLocation("America/Sao_Paulo")
	now := time.Now().UTC().In(loc)
	data := Hora{
		Hora: now.Format("2006-01-02 15:04:05"),
	}

	err := tpl.Execute(w, data)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

// func main() {
// 	http.HandleFunc("/", request.ServeTemplate)
// 	http.ListenAndServe(":3000", nil)
// }

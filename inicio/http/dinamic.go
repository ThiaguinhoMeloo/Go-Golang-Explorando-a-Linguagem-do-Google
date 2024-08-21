package http

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

type PageDate struct {
	HoraCerta string
}

func RightTime(w http.ResponseWriter, r *http.Request) {

	// fs := http.FileServer(http.Dir("http/hora"))
	// http.Handle("/http/", http.StripPrefix("/http/", fs))

	s := time.Now().Format("02/01/2006 03:04:05")
	// fmt.Fprintf(w, "<h1>Hora certa: %s<h1>", s)

	data := PageDate{
		HoraCerta: s,
	}

	// verificar o diretorio de trabalho atual
	cwd, err := os.Getwd()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	log.Printf("Current working directory: %s\n", cwd)

	// Verificar se o arquivo existe
	filePath := filepath.Join("http/hora", "index.html")
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		http.Error(w, "File not found", http.StatusNotFound)
		log.Printf("File not found: %s\n", filePath)
		return
	}

	tmpl := template.Must(template.ParseFiles(filePath))
	if err := tmpl.Execute(w, data); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		log.Printf("Template execution error: %s\n", err)
	}
}

// import (
// 	response "inicio/http"
// 	"log"
// 	"net/http"

// 	_ "github.com/go-sql-driver/mysql"
// )

// func main() {
// 	http.HandleFunc("/horaCerta", response.RightTime)
// 	log.Println("Executando...")
// 	log.Fatal(http.ListenAndServe(":3000", nil))
// }

// OU

// func main() {
// 	fs := http.FileServer(http.Dir("http/hora"))
// 	http.Handle("/http/", http.StripPrefix("/http/", fs))
// 	// http.Handle("/", fs)

// 	http.HandleFunc("/horaCerta", request.RightTime)
// 	log.Println("Executando...")
// 	log.Fatal(http.ListenAndServe(":3000", nil))
// }

package http

import (
	"fmt"
	"net/http"
	"time"
)

func RightTimeV2(w http.ResponseWriter, r *http.Request) {
	s := time.Now().Format("02/01/2006 15:04:05")
	fmt.Fprint(w, s)
}

func ServeTemplate(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "http/hora/index.html")
}

// fs := http.FileServer(http.Dir("http/hora"))
// http.Handle("/http/", http.StripPrefix("/http/", fs))
// http.HandleFunc("/horaCerta", request.RightTimeV2)
// http.HandleFunc("/", request.ServeTemplate)

// log.Println("Executando...")
// log.Fatal(http.ListenAndServe(":3000", nil))

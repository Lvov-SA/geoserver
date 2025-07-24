package handlers

import (
	"html/template"
	"log"
	"net/http"
	"os"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../public/index.html")
	data := struct{ Host string }{Host: os.Getenv("HOST")}
	t.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

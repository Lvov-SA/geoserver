package handlers

import (
	"html/template"
	"log"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../public/index.html")
	t.Execute(w, 1)
	if err != nil {
		log.Println(err)
	}
}

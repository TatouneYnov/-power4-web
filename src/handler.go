package power4

import (
	"html/template"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("page/home.html", "static/style.css")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

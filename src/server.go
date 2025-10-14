package power4

import (
	"net/http"
)

func Server() {
	http.HandleFunc("/", homeHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}

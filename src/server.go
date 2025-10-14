package power4

import (
	"net/http"
)

func Server() {
	http.HandleFunc("/", homeHandler)

	http.HandleFunc("/start-game", gameHandler)
	http.HandleFunc("/move", moveHandler)
	http.HandleFunc("/reset", resetHandler)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
	http.ListenAndServe(":8080", nil)
}

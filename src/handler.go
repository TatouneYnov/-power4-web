package power4

import (
	"html/template"
	"log"
	"net/http"
	"strconv"
)

var game *Game

func homeHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("page/home.html")
	if err != nil {
		log.Fatal(err)
	}
	tmpl.Execute(w, nil)
}

func gameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		r.ParseForm()

		player1Name := r.FormValue("player1")
		player2Name := r.FormValue("player2")
		difficulty := r.FormValue("difficulty")

		if player1Name == "" {
			player1Name = "Joueur 1"
		}
		if player2Name == "" {
			player2Name = "Joueur 2"
		}

		var rows, cols int
		switch difficulty {
		case "difficult":
			rows, cols = 6, 9
		case "expert":
			rows, cols = 7, 8
		default:
			rows, cols = 6, 7
		}

		game = NewGame(rows, cols, player1Name, player2Name)
	} else {
		if game == nil {
			game = NewGame(6, 7, "Joueur 1", "Joueur 2")
		}
	}

	renderGame(w)
}

func moveHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/start-game", http.StatusSeeOther)
		return
	}

	if game == nil {
		http.Redirect(w, r, "/start-game", http.StatusSeeOther)
		return
	}

	r.ParseForm()
	column := r.FormValue("column")

	col, err := strconv.Atoi(column)
	if err == nil {
		game.DropToken(col)
	}

	renderGame(w)
}

func resetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/start-game", http.StatusSeeOther)
		return
	}

	if game != nil {
		game = NewGame(game.Rows, game.Cols, game.Player1Name, game.Player2Name)
	} else {
		game = NewGame(6, 7, "Joueur 1", "Joueur 2")
	}
	http.Redirect(w, r, "/start-game", http.StatusSeeOther)
}

func renderGame(w http.ResponseWriter) {
	funcMap := template.FuncMap{
		"iterate": func(n int) []int {
			result := make([]int, n)
			for i := 0; i < n; i++ {
				result[i] = i
			}
			return result
		},
	}

	tmpl, err := template.New("index.html").Funcs(funcMap).ParseFiles("index.html", "page/template/reset.html")
	if err != nil {
		log.Fatal(err)
	}

	err = tmpl.Execute(w, game)
	if err != nil {
		log.Fatal(err)
	}
}

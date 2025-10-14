package power4

type Player int

const (
	Empty   Player = 0
	Player1 Player = 1
	Player2 Player = 2
)

type Game struct {
	Board         [][]Player
	Rows          int
	Cols          int
	CurrentPlayer Player
	IsGameOver    bool
	Winner        Player
	Player1Name   string
	Player2Name   string
}

func NewGame(rows, cols int, player1Name, player2Name string) *Game {
	board := make([][]Player, rows)
	for i := range board {
		board[i] = make([]Player, cols)
	}
	return &Game{
		Board:         board,
		Rows:          rows,
		Cols:          cols,
		CurrentPlayer: Player1,
		IsGameOver:    false,
		Winner:        Empty,
		Player1Name:   player1Name,
		Player2Name:   player2Name,
	}
}

func (g *Game) DropToken(column int) bool {
	if column < 0 || column >= g.Cols || g.IsGameOver {
		return false
	}

	for row := g.Rows - 1; row >= 0; row-- {
		if g.Board[row][column] == Empty {
			g.Board[row][column] = g.CurrentPlayer

			if g.checkWin(row, column) {
				g.IsGameOver = true
				g.Winner = g.CurrentPlayer
			} else if g.isBoardFull() {
				g.IsGameOver = true
			} else {
				if g.CurrentPlayer == Player1 {
					g.CurrentPlayer = Player2
				} else {
					g.CurrentPlayer = Player1
				}
			}
			return true
		}
	}
	return false
}

func (g *Game) isBoardFull() bool {
	for col := 0; col < g.Cols; col++ {
		if g.Board[0][col] == Empty {
			return false
		}
	}
	return true
}

func (g *Game) checkWin(row, col int) bool {
	player := g.Board[row][col]

	count := 0
	for c := 0; c < g.Cols; c++ {
		if g.Board[row][c] == player {
			count++
			if count == 4 {
				return true
			}
		} else {
			count = 0
		}
	}

	count = 0
	for r := 0; r < g.Rows; r++ {
		if g.Board[r][col] == player {
			count++
			if count == 4 {
				return true
			}
		} else {
			count = 0
		}
	}

	for r := 0; r <= g.Rows-4; r++ {
		for c := 0; c <= g.Cols-4; c++ {
			if g.Board[r][c] == player &&
				g.Board[r+1][c+1] == player &&
				g.Board[r+2][c+2] == player &&
				g.Board[r+3][c+3] == player {
				return true
			}
		}
	}

	for r := 3; r < g.Rows; r++ {
		for c := 0; c <= g.Cols-4; c++ {
			if g.Board[r][c] == player &&
				g.Board[r-1][c+1] == player &&
				g.Board[r-2][c+2] == player &&
				g.Board[r-3][c+3] == player {
				return true
			}
		}
	}
	return false
}

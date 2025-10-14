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

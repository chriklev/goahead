package gogame

import "fmt"

// InvalidMoveError is meant to be returned when an invalid move is attempted
type InvalidMoveError struct{}

func (err InvalidMoveError) Error() string {
	return "An invalid move was attempted."
}

// Game stores the functions and variables for a game instance
type Game struct {
	Board     [][]uint8
	Freedoms  [][]uint8
	BlackTurn bool
	BoardSize int
}

// NewGame sets the values of a Game struct to the initial game conditions
func (g *Game) NewGame(boardSize int) {
	g.BoardSize = boardSize

	g.Board = make([][]uint8, boardSize)
	g.Freedoms = make([][]uint8, boardSize)
	for i := 0; i < boardSize; i++ {
		g.Board[i] = make([]uint8, boardSize)
		g.Freedoms[i] = make([]uint8, boardSize)
		for j := 0; j < boardSize; j++ {
			g.Board[i][j] = 0
			g.Freedoms[i][j] = 0
		}
	}
	g.BlackTurn = true
}

// ValidMove checks wether a given move is legal.
func (g Game) ValidMove(x, y int) bool {
	return g.Board[y][x] == 0
}

// MakeMove checks if a given move is legal. If it is, it registers that move in the Game struct and returns true, else returns false.
func (g *Game) MakeMove(x, y, turn int) bool {
	// Check move move is valid
	if g.Board[y][x] != 0 {
		return false
	}

	dx, dy := 1, 0
	for i := 0; i < 4; i++ {
		fmt.Println(dx, dy)
		dx, dy = dy, -dx
	}

	return nil
}

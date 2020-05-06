package gogame

// Game stores the functions and variables for a game instance
type Game struct {
	Board      [][]int8
	TurnColor  int8
	boardSize  int
	blackKills int
	whiteKills int
	turnNumber int
}

// NewGame sets the values of a Game struct to the initial game conditions
func (g *Game) NewGame(boardSize int) {
	g.boardSize = boardSize

	g.Board = make([][]int8, boardSize)
	for i := 0; i < boardSize; i++ {
		g.Board[i] = make([]int8, boardSize)
	}
	g.TurnColor = 1
	g.blackKills = 0
	g.whiteKills = 0
	g.turnNumber = 0
}

// MakeMove is a function for making moves in a game. To pass, give the coords (-1, -1).
// If the move was succesfull it returns true, if the move is illegal, returns false.
func (g *Game) MakeMove(x, y int) bool {
	// Check if spot is clear
	if g.Board[x][y] != 0 {
		return false
	}
	// Place stone
	g.Board[x][y] = g.TurnColor
	// Remove enemy groups that now have no freedoms
	dx, dy := 1, 0
	for i := 0; i < 4; i++ {
		xSide := x + dx
		ySide := y + dy
		if xSide >= 0 && xSide < g.boardSize && ySide >= 0 && ySide < g.boardSize {
			if g.Board[xSide][ySide] == -g.TurnColor {
				g.removeGroupWithoutFreedom(xSide, ySide)
			}
		}
		dx, dy = dy, -dx
	}

	g.TurnColor *= -1

	return true
}

func (g *Game) removeGroupWithoutFreedom(x, y int) {
	var xVisited []int
	var yVisited []int
	groupColor := g.Board[x][y]

	if !g.groupHasFreedom(x, y, &xVisited, &yVisited, groupColor) {
		if groupColor == 1 {
			g.whiteKills += len(xVisited)
		} else {
			g.blackKills += len(xVisited)
		}
		for i, val := range xVisited {
			g.Board[val][yVisited[i]] = 0
		}
	}
}

func (g Game) groupHasFreedom(x, y int, xVisited, yVisited *[]int, groupColor int8) bool {

	*xVisited = append(*xVisited, x)
	*yVisited = append(*yVisited, y)

	// for each side
	dx, dy := 1, 0
	for i := 0; i < 4; i++ {
		xSide := x + dx
		ySide := y + dy

		// if current coordinates are on the board
		if g.isOnBoard(xSide, ySide) {
			c := g.Board[xSide][ySide]
			// if coords don't have a stone, then we found a freedom
			if c == 0 {
				return true
			}
			// if coords have a stone not in the color of the turn
			if c == groupColor {
				visited := false
				for i := range *xVisited {
					if xSide == (*xVisited)[i] && ySide == (*yVisited)[i] {
						visited = true
						break
					}
				}
				if !visited {
					if g.groupHasFreedom(xSide, ySide, xVisited, yVisited, groupColor) {
						return true
					}
				}
			}
		}
		dx, dy = dy, -dx
	}
	return false
}

func (g Game) isOnBoard(x, y int) bool {
	return x >= 0 && x < g.boardSize && y >= 0 && y < g.boardSize
}

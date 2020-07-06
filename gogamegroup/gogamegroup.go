package gogamegroup

import "log"

const (
	bsize = 9
	black = 1
	white = -1
)

// Game manages the methods and variables for a game instance
type Game struct {
	board      [bsize][bsize]int8
	groupboard [bsize][bsize]*group
	boardhist  [][bsize][bsize]int8
	groups     []*group
	blackKills uint8
	whiteKills uint8
	turncolor  int8
}

type group struct {
	stones   [][2]int
	color    int8
	freedoms [][2]int
}

// NewGame returns a fresh Game structure, initialized for a new game
func NewGame() (g Game) {
	g.blackKills = 0
	g.whiteKills = 0
	g.turncolor = black
	return
}

// Adds a new stone and group to the game, but does not add freedoms or remove freedoms from neighboring groups
func (g *Game) addStone(x, y int) {
	g.board[x][y] = g.turncolor

	var newgroup group
	newgroup.stones = make([][2]int, 1)
	newgroup.stones[0] = [2]int{x, y}
	newgroup.color = g.turncolor

	g.groups = append(g.groups, &newgroup)
	g.groupboard[x][y] = &newgroup
}

// MakeMove registers a move in the game struct. Use x = y = -1 to pass
func (g *Game) MakeMove(x, y int) bool {
	// If move is pass
	if x == -1 && y == -1 {
		g.pass()
		g.boardhist = append(g.boardhist, g.board)
		return true
	}
	// If the move is outside of board bounds
	if x < 0 || x >= bsize || y < 0 || y >= bsize {
		return false
	}
	// If the spot is not empty
	if g.board[x][y] != 0 {
		return false
	}

	g.boardhist = append(g.boardhist, g.board)
	g.addStone(x, y)

	visitedgroups := make(map[*group]bool)

	// For each neighbor
	dx, dy := 1, 0
	for i := 0; i < 4; i++ {
		nx, ny := x+dx, y+dy

		if x >= 0 && x < bsize && y >= 0 && y < bsize {
			switch g.board[nx][ny] {
			case -g.turncolor: // if enemy group, then remove their freedom
				ngroup := g.groupboard[nx][ny]
				if !visitedgroups[ngroup] {
					g.removeFreedom(ngroup, x, y)
					visitedgroups[ngroup] = true
				}
			case g.turncolor: // if allied group, then merge
				ngroup := g.groupboard[nx][ny]
				if !visitedgroups[ngroup] {
					g.mergeGroups(g.groupboard[x][y], ngroup)
					visitedgroups[ngroup] = true
				}
			case 0: // if empty then add as freedom
				g.addFreedom(g.groupboard[x][y], x, y)
			}
		}
		dx, dy = dy, -dx
	}

	return true
}

func (g *Game) mergeGroups(gr1, gr2 *group) {
	gr2.stones = append(gr2.stones, gr1.stones...)

}

func (g *Game) pass() {
	log.Fatalln("Pass function not implemented")
}

func (g *Game) removeFreedom(gr *group, x, y int) {
	if len(gr.freedoms) == 1 {
		g.removeGroup(gr)
		return
	}
	for i, freedom := range gr.freedoms {
		if freedom[0] == x && freedom[1] == y {
			gr.freedoms[i] = gr.freedoms[len(gr.freedoms)-1]
			gr.freedoms = gr.freedoms[:len(gr.freedoms)-1]
			return
		}
	}
	log.Fatalln("The group you tried to remove a freedom from did not have that freedom.")
}

func (g *Game) addFreedom(gr *group, x, y int) {
	gr.freedoms = append(gr.freedoms, [2]int{x, y})
}

func (g *Game) removeGroup(gr *group) {

}

func appendIfDoesNotContainGroup(s *[]*group, gr *group) {
	for _, a := range *s {
		if a == gr {
			return
		}
	}
	*s = append(*s, gr)
}

package gogamegroup

// Game stores the functions and variables for a game instance
type Game struct {
	boards      [][9][9]int8
	groups      []*group
	boardgroups [9][9]*group
	blackKills  uint8
	whiteKills  uint8
}

type group struct {
	stones   [][2]int
	color    int8
	freedoms [][2]int
}

// NewGame returns a fresh Game structure, initialized for a new game
func NewGame() (g Game) {
	g.boards = make([][9][9]int8, 1, 100)
	g.groups = make([]*group, 0)
	g.blackKills = 0
	g.whiteKills = 0
	return
}

func (g *Game) newGroup(x, y int, color int8) {
	var newgroup group
	newgroup.stones = make([][2]int, 1)
	newgroup.stones[0] = [2]int{x, y}
	newgroup.color = color
	// then add freedoms

	g.groups = append(g.groups, &newgroup)
	g.boardgroups[x][y] = &newgroup
}

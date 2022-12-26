package game

type Player interface {

	// move As many speeds as they have
	Move()
}
type Game struct {
	// area of match
	field_length int
	// players in game
	players []Player
	// slice of players as interface
	// types are not similar
}

// create game
func NewGame(fileld_length int) *Game {
	return &Game{
		field_length: fileld_length,
		players:      make([]Player, 0),
		// slice with initial value empty
	}
}

// full that empty slice
// joining player to game
func (g *Game) join(player Player) {
	g.players = append(g.players, player)
	return g
}

func (g *Game) MovePlayers() {
	// move all players by for
	for _, Player := range g.players {
		player.Move()
	}

}

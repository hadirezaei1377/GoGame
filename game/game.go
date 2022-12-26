package game

type Player interface {

	// move As many speeds as they have
	Move()
	// return the player situation
	Position() int
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

func (g *Game) CheckWinner() Player {
	// this returned player is winner of game
	for _, Player := range g.players {
		if player.Position > g.field_length 
		

		return Player
	}
	}
	// if no one was not > len return nil
	return nil


func (g *Game) print() {

}
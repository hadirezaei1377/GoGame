package game

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"
)

type Player interface {

	// move As many speeds as they have
	Move()
	// return the player situation
	Position() int
	Name() string
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


	// game environement with Ready codes
func (g *Game) Print() {
	time.Sleep(time.Millisecond * 250)
	clearTerminal()
	fmt.Println("|" + strings.Repeat("-", g.field_length+6) + "|")
	for _, player := range g.Players {
		pos := player.Position()
		name := player.Name()
		row := "|" + strings.Repeat(" ", pos) + name + strings.Repeat(" ", g.field_length-pos+1) + "|"
		fmt.Println("|" + strings.Repeat("_", g.field_length+6) + "|")

	}


}

func clearTerminal() {
	
	cmd := exec.Command("cmd", "/", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
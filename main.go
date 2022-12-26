package main

import (
	gamepkg "GoGame/game"
	"fmt"
)

func main() {

	// make players
	// call from game package
	dog := gamepkg.NewDog("Axlie")
	cat := gamepkg.NewCat("Binky")
	owl := gamepkg.NewOwl("Hanna")
	// make the game(an object named game which has a series of game)
	game := gamepkg.NewGame(50)
	// change package name to gamepkg
	// join(add) players to the game
	game.Join(Dog)
	game.Join(Cat)
	game.Join(Owl)

	// loop til there is a winner
	// winner is a variable and empty at first
	var winner gamepkg.Player
	for winner == nil {
		/* in loop
		// move the players
		// check is there any winner?
		// print the final statement and print the winner
		*/
		game.MovePlayers()
		winner := game.CheckWinner()

		game.Print()
	}
	fmt.Printf("%s won the game!!!", winner.Name())
}

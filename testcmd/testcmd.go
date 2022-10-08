package main

import (
	"fmt"

	"github.com/mt3hr/nimar"
)

func main() {
	table := nimar.NewTable("tr", "testroom")
	player1 := nimar.NewPlayer("player1", "p1", nil)
	player2 := nimar.NewPlayer("player2", "p2", nil)

	table.SetPlayer1(player1)
	table.SetPlayer2(player2)
	table.GetGameManager().StartGame()

	player1.Rihai()

	for _, tile := range player1.GetHand() {
		fmt.Printf("[%s]", tile.GetName())
	}
	fmt.Printf("\n")
	fmt.Printf("%vしゃんてん\n", table.GetGameManager().GetShantenChecker().CheckCountOfShanten(player1))
}

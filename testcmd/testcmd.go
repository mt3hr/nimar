package main

import (
	"fmt"

	"github.com/mt3hr/nimar"
)

func main() {
	table := nimar.NewTable("tr", "testroom")
	table.SetPlayer1(nimar.NewPlayer("player1", "p1", nil))
	table.SetPlayer2(nimar.NewPlayer("player2", "p2", nil))
	table.GetGameManager().StartGame()

	table.GetPlayer1().Rihai()
	table.GetPlayer2().Rihai()
	fmt.Println("プレイヤー1")
	player1Hand := table.GetPlayer1().GetHand()
	for i := range player1Hand {
		fmt.Printf("[%v]", player1Hand[i].GetName())
	}
	if table.GetPlayer1().GetTsumoriTile() != nil {
		fmt.Printf("[%v]", table.GetPlayer1().GetTsumoriTile().GetName())
	}
	fmt.Printf("\n")
	fmt.Printf("%vしゃんてん\n", table.GetGameManager().GetShantenChecker().CheckCountOfShanten(table.GetPlayer1()).Shanten)
	fmt.Printf("\n")
	fmt.Println("プレイヤー2")
	player2Hand := table.GetPlayer2().GetHand()
	for i := range player2Hand {
		fmt.Printf("[%v]", player2Hand[i].GetName())
	}
	if table.GetPlayer2().GetTsumoriTile() != nil {
		fmt.Printf("[%v]", table.GetPlayer2().GetTsumoriTile().GetName())
	}
	fmt.Printf("\n")
	fmt.Printf("%vしゃんてん\n", table.GetGameManager().GetShantenChecker().CheckCountOfShanten(table.GetPlayer2()).Shanten)

}

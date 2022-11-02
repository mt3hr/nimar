package main

import (
	"fmt"
	"strings"

	"github.com/mt3hr/nimar/mahjong"
)

func main() {
	testShantenAgariWithNaki()
}

func testShantenAgariWithNaki() {
	handStr := "[西1] [1索1][1索2][1索3] [3索1][3索2][3索3] [中1][中2][中3] [7索1][8索1][9索1]"
	openTypes := []mahjong.OpenType{
		mahjong.OPEN_NULL,
		mahjong.OPEN_PON,
		mahjong.OPEN_PON,
		mahjong.OPEN_PON,
		mahjong.OPEN_CHI,
	}
	agariTileStr := "[西3]"
	tsumo := false // falseでロン

	table := mahjong.NewTable("tr", "testroom")
	table.Player1 = mahjong.NewPlayer("player1", "p1")
	table.Player2 = mahjong.NewPlayer("player2", "p2")

	mahjong.DEBUG = true
	go table.GameManager.StartGame()

	for i, tiles := range parseTilesSplited(handStr) {
		openedTiles := &mahjong.OpenedTiles{
			OpenType: &openTypes[i],
			Tiles:    tiles,
		}

		switch i {
		case 0:
			table.Player1.Hand = tiles
		case 1:
			table.Player1.OpenedTile1 = openedTiles
		case 2:
			table.Player1.OpenedTile2 = openedTiles
		case 3:
			table.Player1.OpenedTile3 = openedTiles
		case 4:
			table.Player1.OpenedTile4 = openedTiles
		}
	}
	agariTile := parseTile(agariTileStr)
	if tsumo {
		table.Player1.TsumoriTile = agariTile
	} else {
		table.Player1.RonTile = agariTile
	}
	agarikei := table.GameManager.ShantenChecker.CheckCountOfShanten(table.Player1)
	for _, tile := range table.Player1.Hand {
		if tile == nil {
			continue
		}
		fmt.Printf("[%s]", tile.Name)
	}
	for _, openedTiles := range []*mahjong.OpenedTiles{
		table.Player1.OpenedTile1,
		table.Player1.OpenedTile2,
		table.Player1.OpenedTile3,
		table.Player1.OpenedTile4,
	} {
		for _, tile := range openedTiles.Tiles {
			if tile == nil {
				continue
			}
			fmt.Printf("[%s]", tile.Name)
		}
	}
	if table.Player1.TsumoriTile != nil {
		fmt.Printf("[%s]", table.Player1.TsumoriTile.Name)
	}
	if table.Player1.RonTile != nil {
		fmt.Printf("[%s]", table.Player1.RonTile.Name)
	}
	fmt.Println("")
	fmt.Printf("shanten %+v\n", agarikei.Shanten)
	fmt.Printf("table.Player1.Status = %+v\n", table.Player1.Status)
	fmt.Printf("agarikei.Agarikei.Machi = %+v\n", agarikei.Agarikei.Machi)

	point := table.GameManager.PointCalcrator.CalcratePoint(table.Player1, agarikei, table, table.GameManager.ShantenChecker.GetYakuList())
	fmt.Printf("point = %+v\n", point)
}

func parseTilesSplited(tileNamesStrSpaceSplited string) [][]*mahjong.Tile {
	tiles := [][]*mahjong.Tile{}

	for _, tileNames := range strings.Split(tileNamesStrSpaceSplited, " ") {
		tiles = append(tiles, parseTiles(tileNames))
	}
	return tiles
}

func parseTiles(tileNamesStr string) []*mahjong.Tile {
	if tileNamesStr == "" {
		return nil
	}
	tiles := []*mahjong.Tile{}

	tileNames := strings.Split(tileNamesStr, "][")
	tileNames[0] = strings.ReplaceAll(tileNames[0], "[", "")
	tileNames[len(tileNames)-1] = strings.ReplaceAll(tileNames[len(tileNames)-1], "]", "")
	for _, tilename := range tileNames {
		tiles = append(tiles, parseTile(tilename))
	}
	return tiles
}
func parseTile(tileNameStr string) *mahjong.Tile {
	if tileNameStr == "" {
		return nil
	}
	alltiles := mahjong.NewTable("", "").GameManager.GenerateTiles()

	tileName := tileNameStr

	tileName = strings.ReplaceAll(tileName, "[", "")
	tileName = strings.ReplaceAll(tileName, "]", "")
	for _, tile := range alltiles {
		if tile.Name == tileName {
			return tile
		}
	}
	panic("かずのこ")
}

func testShantenAgari() {
	handStr := "[1索1][1索2][1索3][3索1][3索2][3索3][中1][中2][中3][西1][7索1][8索1][9索1]"
	tsumoriTileStr := "" // "[西3]"

	table := mahjong.NewTable("tr", "testroom")
	table.Player1 = mahjong.NewPlayer("player1", "p1")
	table.Player2 = mahjong.NewPlayer("player2", "p2")

	go table.GameManager.StartGame()

	hand := parseTiles(handStr)
	tsumoriTile := parseTile(tsumoriTileStr)
	table.Player1.Hand = hand
	table.Player1.TsumoriTile = tsumoriTile
	shanten := table.GameManager.ShantenChecker.CheckCountOfShanten(table.Player1)
	for _, tile := range append(hand, tsumoriTile) {
		if tile == nil {
			continue
		}
		fmt.Printf("[%s]", tile.Name)
	}
	fmt.Println("")
	fmt.Printf("shanten = %+v\n", shanten)
}

func testHaipai() {
	table := mahjong.NewTable("tr", "testroom")
	table.Player1 = mahjong.NewPlayer("player1", "p1")
	table.Player2 = mahjong.NewPlayer("player2", "p2")
	table.GameManager.StartGame()

	table.Player1.Rihai()
	table.Player2.Rihai()
	fmt.Println("プレイヤー1")
	player1Hand := table.Player1.Hand
	for i := range player1Hand {
		fmt.Printf("[%v]", player1Hand[i].Name)
	}
	if table.Player1.TsumoriTile != nil {
		fmt.Printf("[%v]", table.Player1.TsumoriTile.Name)
	}
	fmt.Printf("\n")
	fmt.Printf("%vしゃんてん\n", table.GameManager.ShantenChecker.CheckCountOfShanten(table.Player1).Shanten)
	fmt.Printf("\n")
	fmt.Println("プレイヤー2")
	player2Hand := table.Player2.Hand
	for i := range player2Hand {
		fmt.Printf("[%v]", player2Hand[i].Name)
	}
	if table.Player2.TsumoriTile != nil {
		fmt.Printf("[%v]", table.Player2.TsumoriTile.Name)
	}
	fmt.Printf("\n")
	fmt.Printf("%vしゃんてん\n", table.GameManager.ShantenChecker.CheckCountOfShanten(table.Player2).Shanten)

}

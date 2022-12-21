// ˅
package mahjong

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var DEBUG bool = false

// ˄

type GameManager struct {
	// ˅

	// ˄

	Table *Table

	ShantenChecker *ShantenChecker

	PointCalcrator *PointCalcrator

	// ˅
	receivedOperator *Operator

	receiveOperatorWG *sync.WaitGroup
	waitStartWg       *sync.WaitGroup

	finishedGame bool

	// ˄
}

func (g *GameManager) ExecuteOperator(operator *Operator) error {
	// ˅
	defer g.receiveOperatorWG.Done()

	var player *Player
	var opponentPlayer *Player

	if g.Table.ID != operator.RoomID {
		return fmt.Errorf("部屋IDとオペレータの対象部屋IDが一致しません。部屋ID:%s オペレータ対象部屋ID:%s", g.Table.ID, operator.RoomID)
	}
	player = g.Table.GetPlayerByID(operator.PlayerID)
	if player == nil {
		return fmt.Errorf("部屋に命令を実行したプレイヤーがいません。プレイヤーID:%s", operator.PlayerID)
	}
	opponentPlayer = g.Table.GetOpponentByID(operator.PlayerID)
	if opponentPlayer == nil {
		return fmt.Errorf("部屋に対局相手がいません。プレイヤーID:%s", operator.PlayerID)
	}

	g.receivedOperator = g.removeNullForOperator(operator)

	switch *operator.OperatorType {
	case OPERATOR_START_GAME:
		defer g.waitStartWg.Done()
		g.receivedOperator = nil
	}

	return nil
	// ˄
}

func (g *GameManager) StartGame() error {
	// ˅

	if DEBUG {
		g.initializeGame()
	}
	g.waitStartWg.Add(2)
	g.receiveOperatorWG.Add(2)
	g.waitStartWg.Wait()
	g.receiveOperatorWG.Wait()
	if !DEBUG {
		g.initializeGame()
	}

	var err error
	for tsumo := true; !g.finishedGame; {
		g.Table.UpdateView()
		tsumo, err = g.gameLoop(tsumo)
		if err != nil {
			return err
		}
		if g.finishedGame {
			break
		}
		if g.Table.Tsumo.CanPop() && len(g.Table.Tsumo.GetDoraHyoujiHais()) <= 4 {
			g.tradeTurn()
		} else {
			g.ryukyoku()
		}
	}
	g.receiveOperatorWG.Wait()
	g.Table.Player1.GameTableWs.Close()
	g.Table.Player1.MessageWs.Close()
	g.Table.Player1.OperatorWs.Close()
	g.Table.Player2.GameTableWs.Close()
	g.Table.Player2.MessageWs.Close()
	g.Table.Player2.OperatorWs.Close()

	return nil
	// ˄
}

func (g *GameManager) GetShantenChecker() *ShantenChecker {
	// ˅
	return g.ShantenChecker
	// ˄
}

func (g *GameManager) preparateGame() {
	// ˅
	ton := KAZE_TON
	nan := KAZE_NAN
	g.Table.Tsumo.Tiles = g.generateTiles()
	g.shuffleTiles(g.Table.Tsumo.Tiles)
	g.Table.Status.PlayerWithTurn = g.Table.Status.Oya
	g.Table.Status.PlayerWithTurn.Status.Kaze = &ton
	g.Table.Status.PlayerWithNotTurn = g.Table.Status.Ko
	g.Table.Status.PlayerWithNotTurn.Status.Kaze = &nan
	g.distributeTiles()
	g.Table.Tsumo.OpenNextKandora()
	g.applyDora()

	// ˄
}

func (g *GameManager) resetTable() {
	// ˅
	g.Table.Player1.Status = &PlayerStatus{}
	g.Table.Player2.Status = &PlayerStatus{}

	g.Table.Tsumo.Tiles = nil
	g.Table.Player1.TsumoriTile = nil
	g.Table.Player1.RonTile = nil
	g.Table.Player1.Hand = []*Tile{}
	g.Table.Player1.Kawa = []*Tile{}
	g.Table.Player1.OpenedTile1 = &OpenedTiles{}
	g.Table.Player1.OpenedTile2 = &OpenedTiles{}
	g.Table.Player1.OpenedTile3 = &OpenedTiles{}
	g.Table.Player1.OpenedTile4 = &OpenedTiles{}
	g.Table.Player1.OpenedPe = &OpenedTiles{}

	g.Table.Player2.TsumoriTile = nil
	g.Table.Player2.RonTile = nil
	g.Table.Player2.Hand = []*Tile{}
	g.Table.Player2.Kawa = []*Tile{}
	g.Table.Player2.OpenedTile1 = &OpenedTiles{}
	g.Table.Player2.OpenedTile2 = &OpenedTiles{}
	g.Table.Player2.OpenedTile3 = &OpenedTiles{}
	g.Table.Player2.OpenedTile4 = &OpenedTiles{}
	g.Table.Player2.OpenedPe = &OpenedTiles{}
	// ˄
}

func (g *GameManager) initializeGame() {
	// ˅
	g.Table.Player1.Point = 50000
	g.Table.Player2.Point = 50000
	g.resetTable()
	g.determinateOya()
	g.Table.Status.ChichaPlayer = g.Table.Status.Oya
	g.preparateGame()
	// ˄
}

func (g *GameManager) getOyaPlayer() *Player {
	// ˅
	return g.Table.Status.Oya
	// ˄
}

func (g *GameManager) generateTiles() []*Tile {
	// ˅
	tiles := []*Tile{}

	for i := 1; i <= 4; i++ {
		// 萬子
		for j := 0; j < 2; j++ {
			name := ""
			num := 0
			dora := 0
			suit := MANZU
			akadora := false

			if j == 0 {
				num = 1
			} else {
				num = 9
			}

			name = fmt.Sprintf("%d%s%d", num, suit.ToString(), i)
			if i == 1 && j == 5 {
				akadora = true
				name += "赤ドラ"
			}

			tiles = append(tiles, &Tile{
				ID:      num,
				Name:    name,
				Num:     num,
				Dora:    dora,
				Suit:    &suit,
				Akadora: akadora,
			})

		}
		// 索子
		for j := 1; j <= 9; j++ {
			name := ""
			num := j
			dora := 0
			suit := SOZU
			akadora := false

			name = fmt.Sprintf("%d%s%d", num, suit.ToString(), i)
			if i == 1 && j == 5 {
				akadora = true
				name += "赤ドラ"
			}

			tiles = append(tiles, &Tile{
				ID:      j + 10,
				Name:    name,
				Num:     num,
				Dora:    dora,
				Suit:    &suit,
				Akadora: akadora,
			})
		}
		// 筒子
		for j := 0; j < 2; j++ {
			name := ""
			num := 0
			dora := 0
			suit := PINZU
			akadora := false

			if j == 0 {
				num = 1
			} else {
				num = 9
			}

			name = fmt.Sprintf("%d%s%d", num, suit.ToString(), i)
			if i == 1 && j == 5 {
				akadora = true
				name += "赤ドラ"
			}

			tiles = append(tiles, &Tile{
				ID:      num + 20,
				Name:    name,
				Num:     num,
				Dora:    dora,
				Suit:    &suit,
				Akadora: akadora,
			})
		}

		name := ""
		num := 0
		dora := 0
		akadora := false

		{
			suit := TON
			name = fmt.Sprintf("%s%d", suit.ToString(), i)
			num = 31
			tiles = append(tiles, &Tile{
				ID:      num,
				Name:    name,
				Num:     num,
				Dora:    dora,
				Suit:    &suit,
				Akadora: akadora,
			})
		}

		{
			suit := NAN
			name = fmt.Sprintf("%s%d", suit.ToString(), i)
			num = 32
			tiles = append(tiles, &Tile{
				ID:      num,
				Name:    name,
				Num:     num,
				Dora:    dora,
				Suit:    &suit,
				Akadora: akadora,
			})
		}

		{
			suit := SHA
			name = fmt.Sprintf("%s%d", suit.ToString(), i)
			num = 33
			tiles = append(tiles, &Tile{
				ID:      num,
				Name:    name,
				Num:     num,
				Dora:    dora,
				Suit:    &suit,
				Akadora: akadora,
			})
		}

		{
			suit := PE
			name = fmt.Sprintf("%s%d", suit.ToString(), i)
			num = 34
			tiles = append(tiles, &Tile{
				ID:      num,
				Name:    name,
				Num:     num,
				Dora:    dora,
				Suit:    &suit,
				Akadora: akadora,
			})
		}

		{
			suit := HAKU
			name = fmt.Sprintf("%s%d", suit.ToString(), i)
			num = 35
			tiles = append(tiles, &Tile{
				ID:      num,
				Name:    name,
				Num:     num,
				Dora:    dora,
				Suit:    &suit,
				Akadora: akadora,
			})
		}

		{
			suit := HATSU
			name = fmt.Sprintf("%s%d", suit.ToString(), i)
			num = 36
			tiles = append(tiles, &Tile{
				ID:      num,
				Name:    name,
				Num:     num,
				Dora:    dora,
				Suit:    &suit,
				Akadora: akadora,
			})
		}

		{
			suit := CHUN
			name = fmt.Sprintf("%s%d", suit.ToString(), i)
			num = 37
			tiles = append(tiles, &Tile{
				ID:      num,
				Name:    name,
				Num:     num,
				Dora:    dora,
				Suit:    &suit,
				Akadora: akadora,
			})
		}
	}
	return tiles
	// ˄
}

func (g *GameManager) determinateOya() {
	// ˅
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(2)
	if random == 1 {
		g.Table.Status.Oya = g.Table.Player1
		g.Table.Status.Ko = g.Table.Player2
	} else {
		g.Table.Status.Oya = g.Table.Player2
		g.Table.Status.Ko = g.Table.Player1
	}
	// ˄
}

func (g *GameManager) shuffleTiles(tiles []*Tile) {
	// ˅
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(tiles); i++ {
		randomIndex := rand.Intn(len(tiles))
		temp := tiles[i]
		tiles[i] = tiles[randomIndex]
		tiles[randomIndex] = temp
	}
	// ˄
}

func (g *GameManager) distributeTiles() {
	// ˅
	var hand []*Tile
	var tile *Tile
	Tsumo := g.Table.Tsumo
	for i := 0; i < 3; i++ {
		hand = g.Table.Status.Oya.Hand
		for j := 0; j < 4; j++ {
			tile = Tsumo.Pop()
			hand = append(hand, tile)
		}
		g.Table.Status.Oya.Hand = hand
		g.Table.UpdateView()

		hand = g.Table.Status.Ko.Hand
		for j := 0; j < 4; j++ {
			tile = Tsumo.Pop()
			hand = append(hand, tile)
		}
		g.Table.Status.Ko.Hand = hand
		g.Table.UpdateView()
	}

	hand = g.Table.Status.Oya.Hand
	tile = Tsumo.Pop()
	hand = append(hand, tile)
	g.Table.Status.Oya.Hand = hand
	g.Table.UpdateView()

	hand = g.Table.Status.Ko.Hand
	tile = Tsumo.Pop()
	hand = append(hand, tile)
	g.Table.Status.Ko.Hand = hand
	g.Table.UpdateView()
	// ˄
}

func (g *GameManager) appendKyushuKyuhaiOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	kyushukyuhai := OPERATOR_KYUSHUKYUHAI
	if player.Status.KyushuKyuhai {
		operators = append(operators, &Operator{
			RoomID:       g.Table.ID,
			PlayerID:     player.ID,
			OperatorType: &kyushukyuhai,
		})
	}
	return operators
	// ˄
}

func (g *GameManager) appendAnkanOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	if !g.Table.Tsumo.CanPop() && len(g.Table.Tsumo.GetDoraHyoujiHais()) <= 4 {
		return operators
	}
	ankan := OPERATOR_ANKAN
	tileIDs := HandAndAgariTile(player)
	for tileID := range tileIDs {
		if tileIDs[tileID] == 4 {
			ankanTiles := []*Tile{}
			for _, tile := range append(player.Hand, player.TsumoriTile) {
				if tile.ID == tileID {
					ankanTiles = append(ankanTiles, tile)
				}
			}
			operators = append(operators, &Operator{
				RoomID:       g.Table.ID,
				PlayerID:     player.ID,
				OperatorType: &ankan,
				TargetTiles:  ankanTiles,
			})
		}
	}
	return operators
	// ˄
}

func (g *GameManager) appendKakanOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	if !g.Table.Tsumo.CanPop() && len(g.Table.Tsumo.GetDoraHyoujiHais()) <= 4 {
		return operators
	}
	kakan := OPERATOR_KAKAN
	for _, OpenedTiles := range []*OpenedTiles{
		player.OpenedTile1,
		player.OpenedTile2,
		player.OpenedTile3,
		player.OpenedTile4,
	} {
		if OpenedTiles.IsNil() {
			continue
		}
		if *OpenedTiles.OpenType == OPEN_PON {
			for _, tile := range append(player.Hand, player.TsumoriTile) {
				if tile == nil {
					continue
				}
				if tile.ID == OpenedTiles.Tiles[0].ID {
					operators = append(operators, &Operator{
						RoomID:       g.Table.ID,
						PlayerID:     player.ID,
						OperatorType: &kakan,
						TargetTiles:  []*Tile{tile},
					})
				}
			}
		}
	}
	return operators
	// ˄
}

func (g *GameManager) appendPeOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	if !g.Table.Tsumo.CanPop() {
		return operators
	}
	pe := OPERATOR_PE

	for _, tile := range append(player.Hand, player.TsumoriTile) {
		if tile == nil {
			continue
		}
		if tile.ID == 34 {
			operators = append(operators, &Operator{
				RoomID:       g.Table.ID,
				PlayerID:     player.ID,
				OperatorType: &pe,
				TargetTiles:  []*Tile{tile},
			})
		}
	}
	return operators
	// ˄
}

func (g *GameManager) appendTsumoAgariOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	Tsumo := OPERATOR_TSUMO

	agarikei := g.GetShantenChecker().CheckCountOfShanten(player)
	if agarikei.Shanten == -1 {
		if len(g.ShantenChecker.GetYakuList().MatchYakus(player, g.Table, agarikei)) == 0 {
			return operators
		}

		operators = append(operators, &Operator{
			RoomID:       g.Table.ID,
			PlayerID:     player.ID,
			OperatorType: &Tsumo,
			TargetTiles:  []*Tile{player.TsumoriTile},
		})
	}
	return operators
	// ˄
}

func (g *GameManager) appendReachOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	if player.Status.Reach {
		return operators
	}
	reach := OPERATOR_REACH
	canReach := true
	for _, openedTiles := range []OpenedTiles{
		*player.OpenedTile1,
		*player.OpenedTile2,
		*player.OpenedTile3,
		*player.OpenedTile4,
	} {
		if openedTiles.IsNil() {
			continue
		}
		if *openedTiles.OpenType == OPEN_PON ||
			*openedTiles.OpenType == OPEN_CHI ||
			*openedTiles.OpenType == OPEN_DAIMINKAN ||
			*openedTiles.OpenType == OPEN_KAKAN {
			canReach = false
			break
		}
	}
	if !canReach {
		return operators
	}

	handTemp := append([]*Tile{}, player.Hand...)
	TsumoriTileTemp := player.TsumoriTile

	for i, sutehai := range player.Hand {
		if i == len(player.Hand) {
			player.Hand = player.Hand[:i]
		} else {
			player.Hand = append(player.Hand[:i], player.Hand[i+1:]...)
		}
		if g.ShantenChecker.CheckCountOfShanten(player).Shanten == 0 {
			operators = append(operators, &Operator{
				RoomID:       g.Table.ID,
				PlayerID:     player.ID,
				OperatorType: &reach,
				TargetTiles:  []*Tile{sutehai},
			})
		}
		copy(player.Hand, handTemp)
		player.Hand = append([]*Tile{}, handTemp...)
	}

	player.TsumoriTile = nil
	if g.ShantenChecker.CheckCountOfShanten(player).Shanten == 0 {
		operators = append(operators, &Operator{
			RoomID:       g.Table.ID,
			PlayerID:     player.ID,
			OperatorType: &reach,
			TargetTiles:  []*Tile{TsumoriTileTemp},
		})
	}
	player.TsumoriTile = TsumoriTileTemp

	return operators
	// ˄
}

func (g *GameManager) appendDahaiOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	dahai := OPERATOR_DAHAI
	if player.TsumoriTile != nil {
		operators = append(operators, &Operator{
			RoomID:       g.Table.ID,
			PlayerID:     player.ID,
			OperatorType: &dahai,
			TargetTiles:  []*Tile{player.TsumoriTile},
		})
		if player.Status.Reach {
			return operators
		}
	}

	if player.Status.Reach {
		return operators
	}

	for _, tile := range player.Hand {
		operators = append(operators, &Operator{
			RoomID:       g.Table.ID,
			PlayerID:     player.ID,
			OperatorType: &dahai,
			TargetTiles:  []*Tile{tile},
		})
	}
	return operators
	// ˄
}

// ˅

func (g *GameManager) appendRonOperators(player *Player, opponentPlayer *Player, operators []*Operator) []*Operator {
	ron := OPERATOR_RON
	agarikei := g.ShantenChecker.CheckCountOfShanten(opponentPlayer)
	if agarikei.Shanten != 0 {
		return operators
	}

	defer func() { opponentPlayer.RonTile = nil }()
	opponentPlayer.RonTile = player.Kawa[len(player.Kawa)-1]
	if len(g.ShantenChecker.GetYakuList().MatchYakus(opponentPlayer, g.Table, agarikei)) == 0 {
		return operators
	}

	for machihaiID := range agarikei.Agarikei.MachiHai {
		huriten := false
		for _, kawaTile := range opponentPlayer.Kawa {
			if kawaTile.ID == machihaiID {
				huriten = true
			}
		}
		if huriten {
			continue
		}
		if machihaiID == player.Kawa[len(player.Kawa)-1].ID {
			operators = append(operators, &Operator{
				RoomID:       g.Table.ID,
				PlayerID:     opponentPlayer.ID,
				OperatorType: &ron,
				TargetTiles:  []*Tile{player.Kawa[len(player.Kawa)-1]},
			})
		}
	}
	return operators
}
func (g *GameManager) appendPonOperators(player *Player, opponentPlayer *Player, operators []*Operator) []*Operator {
	if opponentPlayer.Status.Reach {
		return operators
	}
	pon := OPERATOR_PON
	haiNum := player.Kawa[len(player.Kawa)-1].ID
	tileIDs := HandAndAgariTile(opponentPlayer)
	for tileID := range tileIDs {
		if tileIDs[tileID] >= 2 && haiNum == tileID {
			targetTiles := []*Tile{player.Kawa[len(player.Kawa)-1]}
			addedCnt := 0
			for _, tile := range opponentPlayer.Hand {
				if tile.ID == haiNum {
					targetTiles = append(targetTiles, tile)
					addedCnt++
				}
				if addedCnt == 2 {
					break
				}
			}
			operators = append(operators, &Operator{
				RoomID:       g.Table.ID,
				PlayerID:     opponentPlayer.ID,
				OperatorType: &pon,
				TargetTiles:  targetTiles,
			})
		}
	}
	return operators
}
func (g *GameManager) appendChiOperators(player *Player, opponentPlayer *Player, operators []*Operator) []*Operator {
	if opponentPlayer.Status.Reach {
		return operators
	}
	chi := OPERATOR_CHI
	haiNum := player.Kawa[len(player.Kawa)-1].ID
	menzenTiles := HandAndAgariTile(opponentPlayer)
	tii := []*TileIDs{}

	for i := 0; i <= 2; i++ {
		for j := 1; j <= 9; j++ {
			if haiNum == i*10+j && j <= 7 {
				if menzenTiles[haiNum+1] >= 1 && menzenTiles[haiNum+2] >= 1 {
					syuntsu := &TileIDs{}
					syuntsu[haiNum+1] = 1
					syuntsu[haiNum+2] = 1
					tii = append(tii, syuntsu)
				}
			}
			if haiNum == i*10+j && j >= 3 && j <= 9 {
				if menzenTiles[haiNum-1] >= 1 && menzenTiles[haiNum-2] >= 1 {
					syuntsu := &TileIDs{}
					syuntsu[haiNum-1] = 1
					syuntsu[haiNum-2] = 1
					tii = append(tii, syuntsu)
				}
			}
			if haiNum == i*10+j && j >= 2 && j <= 8 {
				if menzenTiles[haiNum+1] >= 1 && menzenTiles[haiNum-1] >= 1 {
					syuntsu := &TileIDs{}
					syuntsu[haiNum+1] = 1
					syuntsu[haiNum-1] = 1
					tii = append(tii, syuntsu)
				}
			}
		}
	}
	for _, mentsu := range tii {
		targetTiles := []*Tile{player.Kawa[len(player.Kawa)-1]}
		for tileid, cnt := range mentsu {
			if cnt == 1 {
				for _, tile := range append(opponentPlayer.Hand) {
					if tile.ID == tileid {
						targetTiles = append(targetTiles, tile)
						break
					}
				}
			}
		}
		operators = append(operators, &Operator{
			RoomID:       g.Table.ID,
			PlayerID:     opponentPlayer.ID,
			OperatorType: &chi,
			TargetTiles:  targetTiles,
		})
	}
	return operators
}
func (g *GameManager) appendDaiminkanOperators(player *Player, opponentPlayer *Player, operators []*Operator) []*Operator {
	if !g.Table.Tsumo.CanPop() && len(g.Table.Tsumo.GetDoraHyoujiHais()) <= 4 {
		return operators
	}
	if opponentPlayer.Status.Reach {
		return operators
	}
	daiminkan := OPERATOR_DAIMINKAN
	haiNum := player.Kawa[len(player.Kawa)-1].ID
	for tileid, cnt := range HandAndAgariTile(opponentPlayer) {
		if cnt == 3 && haiNum == tileid {
			targetTiles := []*Tile{player.Kawa[len(player.Kawa)-1]}
			for _, tile := range opponentPlayer.Hand {
				if tile.ID == tileid {
					targetTiles = append(targetTiles, tile)
				}
			}
			operators = append(operators, &Operator{
				RoomID:       g.Table.ID,
				PlayerID:     opponentPlayer.ID,
				OperatorType: &daiminkan,
				TargetTiles:  targetTiles,
			})
		}
	}
	return operators
}

func newGameManager(Table *Table) *GameManager {
	return &GameManager{
		Table:             Table,
		receiveOperatorWG: &sync.WaitGroup{},
		waitStartWg:       &sync.WaitGroup{},
		ShantenChecker:    NewShantenChecker(),
		PointCalcrator:    &PointCalcrator{},
	}
}
func (m *GameManager) generateAgariMessage(player *Player) *Message {
	agarikei := m.ShantenChecker.CheckCountOfShanten(player)
	point := m.PointCalcrator.CalcratePoint(player, agarikei, m.Table, m.Table.GameManager.ShantenChecker.yakuList)
	magari := MessageAgari
	message := &Message{MessageType: &magari}
	openNull := OPEN_NULL
	agari := &Agari{
		ID:                player.ID,
		Name:              player.Name,
		OpenedTile1:       &OpenedTiles{OpenType: &openNull},
		OpenedTile2:       &OpenedTiles{OpenType: &openNull},
		OpenedTile3:       &OpenedTiles{OpenType: &openNull},
		OpenedTile4:       &OpenedTiles{OpenType: &openNull},
		Pe:                &OpenedTiles{OpenType: &openNull},
		DoraHyoujiHais:    m.Table.Tsumo.GetDoraHyoujiHais(),
		UraDoraHyoujiHais: m.Table.Tsumo.GetUraDoraHyoujiHais(),
		Player:            player,
		Oya:               player.ID == m.Table.Status.Oya.ID,
	}
	for _, tile := range player.Hand {
		agari.Hand = append(agari.Hand, tile)
	}
	agari.TsumoriTile = player.TsumoriTile
	agari.RonTile = player.RonTile
	agariOpenedTiles := []*OpenedTiles{
		agari.OpenedTile1,
		agari.OpenedTile2,
		agari.OpenedTile3,
		agari.OpenedTile4,
		agari.Pe,
	}
	for i, OpenedTiles := range []*OpenedTiles{
		player.OpenedTile1,
		player.OpenedTile2,
		player.OpenedTile3,
		player.OpenedTile4,
		player.OpenedPe,
	} {
		if OpenedTiles.IsNil() {
			continue
		}
		(*agariOpenedTiles[i]) = (*OpenedTiles.ToOpenedTiles())
	}
	agari.Point = &Point{}
	agari.Point.Hu = point.Hu
	agari.Point.Han = point.Han
	agari.Point.Point = point.Point
	for _, yaku := range point.MatchYakus {
		agari.Point.MatchYakus = append(agari.Point.MatchYakus, yaku)
		agari.Point.MatchYakusForMessage = append(agari.Point.MatchYakusForMessage, convertYakuForMessage(yaku, player.IsMenzen()))
	}
	message.Agari = agari
	return message
}

func (g *GameManager) printShantenCount(player *Player) {
	shanten := g.Table.GameManager.ShantenChecker.CheckCountOfShanten(player)
	fmt.Printf("向聴数 %+v\n", shanten.Shanten)
	if shanten.Shanten == 0 {
		for machi := range shanten.Agarikei.MachiHai {
			fmt.Printf("machi = %+v\n", machi)
		}
	}
	if shanten.Shanten == -1 {
		fmt.Printf("%s\n", shanten.Agarikei.String())
		point := g.PointCalcrator.CalcratePoint(player, shanten, g.Table, g.ShantenChecker.yakuList)
		for _, yaku := range point.MatchYakus {
			if player.IsMenzen() {
				fmt.Printf("%s %d翻\n", yaku.GetName(), yaku.NumberOfHan())
			} else {
				fmt.Printf("%s %d翻\n", yaku.GetName(), yaku.NumberOfHanWhenNaki())
			}
		}
		fmt.Printf("%+v符%+v翻 %+v点\n", point.Hu, point.Han, point.Point)
	}
}

func (g *GameManager) gameLoop(tsumo bool) (nextTurnCanTsumo bool, err error) {
TOP:
	{
		player, opponentPlayer := g.getPlayers()

		player.Status.Haitei = false
		player.Status.Rinshan = false
		opponentPlayer.Status.Hotei = false

		player.Rihai()
		opponentPlayer.Rihai()
		opponentPlayer.Rihai()
		if tsumo {
			player.TsumoriTile = g.Table.Tsumo.Pop()
		}
		if NewKyushuKyuhai().IsMatch(player, g.Table, nil) {
			player.Status.KyushuKyuhai = true
		} else {
			player.Status.KyushuKyuhai = false
		}

		if NewTenho(0, 0).IsMatch(player, g.Table, nil) {
			player.Status.Tenho = true
		} else {
			player.Status.Tenho = false
		}

		if NewChiho(0, 0).IsMatch(player, g.Table, nil) {
			player.Status.Chiho = true
		} else {
			player.Status.Chiho = false
		}

		//TODO 河底撈魚と海底摸月が重複することが多々見られる
		if g.Table.Tsumo.RemainTilesCount() <= 18 {
			player.Status.Haitei = true
			opponentPlayer.Status.Hotei = true
		}

		playerOperators := []*Operator{}
		playerOperators = g.appendKyushuKyuhaiOperators(player, playerOperators)
		playerOperators = g.appendAnkanOperators(player, playerOperators)
		playerOperators = g.appendKakanOperators(player, playerOperators)
		playerOperators = g.appendPeOperators(player, playerOperators)
		playerOperators = g.appendTsumoAgariOperators(player, playerOperators)
		playerOperators = g.appendReachOperators(player, playerOperators)
		playerOperators = g.appendDahaiOperators(player, playerOperators)
		playerOperators = g.removeNullForOperators(playerOperators)

		// g.printShantenCount(player)

		opponentPlayer.Rihai()
		g.Table.UpdateView()
		b, err := json.Marshal(playerOperators)
		if err != nil {
			panic(err)
		}
		_, err = player.OperatorWs.Write(b)
		if err != nil {
			panic(err)
		}
		g.receiveOperatorWG.Add(1)
		g.receiveOperatorWG.Wait()

		operator := g.receivedOperator
		if operator == nil {
			return false, nil
		}

		switch *(operator.OperatorType) {
		case OPERATOR_KYUSHUKYUHAI:
			kyushukyuhai := MessageKyushuKyuhai
			ok := OPERATOR_OK
			message := &Message{
				MessageType: &kyushukyuhai,
			}
			b, err := json.Marshal(message)
			if err != nil {
				panic(err)
			}
			g.receiveOperatorWG.Add(2)
			player.MessageWs.Write(b)
			opponentPlayer.MessageWs.Write(b)

			operatorForPlayer := &Operator{
				RoomID:       g.Table.ID,
				PlayerID:     player.ID,
				OperatorType: &ok,
			}
			b, err = json.Marshal([]*Operator{operatorForPlayer})
			// player.OperatorWs.Write(b)

			operatorForOpponentPlayer := &Operator{
				RoomID:       g.Table.ID,
				PlayerID:     operator.PlayerID,
				OperatorType: &ok,
			}
			b, err = json.Marshal([]*Operator{operatorForOpponentPlayer})
			// opponentPlayer.OperatorWs.Write(b)

			g.receiveOperatorWG.Wait()
			g.ryukyoku()
			return false, nil
		case OPERATOR_ANKAN:
			//TODO リーチしているときに暗槓しても面子崩れないかの判定をまだしていない
			opponentPlayer.Status.Ippatsu = false
			player.Status.Ippatsu = false
			player.Status.Rinshan = true
			ankan := OPEN_ANKAN
			OpenedTile := &OpenedTiles{
				OpenType: &ankan,
			}

			removeIndexs := []int{}
			for _, targetTile := range operator.TargetTiles {
				if player.TsumoriTile.Name == targetTile.Name {
					OpenedTile.Tiles = append(OpenedTile.Tiles, player.TsumoriTile)
					player.TsumoriTile = nil
					continue
				}

				for i, tile := range player.Hand {
					if tile.Name == targetTile.Name {
						removeIndexs = append(removeIndexs, i)
					}
				}
			}

			for i := 0; i < len(removeIndexs); i++ {
				index := removeIndexs[i] - i
				targetTile := player.Hand[index]
				player.Hand = append(player.Hand[:index], player.Hand[index+1:]...)
				OpenedTile.Tiles = append(OpenedTile.Tiles, targetTile)
			}

			if player.OpenedTile1.IsNil() {
				player.OpenedTile1 = OpenedTile
			} else if player.OpenedTile2.IsNil() {
				player.OpenedTile2 = OpenedTile
			} else if player.OpenedTile3.IsNil() {
				player.OpenedTile3 = OpenedTile
			} else if player.OpenedTile4.IsNil() {
				player.OpenedTile4 = OpenedTile
			} else {
				return false, fmt.Errorf("カンの完了に失敗しました。すでに4つ牌を開いています？")
			}

			player.TsumoriTile = g.Table.Tsumo.PopFromWanpai()
			if !g.Table.Tsumo.OpenNextKandora() {
				g.Table.Status.Sukaikan = true
			}
			tsumo = false

			flush := &Flush{
				Message: "カン",
				Player:  player,
			}
			b, err := json.Marshal(flush)
			if err != nil {
				panic(err)
			}
			g.Table.Player1.FlushWs.Write(b)
			g.Table.Player2.FlushWs.Write(b)

			goto TOP
		case OPERATOR_KAKAN:
			player.Status.Rinshan = true
			opponentPlayer.Status.Ippatsu = false
			player.Status.Ippatsu = false
			kakan := OPEN_KAKAN
			for i, mentsu := range [][]*Tile{
				player.OpenedTile1.Tiles,
				player.OpenedTile2.Tiles,
				player.OpenedTile3.Tiles,
				player.OpenedTile4.Tiles,
			} {
				cnt := 0
				for _, tile := range mentsu {
					if tile.ID == operator.TargetTiles[0].ID {
						cnt++
					}
				}
				if cnt == 3 {
					for j := range player.Hand {
						if player.Hand[j].Name == operator.TargetTiles[0].Name {
							player.Hand = append(player.Hand[:j], player.Hand[j+1])
							break
						}
					}
					mentsu = append(mentsu, operator.TargetTiles[0])
					switch i {
					case 0:
						player.OpenedTile1.Tiles = mentsu
						player.OpenedTile1.OpenType = &kakan
					case 1:
						player.OpenedTile2.Tiles = mentsu
						player.OpenedTile2.OpenType = &kakan
					case 2:
						player.OpenedTile3.Tiles = mentsu
						player.OpenedTile3.OpenType = &kakan
					case 3:
						player.OpenedTile4.Tiles = mentsu
						player.OpenedTile4.OpenType = &kakan
					}
					break
				}
			}
			player.TsumoriTile = g.Table.Tsumo.PopFromWanpai()
			if !g.Table.Tsumo.OpenNextKandora() {
				g.Table.Status.Sukaikan = true
			}
			g.applyDora()
			tsumo = false

			flush := &Flush{
				Message: "カン",
				Player:  player,
			}
			b, err := json.Marshal(flush)
			if err != nil {
				panic(err)
			}
			g.Table.Player1.FlushWs.Write(b)
			g.Table.Player2.FlushWs.Write(b)

			opponentPlayer.Rihai()
			opponentPlayer.Rihai()

			goto TOP
		case OPERATOR_PE:
			pe := OPEN_PE
			player.Status.Rinshan = true
			opponentPlayer.Status.Ippatsu = false
			player.Status.Ippatsu = false
			OpenedTile := player.OpenedPe
			OpenedTile.OpenType = &pe
			player.Hand = append(player.Hand, player.TsumoriTile)
			player.TsumoriTile = nil
			for _, targetTile := range operator.TargetTiles {
				tileIndex := -1
				for i, tile := range player.Hand {
					if tile.Name == targetTile.Name {
						tileIndex = i
						break
					}
				}
				player.Hand = append(player.Hand[:tileIndex], player.Hand[tileIndex+1:]...)
				OpenedTile.Tiles = append(OpenedTile.Tiles, targetTile)
			}
			player.OpenedPe = OpenedTile

			player.TsumoriTile = g.Table.Tsumo.PopFromWanpai()
			tsumo = false

			flush := &Flush{
				Message: "北",
				Player:  player,
			}
			b, err := json.Marshal(flush)
			if err != nil {
				panic(err)
			}
			g.Table.Player1.FlushWs.Write(b)
			g.Table.Player2.FlushWs.Write(b)

			opponentPlayer.Rihai()
			opponentPlayer.Rihai()

			goto TOP
		case OPERATOR_TSUMO:
			// g.printShantenCount(player)
			ok := OPERATOR_OK
			message := g.generateAgariMessage(player)

			flush := &Flush{
				Message: "ツモ",
				Player:  player,
			}
			b, err := json.Marshal(flush)
			if err != nil {
				panic(err)
			}
			g.Table.Player1.FlushWs.Write(b)
			g.Table.Player2.FlushWs.Write(b)

			b, err = json.Marshal(message)
			if err != nil {
				panic(err)
			}
			g.receiveOperatorWG.Add(2)
			player.MessageWs.Write(b)
			opponentPlayer.MessageWs.Write(b)

			operatorForPlayer := &Operator{
				RoomID:       g.Table.ID,
				PlayerID:     player.ID,
				OperatorType: &ok,
			}
			b, err = json.Marshal([]*Operator{operatorForPlayer})
			// player.OperatorWs.Write(b)

			operatorForOpponentPlayer := &Operator{
				RoomID:       g.Table.ID,
				PlayerID:     operator.PlayerID,
				OperatorType: &ok,
			}
			b, err = json.Marshal([]*Operator{operatorForOpponentPlayer})
			// opponentPlayer.OperatorWs.Write(b)

			player.Point += message.Agari.Point.Point + g.Table.Status.ReachTablePoint
			opponentPlayer.Point -= message.Agari.Point.Point
			g.Table.Status.ReachTablePoint = 0

			g.receiveOperatorWG.Wait()

			if (g.Table.Player1.Point >= 30000 || g.Table.Player2.Point >= 30000) && ((*g.Table.Status.Kaze == KAZE_NAN && g.Table.Status.NumberOfKyoku >= 2) || *g.Table.Status.Kaze == KAZE_SHA || *g.Table.Status.Kaze == KAZE_PE) || g.Table.Player1.Point < 0 || g.Table.Player2.Point < 0 {
				g.finishGame()
				return false, nil
			}
			g.nextKyoku(player)
			return true, nil
		case OPERATOR_DAHAI:
			player.Status.Ippatsu = false
			handAndTsumoriTile := append(player.Hand, player.TsumoriTile)
			tileIndex := -1
			for i, tile := range handAndTsumoriTile {
				if tile.Name == operator.TargetTiles[0].Name {
					tileIndex = i
					break
				}
			}

			if tileIndex == -1 {
				return false, fmt.Errorf("手牌にない牌は捨てられません。プレイヤーID:%s 牌Name:%s", operator.PlayerID, operator.TargetTiles[0].Name)
			}

			player.Kawa = append(player.Kawa, handAndTsumoriTile[tileIndex])
			player.Hand = append(handAndTsumoriTile[:tileIndex], handAndTsumoriTile[tileIndex+1:]...)
			player.TsumoriTile = nil
			nextTurnCanTsumo = true

			opponentPlayer.Rihai()
			opponentPlayer.Rihai()
		case OPERATOR_REACH:
			player.Point -= 1000
			g.Table.Status.ReachTablePoint += 1000

			player.Status.Reach = true
			player.Status.Ippatsu = true
			handAndTsumoriTile := append(player.Hand, player.TsumoriTile)
			tileIndex := -1
			for i, tile := range handAndTsumoriTile {
				if tile.Name == operator.TargetTiles[0].Name {
					tileIndex = i
					break
				}
			}

			if tileIndex == -1 {
				return false, fmt.Errorf("手牌にない牌は捨てられません。プレイヤーID:%s 牌Name:%s", operator.PlayerID, operator.TargetTiles[0].Name)
			}

			player.Kawa = append(player.Kawa, handAndTsumoriTile[tileIndex])
			player.Hand = append(handAndTsumoriTile[:tileIndex], handAndTsumoriTile[tileIndex+1:]...)
			player.TsumoriTile = nil
			nextTurnCanTsumo = true

			flush := &Flush{
				Message: "立直",
				Player:  player,
			}
			b, err := json.Marshal(flush)
			if err != nil {
				panic(err)
			}
			g.Table.Player1.FlushWs.Write(b)
			g.Table.Player2.FlushWs.Write(b)

			opponentPlayer.Rihai()
			opponentPlayer.Rihai()

		default:
			return false, fmt.Errorf("変なオペレータが渡されました。オペレータタイプ:%d", operator.OperatorType)
		}

		if NewRenho(0, 0).IsMatch(opponentPlayer, g.Table, nil) {
			opponentPlayer.Status.Renho = true
		} else {
			opponentPlayer.Status.Renho = false
		}

		// 相手のOperator
		opponentOperators := []*Operator{}
		opponentOperators = g.appendRonOperators(player, opponentPlayer, opponentOperators)
		opponentOperators = g.appendPonOperators(player, opponentPlayer, opponentOperators)
		opponentOperators = g.appendChiOperators(player, opponentPlayer, opponentOperators)
		opponentOperators = g.appendDaiminkanOperators(player, opponentPlayer, opponentOperators)
		opponentOperators = g.removeNullForOperators(opponentOperators)
		if len(opponentOperators) != 0 {
			skip := OPERATOR_SKIP
			opponentOperators = append(opponentOperators, &Operator{
				RoomID:       g.Table.ID,
				PlayerID:     opponentPlayer.ID,
				TargetTiles:  []*Tile{},
				OperatorType: &skip,
			})
			b, err = json.Marshal(opponentOperators)
			_, err = opponentPlayer.OperatorWs.Write(b)
			if err != nil {
				panic(err)
			}
			g.Table.UpdateView()
			g.receiveOperatorWG.Add(1)
			g.receiveOperatorWG.Wait()

			operator = g.receivedOperator

			if operator == nil {
				return false, nil
			}

			switch *(operator.OperatorType) {
			case OPERATOR_SKIP:
				break
			case OPERATOR_RON:
				flush := &Flush{
					Message: "ロン",
					Player:  opponentPlayer,
				}
				b, err := json.Marshal(flush)
				if err != nil {
					panic(err)
				}
				g.Table.Player1.FlushWs.Write(b)
				g.Table.Player2.FlushWs.Write(b)

				ok := OPERATOR_OK
				opponentPlayer.RonTile = player.Kawa[len(player.Kawa)-1]
				player.Kawa = player.Kawa[:len(player.Kawa)-1]

				// g.printShantenCount(opponentPlayer)
				message := g.generateAgariMessage(opponentPlayer)

				b, err = json.Marshal(message)
				if err != nil {
					panic(err)
				}
				g.receiveOperatorWG.Add(2)
				player.MessageWs.Write(b)
				opponentPlayer.MessageWs.Write(b)

				operatorForPlayer := &Operator{
					RoomID:       g.Table.ID,
					PlayerID:     player.ID,
					OperatorType: &ok,
				}
				b, err = json.Marshal([]*Operator{operatorForPlayer})
				// player.OperatorWs.Write(b)
				// opponentPlayer.MessageWs.Write(b)

				operatorForOpponentPlayer := &Operator{
					RoomID:       g.Table.ID,
					PlayerID:     operator.PlayerID,
					OperatorType: &ok,
				}
				b, err = json.Marshal([]*Operator{operatorForOpponentPlayer})
				// player.OperatorWs.Write(b)
				// opponentPlayer.MessageWs.Write(b)

				opponentPlayer.Point += message.Agari.Point.Point + g.Table.Status.ReachTablePoint
				player.Point -= message.Agari.Point.Point
				g.Table.Status.ReachTablePoint = 0

				g.receiveOperatorWG.Wait()

				if (g.Table.Player1.Point >= 30000 || g.Table.Player2.Point >= 30000) && ((*g.Table.Status.Kaze == KAZE_NAN && g.Table.Status.NumberOfKyoku >= 2) || *g.Table.Status.Kaze == KAZE_SHA || *g.Table.Status.Kaze == KAZE_PE) || g.Table.Player1.Point < 0 || g.Table.Player2.Point < 0 {
					g.finishGame()
					return false, nil
				}

				g.nextKyoku(opponentPlayer)
				return true, nil
			case OPERATOR_PON:
				opponentPlayer.Status.Ippatsu = false
				pon := OPEN_PON
				OpenedTile := &OpenedTiles{
					OpenType: &pon,
				}

				removeIndexs := []int{}
				for i, targetTile := range operator.TargetTiles {
					if i == 0 {
						if player.Kawa[len(player.Kawa)-1].Name == targetTile.Name {
							player.Kawa = player.Kawa[:len(player.Kawa)-1]
							OpenedTile.Tiles = append(OpenedTile.Tiles, targetTile)
						} else {
							return false, fmt.Errorf("ポンできません。相手の捨てた最後の牌:%s ポンしたい牌:%s", player.Kawa[len(player.Kawa)-1].Name, targetTile.Name)
						}
					} else {
						for j := 0; j < len(opponentPlayer.Hand); j++ {
							tile := opponentPlayer.Hand[j]
							if tile.Name == targetTile.Name {
								removeIndexs = append(removeIndexs, j)
							}
						}
					}
				}

				for i := 0; i < len(removeIndexs); i++ {
					index := removeIndexs[i] - i
					targetTile := opponentPlayer.Hand[index]
					opponentPlayer.Hand = append(opponentPlayer.Hand[:index], opponentPlayer.Hand[index+1:]...)
					OpenedTile.Tiles = append(OpenedTile.Tiles, targetTile)
				}

				if opponentPlayer.OpenedTile1.IsNil() {
					opponentPlayer.OpenedTile1 = OpenedTile
				} else if opponentPlayer.OpenedTile2.IsNil() {
					opponentPlayer.OpenedTile2 = OpenedTile
				} else if opponentPlayer.OpenedTile3.IsNil() {
					opponentPlayer.OpenedTile3 = OpenedTile
				} else if opponentPlayer.OpenedTile4.IsNil() {
					opponentPlayer.OpenedTile4 = OpenedTile
				} else {
					return false, fmt.Errorf("ポンの完了に失敗しました。すでに4つ牌を開いています？")
				}
				nextTurnCanTsumo = false

				flush := &Flush{
					Message: "ポン",
					Player:  opponentPlayer,
				}
				b, err := json.Marshal(flush)
				if err != nil {
					panic(err)
				}
				g.Table.Player1.FlushWs.Write(b)
				g.Table.Player2.FlushWs.Write(b)

				opponentPlayer.Rihai()
				opponentPlayer.Rihai()

			case OPERATOR_CHI:
				opponentPlayer.Status.Ippatsu = false
				chi := OPEN_CHI
				OpenedTile := &OpenedTiles{
					OpenType: &chi,
				}

				removeIndexs := []int{}
				for i, targetTile := range operator.TargetTiles {
					if i == 0 {
						if player.Kawa[len(player.Kawa)-1].Name == targetTile.Name {
							player.Kawa = player.Kawa[:len(player.Kawa)-1]
							OpenedTile.Tiles = append(OpenedTile.Tiles, targetTile)
						} else {
							return false, fmt.Errorf("チーできません。相手の捨てた最後の牌:%s チーしたい牌:%s", player.Kawa[len(player.Kawa)-1].Name, targetTile.Name)
						}
					} else {
						for j := 0; j < len(opponentPlayer.Hand); j++ {
							tile := opponentPlayer.Hand[j]
							if tile.Name == targetTile.Name {
								removeIndexs = append(removeIndexs, j)
							}
						}
					}
				}
				for i := 0; i < len(removeIndexs); i++ {
					index := removeIndexs[i] - i
					targetTile := opponentPlayer.Hand[index]
					opponentPlayer.Hand = append(opponentPlayer.Hand[:index], opponentPlayer.Hand[index+1:]...)
					OpenedTile.Tiles = append(OpenedTile.Tiles, targetTile)
				}

				if opponentPlayer.OpenedTile1.IsNil() {
					opponentPlayer.OpenedTile1 = OpenedTile
				} else if opponentPlayer.OpenedTile2.IsNil() {
					opponentPlayer.OpenedTile2 = OpenedTile
				} else if opponentPlayer.OpenedTile3.IsNil() {
					opponentPlayer.OpenedTile3 = OpenedTile
				} else if opponentPlayer.OpenedTile4.IsNil() {
					opponentPlayer.OpenedTile4 = OpenedTile
				} else {
					return false, fmt.Errorf("チーの完了に失敗しました。すでに4つ牌を開いています？")
				}
				nextTurnCanTsumo = false

				flush := &Flush{
					Message: "チー",
					Player:  opponentPlayer,
				}
				b, err := json.Marshal(flush)
				if err != nil {
					panic(err)
				}
				g.Table.Player1.FlushWs.Write(b)
				g.Table.Player2.FlushWs.Write(b)

				opponentPlayer.Rihai()
				opponentPlayer.Rihai()

			case OPERATOR_DAIMINKAN:
				opponentPlayer.Status.Rinshan = true
				opponentPlayer.Status.Ippatsu = false
				daiminkan := OPEN_DAIMINKAN
				OpenedTile := &OpenedTiles{
					OpenType: &daiminkan,
				}

				for i, targetTile := range operator.TargetTiles {
					if i == 0 {
						if player.Kawa[len(player.Kawa)-1].Name == targetTile.Name {
							player.Kawa = player.Kawa[:len(player.Kawa)-1]
							OpenedTile.Tiles = append(OpenedTile.Tiles, targetTile)
							continue
						} else {
							return false, fmt.Errorf("カンできません。相手の捨てた最後の牌:%s カンしたい牌:%s", player.Kawa[len(player.Kawa)-1].Name, targetTile.Name)
						}
					} else {
						tileIndex := -1
						for j, tile := range opponentPlayer.Hand {
							if tile.Name == targetTile.Name {
								tileIndex = j
								break
							}
						}
						opponentPlayer.Hand = append(opponentPlayer.Hand[:tileIndex], opponentPlayer.Hand[tileIndex+1:]...)
						OpenedTile.Tiles = append(OpenedTile.Tiles, targetTile)
					}
				}
				if opponentPlayer.OpenedTile1.IsNil() {
					opponentPlayer.OpenedTile1 = OpenedTile
				} else if opponentPlayer.OpenedTile2.IsNil() {
					opponentPlayer.OpenedTile2 = OpenedTile
				} else if opponentPlayer.OpenedTile3.IsNil() {
					opponentPlayer.OpenedTile3 = OpenedTile
				} else if opponentPlayer.OpenedTile4.IsNil() {
					opponentPlayer.OpenedTile4 = OpenedTile
				} else {
					return false, fmt.Errorf("カンの完了に失敗しました。すでに4つ牌を開いています？")
				}

				opponentPlayer.TsumoriTile = g.Table.Tsumo.PopFromWanpai()
				if !g.Table.Tsumo.OpenNextKandora() {
					g.Table.Status.Sukaikan = true
				}
				g.applyDora()
				nextTurnCanTsumo = false

				flush := &Flush{
					Message: "カン",
					Player:  opponentPlayer,
				}
				b, err := json.Marshal(flush)
				if err != nil {
					panic(err)
				}
				g.Table.Player1.FlushWs.Write(b)
				g.Table.Player2.FlushWs.Write(b)

				opponentPlayer.Rihai()
				opponentPlayer.Rihai()
			}
		}
	}
	return nextTurnCanTsumo, nil
}

func (g *GameManager) GenerateTiles() []*Tile {
	return g.generateTiles()
}

func (g *GameManager) getPlayers() (*Player, *Player) {
	return g.Table.Status.PlayerWithTurn, g.Table.Status.PlayerWithNotTurn
}

func (g *GameManager) tradeTurn() {
	g.Table.Status.PlayerWithTurn, g.Table.Status.PlayerWithNotTurn = g.Table.Status.PlayerWithNotTurn, g.Table.Status.PlayerWithTurn
}

// こんな関数書きたくなかったけど原因がわからないので書きます
func (g *GameManager) removeNullForOperators(operators []*Operator) []*Operator {
	ops := []*Operator{}
	for _, operator := range operators {
		ops = append(ops, g.removeNullForOperator(operator))
	}
	return ops
}

// こんな関数書きたくなかったけど原因がわからないので書きます
func (g *GameManager) removeNullForOperator(operator *Operator) *Operator {
	op := &Operator{
		RoomID:       operator.RoomID,
		PlayerID:     operator.PlayerID,
		OperatorType: operator.OperatorType,
		TargetTiles:  []*Tile{},
	}
	for _, targetTile := range operator.TargetTiles {
		if targetTile == nil {
			continue
		}
		op.TargetTiles = append(op.TargetTiles, targetTile)
	}
	return op
}

func (g *GameManager) calcNotenBappu(player1 *Player, player2 *Player) (player1Tempai bool, player1Bappu int, player2Tempai bool, player2Bappu int) {
	player1Tempai = g.ShantenChecker.CheckCountOfShanten(player1).Shanten == 0
	player2Tempai = g.ShantenChecker.CheckCountOfShanten(player2).Shanten == 0

	if player1Tempai && player2Tempai || !player1Tempai && !player2Tempai {
		player1Bappu = 0
		player2Bappu = 0
	} else if player1Tempai && !player2Tempai {
		player1Bappu = 3000
		player2Bappu = -3000
	} else if !player1Tempai && player2Tempai {
		player1Bappu = -3000
		player2Bappu = 3000
	}
	return player1Tempai, player1Bappu, player2Tempai, player2Bappu
}

func (g *GameManager) nextKyoku(agariPlayer *Player) {
	if agariPlayer == nil || agariPlayer.ID == g.Table.Status.Oya.ID {
		g.Table.Status.NumberOfHonba++
	} else if g.Table.Status.NumberOfKyoku != 2 {
		g.Table.Status.NumberOfKyoku++
		g.Table.Status.NumberOfHonba = 0
	} else {
		g.Table.Status.NumberOfKyoku = 1
		g.Table.Status.NumberOfHonba = 0
		switch *g.Table.Status.Kaze {
		case KAZE_TON:
			*g.Table.Status.Kaze = KAZE_NAN
		case KAZE_NAN:
			*g.Table.Status.Kaze = KAZE_SHA
		case KAZE_SHA:
			*g.Table.Status.Kaze = KAZE_PE
		}
	}

	if (g.Table.Player1.Point >= 30000 || g.Table.Player2.Point >= 30000) && ((*g.Table.Status.Kaze == KAZE_NAN && g.Table.Status.NumberOfKyoku >= 2) || *g.Table.Status.Kaze == KAZE_SHA || *g.Table.Status.Kaze == KAZE_PE) || g.Table.Player1.Point < 0 || g.Table.Player2.Point < 0 {
		g.finishGame()
		return
	} else {
		// テーブルをリセットしてゲーム再スタート
		g.Table.Status.Oya, g.Table.Status.Ko = g.Table.Status.Ko, g.Table.Status.Oya
	}

	g.resetTable()
	g.preparateGame()
	g.Table.Status.Oya, g.Table.Status.Ko = g.Table.Status.Ko, g.Table.Status.Oya
}

func (g *GameManager) ryukyoku() {
	ok := OPERATOR_OK
	ryukyoku := MessageRyukyoku
	player1Tempai, player1Bappu, player2Tempai, player2Bappu := g.calcNotenBappu(g.Table.Player1, g.Table.Player2)

	message := &Message{
		MessageType: &ryukyoku,
		Ryukyoku: &Ryukyoku{
			Player1Tempai: player1Tempai,
			Player2Tempai: player2Tempai,
			Player1Bappu:  player1Bappu,
			Player2Bappu:  player2Bappu,
		},
	}
	b, err := json.Marshal(message)
	if err != nil {
		panic(err)

	}

	g.Table.Player1.MessageWs.Write(b)
	g.Table.Player2.MessageWs.Write(b)

	operatorForPlayer1 := &Operator{
		RoomID:       g.Table.ID,
		PlayerID:     g.Table.Player1.ID,
		OperatorType: &ok,
	}
	b, err = json.Marshal([]*Operator{operatorForPlayer1})
	// g.Table.Player1.OperatorWs.Write(b)

	operatorForPlayer2 := &Operator{
		RoomID:       g.Table.ID,
		PlayerID:     g.Table.Player2.ID,
		OperatorType: &ok,
	}
	b, err = json.Marshal([]*Operator{operatorForPlayer2})
	// g.Table.Player2.OperatorWs.Write(b)

	g.receiveOperatorWG.Add(2)
	g.receiveOperatorWG.Wait()

	g.Table.Player1.Point += player1Bappu
	g.Table.Player2.Point += player2Bappu

	g.nextKyoku(nil)
}
func (g *GameManager) finishGame() {
	g.finishedGame = true
	var winnerPlayer *Player
	var loserPlayer *Player

	if g.Table.Player1.Point > g.Table.Player2.Point {
		winnerPlayer = g.Table.Player1
		loserPlayer = g.Table.Player2
	} else {
		winnerPlayer = g.Table.Player2
		loserPlayer = g.Table.Player1
	}

	messageMatchResult := MessageMatchResult

	message := &Message{
		MessageType: &messageMatchResult,
		MatchResult: &MatchResult{
			WinnerPlayer: winnerPlayer,
			LoserPlayer:  loserPlayer,
		},
	}
	b, err := json.Marshal(message)
	if err != nil {
		panic(err)
	}
	g.receiveOperatorWG.Add(2)
	g.Table.Player1.MessageWs.Write(b)
	g.Table.Player2.MessageWs.Write(b)
	g.receiveOperatorWG.Wait()
}

func (g *GameManager) applyDora() {
	allTiles := []*Tile{}
	allTiles = append(allTiles, g.Table.Tsumo.Tiles...)
	allTiles = append(allTiles, g.Table.Player1.Hand...)
	allTiles = append(allTiles, g.Table.Player1.TsumoriTile)
	allTiles = append(allTiles, g.Table.Player1.RonTile)
	allTiles = append(allTiles, g.Table.Player1.Kawa...)
	allTiles = append(allTiles, g.Table.Player1.OpenedTile1.Tiles...)
	allTiles = append(allTiles, g.Table.Player1.OpenedTile2.Tiles...)
	allTiles = append(allTiles, g.Table.Player1.OpenedTile3.Tiles...)
	allTiles = append(allTiles, g.Table.Player1.OpenedTile4.Tiles...)
	allTiles = append(allTiles, g.Table.Player1.OpenedPe.Tiles...)
	allTiles = append(allTiles, g.Table.Player2.Hand...)
	allTiles = append(allTiles, g.Table.Player2.TsumoriTile)
	allTiles = append(allTiles, g.Table.Player2.RonTile)
	allTiles = append(allTiles, g.Table.Player2.Kawa...)
	allTiles = append(allTiles, g.Table.Player2.OpenedTile1.Tiles...)
	allTiles = append(allTiles, g.Table.Player2.OpenedTile2.Tiles...)
	allTiles = append(allTiles, g.Table.Player2.OpenedTile3.Tiles...)
	allTiles = append(allTiles, g.Table.Player2.OpenedTile4.Tiles...)
	allTiles = append(allTiles, g.Table.Player2.OpenedPe.Tiles...)
	allTilesTemp := []*Tile{}
	for i := range allTiles {
		if i < len(allTiles) && allTiles[i] != nil {
			allTilesTemp = append(allTilesTemp, allTiles[i])
		}
	}
	allTiles = allTilesTemp

	for i := range allTiles {
		if allTiles[i] == nil {
			continue
		}
		allTiles[i].Dora = 0
		allTiles[i].UraDora = 0
	}

	for _, doraHyoujiHai := range g.Table.Tsumo.GetDoraHyoujiHais() {
		if doraHyoujiHai.ID == 11 ||
			doraHyoujiHai.ID == 21 {
			for i, tile := range allTiles {
				if doraHyoujiHai.ID+8 == tile.ID {
					allTiles[i].Dora++
				}
			}
		}
		if doraHyoujiHai.ID == 19 ||
			doraHyoujiHai.ID == 29 {
			for i, tile := range allTiles {
				if doraHyoujiHai.ID-8 == tile.ID {
					allTiles[i].Dora++
				}
			}
		}

		if doraHyoujiHai.ID > 0 && doraHyoujiHai.ID < 8 ||
			doraHyoujiHai.ID == 31 ||
			doraHyoujiHai.ID == 32 ||
			doraHyoujiHai.ID == 33 ||
			doraHyoujiHai.ID == 35 ||
			doraHyoujiHai.ID == 36 {
			for i, tile := range allTiles {
				if doraHyoujiHai.ID+1 == tile.ID {
					allTiles[i].Dora++
				}
			}
		}

		if doraHyoujiHai.ID == 34 {
			for i, tile := range allTiles {
				if doraHyoujiHai.ID-3 == tile.ID {
					allTiles[i].Dora++
				}
			}
		}

		if doraHyoujiHai.ID == 37 {
			for i, tile := range allTiles {
				if doraHyoujiHai.ID-2 == tile.ID {
					allTiles[i].Dora++
				}
			}
		}
	}

	for _, uraDoraHyoujiHai := range g.Table.Tsumo.GetUraDoraHyoujiHais() {
		if uraDoraHyoujiHai.ID == 1 ||
			uraDoraHyoujiHai.ID == 11 ||
			uraDoraHyoujiHai.ID == 21 {
			for i, tile := range allTiles {
				if uraDoraHyoujiHai.ID+8 == tile.ID {
					allTiles[i].UraDora++
				}
			}
		}
		if uraDoraHyoujiHai.ID == 9 ||
			uraDoraHyoujiHai.ID == 19 ||
			uraDoraHyoujiHai.ID == 29 {
			for i, tile := range allTiles {
				if uraDoraHyoujiHai.ID-8 == tile.ID {
					allTiles[i].UraDora++
				}
			}
		}

		if uraDoraHyoujiHai.ID > 10 && uraDoraHyoujiHai.ID < 18 ||
			uraDoraHyoujiHai.ID == 31 ||
			uraDoraHyoujiHai.ID == 32 ||
			uraDoraHyoujiHai.ID == 33 ||
			uraDoraHyoujiHai.ID == 35 ||
			uraDoraHyoujiHai.ID == 36 {
			for i, tile := range allTiles {
				if uraDoraHyoujiHai.ID+1 == tile.ID {
					allTiles[i].UraDora++
				}
			}
		}

		if uraDoraHyoujiHai.ID == 34 {
			for i, tile := range allTiles {
				if uraDoraHyoujiHai.ID-3 == tile.ID {
					allTiles[i].UraDora++
				}
			}
		}

		if uraDoraHyoujiHai.ID == 37 {
			for i, tile := range allTiles {
				if uraDoraHyoujiHai.ID-2 == tile.ID {
					allTiles[i].UraDora++
				}
			}
		}
	}
}

// TODO 槍槓
// TODO 北ロン
// TODO 流し満貫
// TODO 三暗刻
// TODO 四暗刻
// TODO 四暗刻単騎

// ˄

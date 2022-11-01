// ˅
package mahjong

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ˄

type GameManager struct {
	// ˅

	// ˄

	dealerPlayer *Player

	notDealerPlayer *Player

	ShantenChecker *ShantenChecker

	PointCalcrator *PointCalcrator

	// ˅

	Table *Table `json:"-"`

	receivedOperator *Operator

	receiveOperatorWG *sync.WaitGroup
	okWG              *sync.WaitGroup
	waitStartWg       *sync.WaitGroup

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
	g.waitStartWg.Add(2)
	g.receiveOperatorWG.Add(2)
	g.waitStartWg.Wait()
	g.receiveOperatorWG.Wait()

	g.initializeGame()

	var err error
	for tsumo := true; true; {
		tsumo, err = g.gameLoop(tsumo)
		if err != nil {
			return err
		}
		g.tradeTurn()
	}
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
	// ˄
}

func (g *GameManager) resetGame() {
	// ˅
	g.Table.Player1.TsumoriTile = nil
	g.Table.Player1.Hand = []*Tile{}
	g.Table.Player1.Kawa = []*Tile{}
	g.Table.Player1.OpenedTile1 = &OpenedTiles{}
	g.Table.Player1.OpenedTile2 = &OpenedTiles{}
	g.Table.Player1.OpenedTile3 = &OpenedTiles{}
	g.Table.Player1.OpenedTile4 = &OpenedTiles{}

	g.Table.Player2.TsumoriTile = nil
	g.Table.Player2.Hand = []*Tile{}
	g.Table.Player2.Kawa = []*Tile{}
	g.Table.Player2.OpenedTile1 = &OpenedTiles{}
	g.Table.Player2.OpenedTile2 = &OpenedTiles{}
	g.Table.Player2.OpenedTile3 = &OpenedTiles{}
	g.Table.Player2.OpenedTile4 = &OpenedTiles{}
	// ˄
}

func (g *GameManager) initializeGame() {
	// ˅
	ton := KAZE_TON
	nan := KAZE_NAN
	g.resetGame()
	g.determinateDealer()
	g.Table.Tsumo.Tiles = g.generateTiles()
	g.shuffleTiles(g.Table.Tsumo.Tiles)
	g.Table.Status.ChichaPlayer = g.dealerPlayer
	g.Table.Status.PlayerWithTurn = g.dealerPlayer
	g.dealerPlayer.Status.Kaze = &ton
	g.Table.Status.PlayerWithNotTurn = g.notDealerPlayer
	g.notDealerPlayer.Status.Kaze = &nan
	g.distributeTiles()
	//TODO
	// ˄
}

func (g *GameManager) getDealerPlayer() *Player {
	// ˅
	return g.dealerPlayer
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
			dora := false
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
			dora := false
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
			dora := false
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
		dora := false
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

func (g *GameManager) determinateDealer() {
	// ˅
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(2)
	if random == 1 {
		g.dealerPlayer = g.Table.Player1
		g.notDealerPlayer = g.Table.Player2
	} else {
		g.dealerPlayer = g.Table.Player2
		g.notDealerPlayer = g.Table.Player1
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
		hand = g.dealerPlayer.Hand
		for j := 0; j < 4; j++ {
			tile = Tsumo.Pop()
			hand = append(hand, tile)
		}
		g.dealerPlayer.Hand = hand

		hand = g.notDealerPlayer.Hand
		for j := 0; j < 4; j++ {
			tile = Tsumo.Pop()
			hand = append(hand, tile)
		}
		g.notDealerPlayer.Hand = hand
	}

	hand = g.dealerPlayer.Hand
	tile = Tsumo.Pop()
	hand = append(hand, tile)
	g.dealerPlayer.Hand = hand

	hand = g.notDealerPlayer.Hand
	tile = Tsumo.Pop()
	hand = append(hand, tile)
	g.notDealerPlayer.Hand = hand
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
	ankan := OPERATOR_ANKAN
	tileIDs := handAndTsumoriTile(player)
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
	if g.GetShantenChecker().CheckCountOfShanten(player).Shanten == -1 {
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
	reach := OPERATOR_REACH
	canReach := true
	for _, OpenedTiles := range []OpenedTiles{
		*player.OpenedTile1,
		*player.OpenedTile2,
		*player.OpenedTile3,
		*player.OpenedTile4,
	} {
		if OpenedTiles.IsNil() {
			continue
		}
		if *OpenedTiles.OpenType == OPEN_PON ||
			*OpenedTiles.OpenType == OPEN_CHI ||
			*OpenedTiles.OpenType == OPEN_DAIMINKAN ||
			*OpenedTiles.OpenType == OPEN_KAKAN {
			canReach = false
			break
		}
	}
	if !canReach {
		return operators
	}

	var playerTemp Player
	playerTemp = *player
	handTemp := []*Tile{}
	playerTemp.Hand = handTemp
	for _, tile := range playerTemp.Hand {
		handTemp = append(handTemp, tile)
	}

	for i, sutehai := range playerTemp.Hand {
		playerTemp.Hand = append(playerTemp.Hand[:i], playerTemp.Hand[i+1:]...)
		if g.ShantenChecker.CheckCountOfShanten(&playerTemp).Shanten == 0 {
			operators = append(operators, &Operator{
				RoomID:       g.Table.ID,
				PlayerID:     player.ID,
				OperatorType: &reach,
				TargetTiles:  []*Tile{sutehai},
			})
		}
		playerTemp.Hand = handTemp
	}

	TsumoriTileTemp := playerTemp.TsumoriTile
	playerTemp.TsumoriTile = nil
	if g.ShantenChecker.CheckCountOfShanten(&playerTemp).Shanten == 0 {
		operators = append(operators, &Operator{
			RoomID:       g.Table.ID,
			PlayerID:     player.ID,
			OperatorType: &reach,
			TargetTiles:  []*Tile{TsumoriTileTemp},
		})
	}
	playerTemp.TsumoriTile = TsumoriTileTemp

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
	for machihaiID := range agarikei.Machihai {
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
	pon := OPERATOR_PON
	haiNum := player.Kawa[len(player.Kawa)-1].ID
	tileIDs := handAndTsumoriTile(opponentPlayer)
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
	chi := OPERATOR_CHI
	haiNum := player.Kawa[len(player.Kawa)-1].ID
	menzenTiles := handAndTsumoriTile(opponentPlayer)
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
	daiminkan := OPERATOR_DAIMINKAN
	haiNum := player.Kawa[len(player.Kawa)-1].ID
	for tileid, cnt := range handAndTsumoriTile(opponentPlayer) {
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
		okWG:              &sync.WaitGroup{},
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
	agari := &Agari{
		ID:   player.ID,
		Name: player.Name,
	}
	for _, tile := range player.Hand {
		agari.Hand = append(agari.Hand, tile)
	}
	if player.TsumoriTile != nil {
		agari.TsumoriTile = player.TsumoriTile
	}
	if player.RonTile != nil {
		agari.RonTile = player.RonTile
	}
	agariOpenedTiles := []*OpenedTiles{
		agari.OpenedTile1,
		agari.OpenedTile1,
		agari.OpenedTile1,
		agari.OpenedTile1,
		agari.Pe,
	}
	for i, OpenedTiles := range []*OpenedTiles{
		player.OpenedTile1,
		player.OpenedTile2,
		player.OpenedTile3,
		player.OpenedTile4,
		player.OpenedPe,
	} {
		(*agariOpenedTiles[i]) = (*OpenedTiles.ToOpenedTiles())
	}
	agari.Point = &Point{}
	agari.Point.Hu = point.Hu
	agari.Point.Han = point.Han
	agari.Point.Point = point.Point
	for _, yaku := range point.MatchYakus {
		agari.Point.MatchYakus = append(agari.Point.MatchYakus, yaku)
	}
	message.Agari = agari
	return message
}

func (g *GameManager) gameLoop(tsumo bool) (nextTurnCanTsumo bool, err error) {
TOP:
	{
		player, opponentPlayer := g.getPlayers()

		player.Rihai()
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

		if g.Table.Tsumo.RemainTilesCount() <= 18 {
			player.Status.Haitei = true
			opponentPlayer.Status.Hotei = true
		} else {
			player.Status.Haitei = false
			opponentPlayer.Status.Hotei = false
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

		for _, op := range playerOperators {
			for _, tile := range op.TargetTiles {
				fmt.Printf("tile = %+v\n", tile)
			}
		}

		func() {
			shanten := g.Table.GameManager.ShantenChecker.CheckCountOfShanten(player)
			fmt.Printf("向聴数 %+v\n", shanten.Shanten)
			if shanten.Shanten == -1 {
				fmt.Printf("shanten.Agarikei = %+v\n", shanten.Agarikei)
				point := g.PointCalcrator.CalcratePoint(player, shanten, g.Table, g.ShantenChecker.yakuList)
				for _, yaku := range point.MatchYakus {
					fmt.Println(yaku.GetName())
				}
				fmt.Printf("%+v符%+v翻 %+v点\n", point.Hu, point.Han, point.Point)
			}
		}()

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
			//TODO
		case OPERATOR_ANKAN:
			ankan := OPEN_ANKAN
			OpenedTile := &OpenedTiles{
				OpenType: &ankan,
			}

			for _, targetTile := range operator.TargetTiles {
				if player.TsumoriTile.Name == targetTile.Name {
					OpenedTile.Tiles = append(OpenedTile.Tiles, player.TsumoriTile)
					player.TsumoriTile = nil
					continue
				}

				tileIndex := 0
				for i, tile := range player.Hand {
					if tile.Name == targetTile.Name {
						tileIndex = i
						break
					}
				}
				hand := player.Hand
				hand = append(hand[:tileIndex], hand[tileIndex+1:]...)
				player.Hand = hand
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
			goto TOP
		case OPERATOR_KAKAN:
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
					for i := range player.Hand {
						if player.Hand[i].Name == operator.TargetTiles[0].Name {
							player.Hand = append(player.Hand[:i], player.Hand[i+1])
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
			tsumo = false
			goto TOP
		case OPERATOR_PE:
			pe := OPEN_PE
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
			goto TOP
		case OPERATOR_TSUMO:
			//TODO
			message := g.generateAgariMessage(player)
			_ = message

			g.okWG.Add(2)

			//TODO

			/*
				(*player.GetNimaRMessageStreamServer()).Send(message)
				(*opponentPlayer.GetNimaRMessageStreamServer()).Send(message)
				(*player.GetNimaROperatorsStreamServer()).Send(&Operators{
					Operators: []*Operator{
						&Operator{
							RoomID:       g.Table.GetID(),
							PlayerID:     player.GetID(),
							OperatorType: OPERATOR_OK,
						},
					},
				})
				(*opponentPlayer.GetNimaROperatorsStreamServer()).Send(&Operators{
					Operators: []*Operator{
						&Operator{
							RoomID:       g.Table.GetID(),
							PlayerID:     opponentPlayer.GetID(),
							OperatorType: OPERATOR_OK,
						},
					},
				})
			*/
			g.okWG.Wait()
			//TODO 次の局にすすむ

		case OPERATOR_DAHAI:
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
		case OPERATOR_REACH:
			//TODO
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
		default:
			return false, fmt.Errorf("変なオペレータが渡されました。オペレータタイプ:%d", operator.OperatorType)
		}

		//TODO  相手のOperator
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
				opponentPlayer.RonTile = opponentPlayer.Kawa[len(opponentPlayer.Kawa)-1]
				opponentPlayer.Kawa = opponentPlayer.Kawa[:len(opponentPlayer.Kawa)-1]

				message := g.generateAgariMessage(opponentPlayer)
				_ = message

				g.okWG.Add(2)

				//TODO
				/*
					(*player.GetNimaRMessageStreamServer()).Send(message)
					(*opponentPlayer.GetNimaRMessageStreamServer()).Send(message)
					(*player.GetNimaROperatorsStreamServer()).Send(&Operators{
						Operators: []*Operator{
							&Operator{
								RoomID:       g.Table.GetID(),
								PlayerID:     player.GetID(),
								OperatorType: OPERATOR_OK,
							},
						},
					})
					(*opponentPlayer.GetNimaROperatorsStreamServer()).Send(&Operators{
						Operators: []*Operator{
							&Operator{
								RoomID:       g.Table.GetID(),
								PlayerID:     opponentPlayer.GetID(),
								OperatorType: OPERATOR_OK,
							},
						},
					})
				*/
				g.okWG.Wait()
				//TODO 次の局にすすむ
			case OPERATOR_PON:
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
						for i := 0; i < len(opponentPlayer.Hand); i++ {
							tile := opponentPlayer.Hand[i]
							if tile.Name == targetTile.Name {
								removeIndexs = append(removeIndexs, i)
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
			case OPERATOR_CHI:
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
							continue
						} else {
							return false, fmt.Errorf("チーできません。相手の捨てた最後の牌:%s チーしたい牌:%s", player.Kawa[len(player.Kawa)-1].Name, targetTile.Name)
						}
					} else {
						for i := 0; i < len(opponentPlayer.Hand); i++ {
							tile := opponentPlayer.Hand[i]
							if tile.Name == targetTile.Name {
								removeIndexs = append(removeIndexs, i)
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
				} else if opponentPlayer.OpenedTile3.IsNil() {
					opponentPlayer.OpenedTile2 = OpenedTile
				} else if opponentPlayer.OpenedTile3.IsNil() {
					opponentPlayer.OpenedTile3 = OpenedTile
				} else if opponentPlayer.OpenedTile4.IsNil() {
					opponentPlayer.OpenedTile4 = OpenedTile
				} else {
					return false, fmt.Errorf("チーの完了に失敗しました。すでに4つ牌を開いています？")
				}
				nextTurnCanTsumo = false
			case OPERATOR_DAIMINKAN:
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
						for i, tile := range opponentPlayer.Hand {
							if tile.Name == targetTile.Name {
								tileIndex = i
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
				nextTurnCanTsumo = false
			}
		}
	}
	//TODO
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

// ˄

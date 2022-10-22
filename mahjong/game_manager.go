// ˅
package mahjong

import (
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

	table *Table

	shantenChecker *ShantenChecker

	pointCalcrator *PointCalcrator

	// ˅

	receivedOperator *Operator

	receiveOperatorWG *sync.WaitGroup
	okWG              *sync.WaitGroup

	// ˄
}

func (g *GameManager) ExecuteOperator(operator *Operator) error {
	// ˅
	defer g.receiveOperatorWG.Done()
	var player *Player
	var opponentPlayer *Player

	if g.table.id != operator.GetRoomID() {
		return fmt.Errorf("部屋IDとオペレータの対象部屋IDが一致しません。部屋ID:%s オペレータ対象部屋ID:%s", g.table.id, operator.GetRoomID())
	}
	player = g.table.GetPlayerByID(operator.GetPlayerID())
	if player == nil {
		return fmt.Errorf("部屋に命令を実行したプレイヤーがいません。プレイヤーID:%s", operator.GetPlayerID())
	}
	opponentPlayer = g.table.GetOpponentByID(operator.GetPlayerID())
	if opponentPlayer == nil {
		return fmt.Errorf("部屋に対局相手がいません。プレイヤーID:%s", operator.GetPlayerID())
	}

	g.receivedOperator = operator
	return nil
	// ˄
}

func (g *GameManager) StartGame() error {
	// ˅
	g.preparateGame()
	g.initializeGame()

	player := g.table.tableStatus.PlayerWithTurn
	opponentPlayer := g.table.tableStatus.PlayerWithNotTurn

	player.SetTsumoriTile(g.table.GetTsumo().Pop())
	if g.shantenChecker.GetYakuList()["九種九牌"].IsMatch(player, g.table, nil) {
		player.status.KyushuKyuhai = true
	} else {
		player.status.KyushuKyuhai = false
	}

	if g.shantenChecker.GetYakuList()["天和"].IsMatch(player, g.table, nil) {
		player.status.Tenho = true
	} else {
		player.status.Tenho = false
	}

	if g.shantenChecker.GetYakuList()["地和"].IsMatch(player, g.table, nil) {
		player.status.Chiho = true
	} else {
		player.status.Chiho = false
	}

	if g.table.tsumo.RemainTilesCount() <= 18 {
		player.status.Haitei = true
		opponentPlayer.status.Hotei = true
	} else {
		player.status.Haitei = false
		opponentPlayer.status.Hotei = false
	}

	playerOperators := []*Operator{}
	playerOperators = g.appendKyushuKyuhaiOperators(player, playerOperators)
	playerOperators = g.appendAnkanOperators(player, playerOperators)
	playerOperators = g.appendKakanOperators(player, playerOperators)
	playerOperators = g.appendPeOperators(player, playerOperators)
	playerOperators = g.appendTsumoAgariOperators(player, playerOperators)
	playerOperators = g.appendReachOperators(player, playerOperators)
	playerOperators = g.appendDahaiOperators(player, playerOperators)

	//TODO
	/*
		g.receiveOperatorWG.Add(1)
		if player.GetNimaROperatorsStreamServer() != nil {
			(*player.GetNimaROperatorsStreamServer()).Send(&Operators{
				Operators: playerOperators,
			})
		}
	*/
	g.table.UpdateView()
	g.receiveOperatorWG.Wait()

	operator := g.receivedOperator
	if operator == nil {
		return nil
	}
	switch *(operator.GetOperatorType()) {
	case OPERATOR_OK:
		g.okWG.Done()
	case OPERATOR_DRAW:
		if player.GetTsumoriTile() != nil {
			return fmt.Errorf("すでに引いているのに更にひこうとしています。プレイヤーID:%s", operator.GetPlayerID())
		}

		tsumo := g.table.GetTsumo()
		if !tsumo.CanPop() {
			return fmt.Errorf("ツモが18枚を下回ったので引けません。プレイヤーID:%s", operator.GetPlayerID())
		}
		player.SetTsumoriTile(tsumo.Pop())
	case OPERATOR_KYUSHUKYUHAI:
		//TODO
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
						RoomID:       g.table.GetID(),
						PlayerID:     player.GetID(),
						OperatorType: OPERATOR_OK,
					},
				},
			})
			(*opponentPlayer.GetNimaROperatorsStreamServer()).Send(&Operators{
				Operators: []*Operator{
					&Operator{
						RoomID:       g.table.GetID(),
						PlayerID:     opponentPlayer.GetID(),
						OperatorType: OPERATOR_OK,
					},
				},
			})
		*/
		g.okWG.Wait()
		//TODO 次の局にすすむ

	case OPERATOR_RON:
		player.SetRonTile(opponentPlayer.GetKawa()[len(opponentPlayer.GetKawa())-1])
		opponentPlayer.SetKawa(opponentPlayer.GetKawa()[:len(opponentPlayer.GetKawa())-1])

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
						RoomID:       g.table.GetID(),
						PlayerID:     player.GetID(),
						OperatorType: OPERATOR_OK,
					},
				},
			})
			(*opponentPlayer.GetNimaROperatorsStreamServer()).Send(&Operators{
				Operators: []*Operator{
					&Operator{
						RoomID:       g.table.GetID(),
						PlayerID:     opponentPlayer.GetID(),
						OperatorType: OPERATOR_OK,
					},
				},
			})
		*/
		g.okWG.Wait()
		//TODO 次の局にすすむ

	case OPERATOR_KAKAN:
		//TODO
	case OPERATOR_DAHAI:
		hand := player.GetHand()
		kawa := player.GetKawa()
		tileIndex := -1
		for i, tile := range hand {
			if tile.GetID() == operator.GetTargetTiles()[0].id {
				tileIndex = i
				break
			}
		}

		if tileIndex == -1 {
			return fmt.Errorf("手牌にない牌は捨てられません。プレイヤーID:%s 牌ID:%s", operator.GetPlayerID(), operator.GetTargetTiles()[0].id)
		}

		tile := hand[tileIndex]
		hand = append(hand[:tileIndex], hand[tileIndex+1:]...)
		hand = append(hand, player.GetTsumoriTile())
		kawa = append(kawa, tile)
		player.SetHand(hand)
		player.SetKawa(kawa)
		player.SetTsumoriTile(nil)

	case OPERATOR_START_GAME:
		//TODO
	case OPERATOR_SKIP:
		//TODO
	case OPERATOR_PON:
		pon := OPEN_PON
		openedTile := &OpenedTiles{
			openType: &pon,
		}

		for i, targetTile := range operator.GetTargetTiles() {
			if i == 0 {
				if opponentPlayer.GetKawa()[len(opponentPlayer.GetKawa())-1].GetName() == targetTile.GetName() {
					opponentPlayer.SetKawa(opponentPlayer.GetKawa()[:len(opponentPlayer.GetKawa())-1])
					openedTile.tiles = append(openedTile.tiles, targetTile)
					continue
				} else {
					return fmt.Errorf("ポンできません。相手の捨てた最後の牌:%s ポンしたい牌:%s", opponentPlayer.GetKawa()[len(opponentPlayer.GetKawa())-1].GetName(), targetTile.GetName())
				}
			} else {
				tileIndex := 0
				for i, tile := range player.GetHand() {
					if tile.GetName() == targetTile.GetName() {
						tileIndex = i
						break
					}
				}
				hand := player.GetHand()
				hand = append(hand[:tileIndex], hand[tileIndex+1:]...)
				player.SetHand(hand)
				openedTile.tiles = append(openedTile.tiles, targetTile)
			}
		}
		if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile1 = openedTile
		} else if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile2 = openedTile
		} else if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile3 = openedTile
		} else if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile4 = openedTile
		} else {
			return fmt.Errorf("ポンの完了に失敗しました。すでに4つ牌を開いています？")
		}
	case OPERATOR_CHI:
		chi := OPEN_CHI
		openedTile := &OpenedTiles{
			openType: &chi,
		}

		for i, targetTile := range operator.GetTargetTiles() {
			if i == 0 {
				if opponentPlayer.GetKawa()[len(opponentPlayer.GetKawa())-1].GetName() == targetTile.GetName() {
					opponentPlayer.SetKawa(opponentPlayer.GetKawa()[:len(opponentPlayer.GetKawa())-1])
					openedTile.tiles = append(openedTile.tiles, targetTile)
					continue
				} else {
					return fmt.Errorf("チーできません。相手の捨てた最後の牌:%s チーしたい牌:%s", opponentPlayer.GetKawa()[len(opponentPlayer.GetKawa())-1].GetName(), targetTile.GetName())
				}
			} else {
				tileIndex := 0
				for i, tile := range player.GetHand() {
					if tile.GetName() == targetTile.GetName() {
						tileIndex = i
						break
					}
				}
				hand := player.GetHand()
				hand = append(hand[:tileIndex], hand[tileIndex+1:]...)
				player.SetHand(hand)
				openedTile.tiles = append(openedTile.tiles, targetTile)
			}
		}
		if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile1 = openedTile
		} else if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile2 = openedTile
		} else if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile3 = openedTile
		} else if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile4 = openedTile
		} else {
			return fmt.Errorf("チーの完了に失敗しました。すでに4つ牌を開いています？")
		}
	case OPERATOR_DAIMINKAN:
		daiminkan := OPEN_DAIMINKAN
		openedTile := &OpenedTiles{
			openType: &daiminkan,
		}

		for i, targetTile := range operator.GetTargetTiles() {
			if i == 0 {
				if opponentPlayer.GetKawa()[len(opponentPlayer.GetKawa())-1].GetName() == targetTile.GetName() {
					opponentPlayer.SetKawa(opponentPlayer.GetKawa()[:len(opponentPlayer.GetKawa())-1])
					openedTile.tiles = append(openedTile.tiles, targetTile)
					continue
				} else {
					return fmt.Errorf("カンできません。相手の捨てた最後の牌:%s カンしたい牌:%s", opponentPlayer.GetKawa()[len(opponentPlayer.GetKawa())-1].GetName(), targetTile.GetName())
				}
			} else {
				tileIndex := 0
				for i, tile := range player.GetHand() {
					if tile.GetName() == targetTile.GetName() {
						tileIndex = i
						break
					}
				}
				hand := player.GetHand()
				hand = append(hand[:tileIndex], hand[tileIndex+1:]...)
				player.SetHand(hand)
				openedTile.tiles = append(openedTile.tiles, targetTile)
			}
		}
		if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile1 = openedTile
		} else if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile2 = openedTile
		} else if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile3 = openedTile
		} else if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile4 = openedTile
		} else {
			return fmt.Errorf("カンの完了に失敗しました。すでに4つ牌を開いています？")
		}

		player.SetTsumoriTile(g.table.tsumo.PopFromWanpai())
		if !g.table.GetTsumo().OpenNextKandora() {
			g.table.GetStatus().Sukaikan = true
		}

	case OPERATOR_ANKAN:
		ankan := OPEN_ANKAN
		openedTile := &OpenedTiles{
			openType: &ankan,
		}

		for _, targetTile := range operator.GetTargetTiles() {
			if player.GetTsumoriTile().GetName() == targetTile.GetName() {
				openedTile.tiles = append(openedTile.tiles, player.GetTsumoriTile())
				player.SetTsumoriTile(nil)
				continue
			}

			tileIndex := 0
			for i, tile := range player.GetHand() {
				if tile.GetName() == targetTile.GetName() {
					tileIndex = i
					break
				}
			}
			hand := player.GetHand()
			hand = append(hand[:tileIndex], hand[tileIndex+1:]...)
			player.SetHand(hand)
			openedTile.tiles = append(openedTile.tiles, targetTile)
		}
		if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile1 = openedTile
		} else if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile2 = openedTile
		} else if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile3 = openedTile
		} else if *player.openedTile1.openType == OPEN_NULL {
			player.openedTile4 = openedTile
		} else {
			return fmt.Errorf("カンの完了に失敗しました。すでに4つ牌を開いています？")
		}

		player.SetTsumoriTile(g.table.tsumo.PopFromWanpai())
		if !g.table.GetTsumo().OpenNextKandora() {
			g.table.GetStatus().Sukaikan = true
		}
	case OPERATOR_PE:
		pe := OPEN_PE
		openedTile := player.openedPe
		openedTile.openType = &pe
		for _, targetTile := range operator.GetTargetTiles() {
			if player.GetTsumoriTile().GetName() == targetTile.GetName() {
				openedTile.tiles = append(openedTile.tiles, player.GetTsumoriTile())
				player.SetTsumoriTile(nil)
				continue
			}

			tileIndex := 0
			for i, tile := range player.GetHand() {
				if tile.GetName() == targetTile.GetName() {
					tileIndex = i
					break
				}
			}
			hand := player.GetHand()
			hand = append(hand[:tileIndex], hand[tileIndex+1:]...)
			player.SetHand(hand)
			openedTile.tiles = append(openedTile.tiles, targetTile)
		}
		player.openedPe = openedTile

		player.SetTsumoriTile(g.table.tsumo.PopFromWanpai())
	default:
		return fmt.Errorf("変なオペレータが渡されました。オペレータタイプ:%d", operator.GetOperatorType())
	}
	g.table.UpdateView()

	//TODO
	return nil
	// ˄
}

func (g *GameManager) GetShantenChecker() *ShantenChecker {
	// ˅
	return g.shantenChecker
	// ˄
}

func (g *GameManager) preparateGame() {
	// ˅
	g.determinateDealer()
	// ˄
}

func (g *GameManager) resetGame() {
	// ˅
	g.table.player1.tsumoriTile = nil
	g.table.player1.hand = []*Tile{}
	g.table.player1.kawa = []*Tile{}
	g.table.player1.openedTile1 = &OpenedTiles{}
	g.table.player1.openedTile2 = &OpenedTiles{}
	g.table.player1.openedTile3 = &OpenedTiles{}
	g.table.player1.openedTile4 = &OpenedTiles{}

	g.table.player2.tsumoriTile = nil
	g.table.player2.hand = []*Tile{}
	g.table.player2.kawa = []*Tile{}
	g.table.player2.openedTile1 = &OpenedTiles{}
	g.table.player2.openedTile2 = &OpenedTiles{}
	g.table.player2.openedTile3 = &OpenedTiles{}
	g.table.player2.openedTile4 = &OpenedTiles{}
	// ˄
}

func (g *GameManager) initializeGame() {
	// ˅
	ton := KAZE_TON
	nan := KAZE_NAN
	g.resetGame()
	g.table.tsumo.tiles = g.generateTiles()
	g.shuffleTiles(g.table.tsumo.tiles)
	g.table.tableStatus.ChichaPlayer = g.dealerPlayer
	g.table.tableStatus.PlayerWithTurn = g.dealerPlayer
	g.dealerPlayer.status.Kaze = &ton
	g.table.tableStatus.PlayerWithNotTurn = g.notDealerPlayer
	g.notDealerPlayer.status.Kaze = &nan
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
				id:      num,
				name:    name,
				num:     num,
				dora:    dora,
				suit:    &suit,
				akadora: akadora,
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
				id:      j + 10,
				name:    name,
				num:     num,
				dora:    dora,
				suit:    &suit,
				akadora: akadora,
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
				id:      num + 20,
				name:    name,
				num:     num,
				dora:    dora,
				suit:    &suit,
				akadora: akadora,
			})
		}

		name := ""
		num := 0
		dora := false
		akadora := false
		suit := NONE

		suit = TON
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 31
		tiles = append(tiles, &Tile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = NAN
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 32
		tiles = append(tiles, &Tile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = SHA
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 33
		tiles = append(tiles, &Tile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = PE
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 34
		tiles = append(tiles, &Tile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = HAKU
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 35
		tiles = append(tiles, &Tile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = HATSU
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 36
		tiles = append(tiles, &Tile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = CHUN
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 37
		tiles = append(tiles, &Tile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})
	}
	return tiles
	// ˄
}

func (g *GameManager) determinateDealer() {
	// ˅
	rand.Seed(time.Now().UnixNano())
	random := rand.Intn(2)
	if random == 1 {
		g.dealerPlayer = g.table.player1
		g.notDealerPlayer = g.table.player2
	} else {
		g.dealerPlayer = g.table.player2
		g.notDealerPlayer = g.table.player1
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
	tsumo := g.table.GetTsumo()
	for i := 0; i < 3; i++ {
		hand = g.dealerPlayer.GetHand()
		for j := 0; j < 4; j++ {
			tile = tsumo.Pop()
			hand = append(hand, tile)
		}
		g.dealerPlayer.SetHand(hand)

		hand = g.notDealerPlayer.GetHand()
		for j := 0; j < 4; j++ {
			tile = tsumo.Pop()
			hand = append(hand, tile)
		}
		g.notDealerPlayer.SetHand(hand)
	}

	hand = g.dealerPlayer.GetHand()
	tile = tsumo.Pop()
	hand = append(hand, tile)
	g.dealerPlayer.SetHand(hand)

	hand = g.notDealerPlayer.GetHand()
	tile = tsumo.Pop()
	hand = append(hand, tile)
	g.notDealerPlayer.SetHand(hand)
	// ˄
}

func (g *GameManager) appendKyushuKyuhaiOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	kyushukyuhai := OPERATOR_KYUSHUKYUHAI
	if player.status.KyushuKyuhai {
		operators = append(operators, &Operator{
			roomID:       g.table.GetID(),
			playerID:     player.GetID(),
			operatorType: &kyushukyuhai,
		})
	}
	return operators
	// ˄
}

func (g *GameManager) appendAnkanOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	ankan := OPERATOR_ANKAN
	tileIDs := handAndTsumoriTile(player)
	for _, tileID := range tileIDs {
		if tileIDs[tileID] == 4 {
			ankanTiles := []*Tile{}
			for _, tile := range player.GetHand() {
				if tile.GetID() == tileID {
					ankanTiles = append(ankanTiles, tile)
				}
			}
			operators = append(operators, &Operator{
				roomID:       g.table.GetID(),
				playerID:     player.GetID(),
				operatorType: &ankan,
				targetTiles:  ankanTiles,
			})
		}
	}
	return operators
	// ˄
}

func (g *GameManager) appendKakanOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	kakan := OPERATOR_KAKAN
	for _, openedTiles := range []*OpenedTiles{
		player.openedTile1,
		player.openedTile2,
		player.openedTile3,
		player.openedTile4,
	} {
		if !openedTiles.IsNil() &&
			*openedTiles.openType == OPEN_PON {
			for _, tile := range append(player.hand, player.GetTsumoriTile()) {
				if tile.GetID() == openedTiles.tiles[0].GetID() {
					operators = append(operators, &Operator{
						roomID:       g.table.GetID(),
						playerID:     player.GetID(),
						operatorType: &kakan,
						targetTiles:  []*Tile{tile},
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
	for _, tile := range append(player.hand, player.GetTsumoriTile()) {
		if tile.GetID() == 34 {
			operators = append(operators, &Operator{
				roomID:       g.table.GetID(),
				playerID:     player.GetID(),
				operatorType: &pe,
				targetTiles:  []*Tile{tile},
			})
		}
	}
	return operators
	// ˄
}

func (g *GameManager) appendTsumoAgariOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	tsumo := OPERATOR_TSUMO
	if g.GetShantenChecker().CheckCountOfShanten(player).Shanten == -1 {
		operators = append(operators, &Operator{
			roomID:       g.table.GetID(),
			playerID:     player.GetID(),
			operatorType: &tsumo,
			targetTiles:  []*Tile{player.GetTsumoriTile()},
		})
	}
	return operators
	// ˄
}

func (g *GameManager) appendReachOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	reach := OPERATOR_REACH
	canReach := true
	for _, openedTiles := range []OpenedTiles{
		*player.openedTile1,
		*player.openedTile2,
		*player.openedTile3,
		*player.openedTile4,
	} {
		if openedTiles.IsNil() {
			continue
		}
		if *openedTiles.openType == OPEN_PON ||
			*openedTiles.openType == OPEN_CHI ||
			*openedTiles.openType == OPEN_DAIMINKAN ||
			*openedTiles.openType == OPEN_KAKAN {
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
	for _, tile := range playerTemp.hand {
		handTemp = append(handTemp, tile)
	}

	for i, sutehai := range playerTemp.hand {
		playerTemp.hand = append(playerTemp.hand[:i], playerTemp.hand[i+1:]...)
		if g.shantenChecker.CheckCountOfShanten(&playerTemp).Shanten == 0 {
			operators = append(operators, &Operator{
				roomID:       g.table.GetID(),
				playerID:     player.GetID(),
				operatorType: &reach,
				targetTiles:  []*Tile{sutehai},
			})
		}
		playerTemp.hand = handTemp
	}

	playerTemp.GetTsumoriTile()
	tsumoriTileTemp := playerTemp.GetTsumoriTile()
	playerTemp.SetTsumoriTile(nil)
	if g.shantenChecker.CheckCountOfShanten(&playerTemp).Shanten == 0 {
		operators = append(operators, &Operator{
			roomID:       g.table.GetID(),
			playerID:     player.GetID(),
			operatorType: &reach,
			targetTiles:  []*Tile{tsumoriTileTemp},
		})
	}
	playerTemp.SetTsumoriTile(tsumoriTileTemp)

	return operators
	// ˄
}

func (g *GameManager) appendDahaiOperators(player *Player, operators []*Operator) []*Operator {
	// ˅
	dahai := OPERATOR_DAHAI
	operators = append(operators, &Operator{
		roomID:       g.table.GetID(),
		playerID:     player.GetID(),
		operatorType: &dahai,
		targetTiles:  []*Tile{player.GetTsumoriTile()},
	})

	if player.status.Reach {
		return operators
	}

	for _, tile := range player.GetHand() {
		operators = append(operators, &Operator{
			roomID:       g.table.GetID(),
			playerID:     player.GetID(),
			operatorType: &dahai,
			targetTiles:  []*Tile{tile},
		})
	}
	return operators
	// ˄
}

// ˅

func newGameManager(table *Table) *GameManager {
	return &GameManager{
		table:             table,
		receiveOperatorWG: &sync.WaitGroup{},
		okWG:              &sync.WaitGroup{},
		shantenChecker:    NewShantenChecker(),
		pointCalcrator:    &PointCalcrator{},
	}
}
func (m *GameManager) generateAgariMessage(player *Player) *Message {
	agarikei := m.shantenChecker.CheckCountOfShanten(player)
	point := m.pointCalcrator.CalcratePoint(player, agarikei, m.table, m.table.gameManager.shantenChecker.yakuList)
	magari := MessageAgari
	message := &Message{messageType: &magari}
	agari := &Agari{
		id:   player.GetID(),
		name: player.GetName(),
	}
	for _, tile := range player.GetHand() {
		agari.hand = append(agari.hand, tile)
	}
	if player.GetTsumoriTile != nil {
		agari.tsumoriTile = player.GetTsumoriTile()
	}
	if player.GetRonTile() != nil {
		agari.ronTile = player.GetRonTile()
	}
	agariOpenedTiles := []*OpenedTiles{
		agari.openedTile1,
		agari.openedTile1,
		agari.openedTile1,
		agari.openedTile1,
		agari.pe,
	}
	for i, openedTiles := range []*OpenedTiles{
		player.openedTile1,
		player.openedTile2,
		player.openedTile3,
		player.openedTile4,
		player.openedPe,
	} {
		(*agariOpenedTiles[i]) = (*openedTiles.ToOpenedTiles())
	}
	agari.point = &Point{}
	agari.point.Hu = point.Hu
	agari.point.Han = point.Han
	agari.point.Point = point.Point
	for _, yaku := range point.MatchYakus {
		agari.point.MatchYakus = append(agari.point.MatchYakus, yaku)
	}
	message.agari = agari
	return message
}

// ˄

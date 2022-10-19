// ˅
package nimar

import (
	"fmt"
	"math/rand"
	"time"
)

// ˄

type MGameManager struct {
	// ˅

	// ˄

	dealerPlayer *MPlayer

	notDealerPlayer *MPlayer

	table *MTable

	shantenChecker *ShantenChecker

	// ˅

	// ˄
}

func (m *MGameManager) ExecuteOperator(operator *Operator) error {
	// ˅
	var player *MPlayer
	var opponentPlayer *MPlayer

	if m.table.id != operator.RoomID {
		return fmt.Errorf("部屋IDとオペレータの対象部屋IDが一致しません。部屋ID:%s オペレータ対象部屋ID:%s", m.table.id, operator.RoomID)
	}
	player = m.table.GetPlayerByID(operator.PlayerID)
	if player == nil {
		return fmt.Errorf("部屋に命令を実行したプレイヤーがいません。プレイヤーID:%s", operator.PlayerID)
	}
	opponentPlayer = m.table.GetOpponentByID(operator.PlayerID)
	if opponentPlayer == nil {
		return fmt.Errorf("部屋に対局相手がいません。プレイヤーID:%s", operator.PlayerID)
	}

	switch operator.OperatorType {
	case OperatorType_OPERATOR_DRAW:
		if player.GetTsumoriTile() != nil {
			return fmt.Errorf("すでに引いているのに更にひこうとしています。プレイヤーID:%s", operator.PlayerID)
		}

		tsumo := m.table.GetTsumo()
		if !tsumo.CanPop() {
			return fmt.Errorf("ツモが18枚を下回ったので引けません。プレイヤーID:%s", operator.PlayerID)
		}
		player.SetTsumoriTile(tsumo.Pop())
	case OperatorType_OPERATOR_DAHAI:
		hand := player.GetHand()
		kawa := player.GetKawa()
		tileIndex := -1
		for i, tile := range hand {
			if tile.GetID() == int(operator.TargetTiles.Tiles[0].Id) {
				tileIndex = i
				break
			}
		}

		if tileIndex == -1 {
			return fmt.Errorf("手牌にない牌は捨てられません。プレイヤーID:%s 牌ID:%s", operator.PlayerID, operator.TargetTiles.Tiles[0].Id)
		}

		tile := hand[tileIndex]
		hand = append(hand[:tileIndex], hand[tileIndex+1:]...)
		hand = append(hand, player.GetTsumoriTile())
		kawa = append(kawa, tile)
		player.SetHand(hand)
		player.SetKawa(kawa)
		player.SetTsumoriTile(nil)

	case OperatorType_OPERATOR_START_GAME:
		//TODO
	case OperatorType_OPERATOR_SKIP:
		//TODO
	case OperatorType_OPERATOR_RON:
		player.SetRonTile(opponentPlayer.GetKawa()[len(opponentPlayer.GetKawa())-1])
		opponentPlayer.SetKawa(opponentPlayer.GetKawa()[:len(opponentPlayer.GetKawa())-1])
		//TODO
	case OperatorType_OPERATOR_PON:
		pon := OpenType_OPEN_PON
		openedTile := &MOpenedTiles{
			openType: &pon,
		}

		for i, targetTile := range operator.TargetTiles.GetTiles() {
			if i == 0 {
				if opponentPlayer.GetKawa()[len(opponentPlayer.GetKawa())-1].GetName() == targetTile.GetName() {
					opponentPlayer.SetKawa(opponentPlayer.GetKawa()[:len(opponentPlayer.GetKawa())-1])
					openedTile.tiles = append(openedTile.tiles, targetTile.ToMTile())
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
				openedTile.tiles = append(openedTile.tiles, targetTile.ToMTile())
			}
		}
		if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile1 = openedTile
		} else if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile2 = openedTile
		} else if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile3 = openedTile
		} else if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile4 = openedTile
		} else {
			return fmt.Errorf("ポンの完了に失敗しました。すでに4つ牌を開いています？")
		}
	case OperatorType_OPERATOR_CHI:
		chi := OpenType_OPEN_CHI
		openedTile := &MOpenedTiles{
			openType: &chi,
		}

		for i, targetTile := range operator.TargetTiles.GetTiles() {
			if i == 0 {
				if opponentPlayer.GetKawa()[len(opponentPlayer.GetKawa())-1].GetName() == targetTile.GetName() {
					opponentPlayer.SetKawa(opponentPlayer.GetKawa()[:len(opponentPlayer.GetKawa())-1])
					openedTile.tiles = append(openedTile.tiles, targetTile.ToMTile())
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
				openedTile.tiles = append(openedTile.tiles, targetTile.ToMTile())
			}
		}
		if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile1 = openedTile
		} else if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile2 = openedTile
		} else if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile3 = openedTile
		} else if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile4 = openedTile
		} else {
			return fmt.Errorf("チーの完了に失敗しました。すでに4つ牌を開いています？")
		}
	case OperatorType_OPERATOR_DAIMINKAN:
		daiminkan := OpenType_OPEN_DAIMINKAN
		openedTile := &MOpenedTiles{
			openType: &daiminkan,
		}

		for i, targetTile := range operator.TargetTiles.GetTiles() {
			if i == 0 {
				if opponentPlayer.GetKawa()[len(opponentPlayer.GetKawa())-1].GetName() == targetTile.GetName() {
					opponentPlayer.SetKawa(opponentPlayer.GetKawa()[:len(opponentPlayer.GetKawa())-1])
					openedTile.tiles = append(openedTile.tiles, targetTile.ToMTile())
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
				openedTile.tiles = append(openedTile.tiles, targetTile.ToMTile())
			}
		}
		if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile1 = openedTile
		} else if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile2 = openedTile
		} else if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile3 = openedTile
		} else if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile4 = openedTile
		} else {
			return fmt.Errorf("カンの完了に失敗しました。すでに4つ牌を開いています？")
		}

		player.SetTsumoriTile(m.table.tsumo.PopFromWanpai())
		if !m.table.GetTsumo().OpenNextKandora() {
			m.table.GetStatus().Sukaikan = true
		}

	case OperatorType_OPERATOR_TSUMO:
		//TODO
	case OperatorType_OPERATOR_ANKAN:
		ankan := OpenType_OPEN_ANKAN
		openedTile := &MOpenedTiles{
			openType: &ankan,
		}

		for _, targetTile := range operator.TargetTiles.GetTiles() {
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
			openedTile.tiles = append(openedTile.tiles, targetTile.ToMTile())
		}
		if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile1 = openedTile
		} else if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile2 = openedTile
		} else if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile3 = openedTile
		} else if *player.openedTile1.openType.Enum() == OpenType_OPEN_NULL {
			player.openedTile4 = openedTile
		} else {
			return fmt.Errorf("カンの完了に失敗しました。すでに4つ牌を開いています？")
		}

		player.SetTsumoriTile(m.table.tsumo.PopFromWanpai())
		if !m.table.GetTsumo().OpenNextKandora() {
			m.table.GetStatus().Sukaikan = true
		}
	case OperatorType_OPERATOR_KAKAN:
		//TODO
	case OperatorType_OPERATOR_PE:
		pe := OpenType_OPEN_PE
		openedTile := player.openedPe
		openedTile.openType = &pe
		for _, targetTile := range operator.TargetTiles.GetTiles() {
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
			openedTile.tiles = append(openedTile.tiles, targetTile.ToMTile())
		}
		player.openedPe = openedTile

		player.SetTsumoriTile(m.table.tsumo.PopFromWanpai())
	default:
		return fmt.Errorf("変なオペレータが渡されました。オペレータタイプ:%d", operator.OperatorType)
	}

	m.table.UpdateView()
	return nil
	// ˄
}

func (m *MGameManager) StartGame() {
	// ˅
	m.preparateGame()
	m.initializeGame()
	tsumo := m.table.GetTsumo()

	player := m.table.tableStatus.PlayerWithTurn
	opponentPlayer := m.table.tableStatus.PlayerWithNotTurn

	player.SetTsumoriTile(tsumo.Pop())
	if m.shantenChecker.GetYakuList()["九種九牌"].IsMatch(player, m.table, nil) {
		player.status.KyushuKyuhai = true
	} else {
		player.status.KyushuKyuhai = false
	}

	if m.shantenChecker.GetYakuList()["天和"].IsMatch(player, m.table, nil) {
		player.status.Tenho = true
	} else {
		player.status.Tenho = false
	}

	if m.shantenChecker.GetYakuList()["地和"].IsMatch(player, m.table, nil) {
		player.status.Chiho = true
	} else {
		player.status.Chiho = false
	}

	if m.table.tsumo.RemainTilesCount() <= 18 {
		player.status.Haitei = true
		opponentPlayer.status.Hotei = true
	} else {
		player.status.Haitei = false
		opponentPlayer.status.Hotei = false
	}

	//TODO

	// ˄
}

func (m *MGameManager) GetShantenChecker() *ShantenChecker {
	// ˅
	return m.shantenChecker
	// ˄
}

func (m *MGameManager) preparateGame() {
	// ˅
	m.determinateDealer()
	// ˄
}

func (m *MGameManager) resetGame() {
	// ˅
	m.table.player1.tsumoriTile = nil
	m.table.player1.hand = []*MTile{}
	m.table.player1.kawa = []*MTile{}
	m.table.player1.openedTile1 = &MOpenedTiles{}
	m.table.player1.openedTile2 = &MOpenedTiles{}
	m.table.player1.openedTile3 = &MOpenedTiles{}
	m.table.player1.openedTile4 = &MOpenedTiles{}

	m.table.player2.tsumoriTile = nil
	m.table.player2.hand = []*MTile{}
	m.table.player2.kawa = []*MTile{}
	m.table.player2.openedTile1 = &MOpenedTiles{}
	m.table.player2.openedTile2 = &MOpenedTiles{}
	m.table.player2.openedTile3 = &MOpenedTiles{}
	m.table.player2.openedTile4 = &MOpenedTiles{}

	m.table.tsumo.tiles = []*MTile{}
	// ˄
}

func (m *MGameManager) initializeGame() {
	// ˅
	m.resetGame()
	m.table.tsumo.tiles = m.generateTiles()
	m.shuffleTiles(m.table.tsumo.tiles)
	m.distributeTiles()
	m.table.tableStatus.ChichaPlayer = m.dealerPlayer
	m.table.tableStatus.PlayerWithTurn = m.dealerPlayer
	m.dealerPlayer.status.Kaze = Kaze_KAZE_TON.Enum()
	m.table.tableStatus.PlayerWithNotTurn = m.notDealerPlayer
	m.notDealerPlayer.status.Kaze = Kaze_KAZE_NAN.Enum()
	//TODO
	// ˄
}

func (m *MGameManager) getDealerPlayer() *MPlayer {
	// ˅
	return m.dealerPlayer
	// ˄
}

func (m *MGameManager) generateTiles() []*MTile {
	// ˅
	tiles := []*MTile{}

	for i := 1; i <= 4; i++ {
		// 萬子
		for j := 0; j < 2; j++ {
			name := ""
			num := 0
			dora := false
			suit := Suit_MANZU
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

			tiles = append(tiles, &MTile{
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
			suit := Suit_SOZU
			akadora := false

			name = fmt.Sprintf("%d%s%d", num, suit.ToString(), i)
			if i == 1 && j == 5 {
				akadora = true
				name += "赤ドラ"
			}

			tiles = append(tiles, &MTile{
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
			suit := Suit_PINZU
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

			tiles = append(tiles, &MTile{
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
		suit := Suit_NONE

		suit = Suit_TON
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 31
		tiles = append(tiles, &MTile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = Suit_NAN
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 32
		tiles = append(tiles, &MTile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = Suit_SHA
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 33
		tiles = append(tiles, &MTile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = Suit_PE
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 34
		tiles = append(tiles, &MTile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = Suit_HAKU
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 35
		tiles = append(tiles, &MTile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = Suit_HATSU
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 36
		tiles = append(tiles, &MTile{
			id:      num,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = Suit_CHUN
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		num = 37
		tiles = append(tiles, &MTile{
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

func (m *MGameManager) determinateDealer() {
	// ˅
	rand.Seed(time.Now().Unix())
	random := rand.Intn(2)
	if random == 0 {
		m.dealerPlayer = m.table.player1
		m.notDealerPlayer = m.table.player2
	} else {
		m.dealerPlayer = m.table.player2
		m.notDealerPlayer = m.table.player1
	}
	// ˄
}

func (m *MGameManager) shuffleTiles(tiles []*MTile) {
	// ˅
	temp := &MTile{}
	randomIndex := 1
	for i := 0; i < len(tiles); i++ {
		rand.Seed(time.Now().UnixNano())
		randomIndex = rand.Intn(len(tiles))

		temp = tiles[i]
		tiles[i] = tiles[randomIndex]
		tiles[randomIndex] = temp
	}
	// ˄
}

func (m *MGameManager) distributeTiles() {
	// ˅
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			m.dealerPlayer.hand = append(m.dealerPlayer.hand, m.table.tsumo.Pop())
		}
		for j := 0; j < 4; j++ {
			m.notDealerPlayer.hand = append(m.notDealerPlayer.hand, m.table.tsumo.Pop())
		}
	}
	m.dealerPlayer.hand = append(m.dealerPlayer.hand, m.table.tsumo.Pop())
	m.notDealerPlayer.hand = append(m.notDealerPlayer.hand, m.table.tsumo.Pop())
	// ˄
}

// ˅
func newGameManager(table *MTable) *MGameManager {
	return &MGameManager{
		table:          table,
		shantenChecker: NewShantenChecker(),
	}
}

// ˄

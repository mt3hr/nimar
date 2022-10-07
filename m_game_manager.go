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
		//TODO
	case OperatorType_OPERATOR_PON:
		//TODO
	case OperatorType_OPERATOR_CHI:
		//TODO
	case OperatorType_OPERATOR_DAIMINKAN:
		//TODO
	case OperatorType_OPERATOR_TSUMO:
		//TODO
	case OperatorType_OPERATOR_ANKAN:
		//TODO
	case OperatorType_OPERATOR_KAKAN:
		//TODO
	case OperatorType_OPERATOR_PE:
		//TODO
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
				id:      j,
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
				id:      j + 20,
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
		tiles = append(tiles, &MTile{
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = Suit_NAN
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		tiles = append(tiles, &MTile{
			id:      31,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = Suit_SHA
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		tiles = append(tiles, &MTile{
			id:      32,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = Suit_PE
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		tiles = append(tiles, &MTile{
			id:      33,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = Suit_HAKU
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		tiles = append(tiles, &MTile{
			id:      34,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = Suit_HATSU
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		tiles = append(tiles, &MTile{
			id:      35,
			name:    name,
			num:     num,
			dora:    dora,
			suit:    &suit,
			akadora: akadora,
		})

		suit = Suit_CHUN
		name = fmt.Sprintf("%s%d", suit.ToString(), i)
		tiles = append(tiles, &MTile{
			id:      36,
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
	} else if random == 1 {
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
		temp = tiles[i]
		rand.Seed(time.Now().Unix())
		randomIndex = rand.Intn(len(tiles))
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
	m.dealerPlayer.hand = append(m.dealerPlayer.hand, m.table.tsumo.Pop())
	// ˄
}

// ˅
func newGameManager(table *MTable) *MGameManager {
	return &MGameManager{
		table: table,
	}
}

// ˄

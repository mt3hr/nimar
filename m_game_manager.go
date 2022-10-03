// ˅
package nimar

import "fmt"

// ˄

type MGameManager struct {
	// ˅

	// ˄

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
			if tile.GetID() == operator.TargetTiles.Tiles[0].Id {
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

// ˅
func newGameManager(table *MTable) *MGameManager {
	return &MGameManager{
		table: table,
	}
}

// ˄

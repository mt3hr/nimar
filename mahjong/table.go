// ˅
package mahjong

// ˄

type Table struct {
	// ˅

	// ˄

	ID string

	Name string

	Tsumo *Tsumo

	Player2 *Player

	Player1 *Player

	GameManager *GameManager

	Status *TableStatus

	// ˅

	// ˄
}

func (t *Table) GetPlayerByID(playerID string) *Player {
	// ˅
	if t.Player1.ID == playerID {
		return t.Player1
	} else if t.Player2.ID == playerID {
		return t.Player2
	}
	return nil
	// ˄
}

func (t *Table) GetOpponentByID(playerID string) *Player {
	// ˅
	if t.Player1.ID == playerID {
		return t.Player2
	} else if t.Player2.ID == playerID {
		return t.Player1
	}
	return nil
	// ˄
}

func (t *Table) UpdateView() {
	// ˅
	Tsumo := &Tsumo{}
	for _, tile := range t.Tsumo.Tiles {
		Tsumo.Tiles = append(Tsumo.Tiles, tile)
	}
	p1 := t.Player1
	p2 := t.Player2

	gameTable := &Table{
		Tsumo:   Tsumo,
		Player1: p1,
		Player2: p2,
	}
	_ = gameTable
	//TODO
	// ˄
}

// ˅
func NewTable(roomID string, roomName string) *Table {
	ton := KAZE_TON
	Table := &Table{
		ID:    roomID,
		Name:  roomName,
		Tsumo: &Tsumo{},
		Status: &TableStatus{
			Sukaikan:      false,
			Kaze:          &ton,
			NumberOfKyoku: 1,
			NumberOfHonba: 0,
		},
	}
	manager := newGameManager(Table)
	Table.GameManager = manager
	return Table
}

// ˄

// ˅
package mahjong

import "encoding/json"

// ˄

type Table struct {
	// ˅

	// ˄

	ID string

	Name string

	Tsumo *Tsumo

	Player2 *Player

	Player1 *Player

	GameManager *GameManager `json:"-"`

	Status *TableStatus

	// ˅

	AllTiles []*Tile

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
	b, err := json.Marshal(t)
	if err != nil {
		panic(err)
	}
	_, err = t.Player1.GameTableWs.Write(b)
	if err != nil {
		panic(err)
	}

	_, err = t.Player2.GameTableWs.Write(b)
	if err != nil {
		panic(err)
	}
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

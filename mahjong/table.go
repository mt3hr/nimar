// ˅
package mahjong

// ˄

type Table struct {
	// ˅

	// ˄

	id string

	name string

	tsumo *Tsumo

	player2 *Player

	player1 *Player

	gameManager *GameManager

	tableStatus *TableStatus

	// ˅

	// ˄
}

func (t *Table) GetID() string {
	// ˅
	return t.id
	// ˄
}

func (t *Table) GetName() string {
	// ˅
	return t.name
	// ˄
}

func (t *Table) GetTsumo() *Tsumo {
	// ˅
	return t.tsumo
	// ˄
}

func (t *Table) GetPlayerByID(playerID string) *Player {
	// ˅
	if t.GetPlayer1().GetID() == playerID {
		return t.player1
	} else if t.GetPlayer2().GetID() == playerID {
		return t.player2
	}
	return nil
	// ˄
}

func (t *Table) GetOpponentByID(playerID string) *Player {
	// ˅
	if t.GetPlayer1().GetID() == playerID {
		return t.player2
	} else if t.GetPlayer2().GetID() == playerID {
		return t.player1
	}
	return nil
	// ˄
}

func (t *Table) GetPlayer1() *Player {
	// ˅
	return t.player1
	// ˄
}

func (t *Table) SetPlayer1(player1 *Player) {
	// ˅
	t.player1 = player1
	// ˄
}

func (t *Table) GetPlayer2() *Player {
	// ˅
	return t.player2
	// ˄
}

func (t *Table) SetPlayer2(player2 *Player) {
	// ˅
	t.player2 = player2
	// ˄
}

func (t *Table) GetGameManager() *GameManager {
	// ˅
	return t.gameManager
	// ˄
}

func (t *Table) GetStatus() *TableStatus {
	// ˅
	return t.tableStatus
	// ˄
}

func (t *Table) SetStatus(status *TableStatus) {
	// ˅
	t.tableStatus = status
	// ˄
}

func (t *Table) UpdateView() {
	// ˅
	tsumo := &Tsumo{}
	for _, tile := range t.GetTsumo().tiles {
		tsumo.tiles = append(tsumo.tiles, tile)
	}
	p1 := t.player1
	p2 := t.player2

	gameTable := &Table{
		tsumo:   tsumo,
		player1: p1,
		player2: p2,
	}
	_ = gameTable
	//TODO
	// ˄
}

// ˅
func NewTable(roomID string, roomName string) *Table {
	ton := KAZE_TON
	table := &Table{
		id:    roomID,
		name:  roomName,
		tsumo: &Tsumo{},
		tableStatus: &TableStatus{
			Sukaikan:      false,
			Kaze:          &ton,
			NumberOfKyoku: 1,
			NumberOfHonba: 0,
		},
	}
	manager := newGameManager(table)
	table.gameManager = manager
	return table
}

// ˄

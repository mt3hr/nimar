// ˅
package nimar

// ˄

type MTable struct {
	// ˅

	// ˄

	id string

	name string

	tsumo *MTsumo

	player2 *MPlayer

	player1 *MPlayer

	gameManager *MGameManager

	tableStatus *MTableStatus

	// ˅

	// ˄
}

func (m *MTable) GetID() string {
	// ˅
	return m.id
	// ˄
}

func (m *MTable) GetName() string {
	// ˅
	return m.name
	// ˄
}

func (m *MTable) GetTsumo() *MTsumo {
	// ˅
	return m.tsumo
	// ˄
}

func (m *MTable) GetPlayerByID(playerID string) *MPlayer {
	// ˅
	if m.GetPlayer1().GetID() == playerID {
		return m.player1
	} else if m.GetPlayer2().GetID() == playerID {
		return m.player2
	}
	return nil
	// ˄
}

func (m *MTable) GetOpponentByID(playerID string) *MPlayer {
	// ˅
	if m.GetPlayer1().GetID() == playerID {
		return m.player2
	} else if m.GetPlayer2().GetID() == playerID {
		return m.player1
	}
	return nil
	// ˄
}

func (m *MTable) GetPlayer1() *MPlayer {
	// ˅
	return m.player1
	// ˄
}

func (m *MTable) SetPlayer1(player1 *MPlayer) {
	// ˅
	m.player1 = player1
	// ˄
}

func (m *MTable) GetPlayer2() *MPlayer {
	// ˅
	return m.player2
	// ˄
}

func (m *MTable) SetPlayer2(player2 *MPlayer) {
	// ˅
	m.player2 = player2
	// ˄
}

func (m *MTable) GetGameManager() *MGameManager {
	// ˅
	return m.gameManager
	// ˄
}

func (m *MTable) GetStatus() *MTableStatus {
	// ˅
	return m.tableStatus
	// ˄
}

func (m *MTable) SetStatus(status *MTableStatus) {
	// ˅
	m.tableStatus = status
	// ˄
}

func (m *MTable) UpdateView() {
	// ˅
	t := &Tsumo{}
	for _, tile := range m.GetTsumo().tiles {
		t.Tiles.Tiles = append(t.Tiles.Tiles, tile.ToTile())
	}
	p1 := m.player1.ToPlayer()
	p2 := m.player2.ToPlayer()

	gameTable := &GameTable{
		Tsumo:   t,
		Player1: p1,
		Player2: p2,
	}

	if m.GetPlayer1().GetNimaRTableStreamServer() != nil {
		(*m.GetPlayer1().GetNimaRTableStreamServer()).Send(gameTable)
	}
	if m.GetPlayer2().GetNimaRTableStreamServer() != nil {
		(*m.GetPlayer2().GetNimaRTableStreamServer()).Send(gameTable)
	}
	// ˄
}

// ˅
func NewTable(roomID string, roomName string) *MTable {
	ton := Kaze_KAZE_TON
	table := &MTable{
		id:    roomID,
		name:  roomName,
		tsumo: &MTsumo{},
		tableStatus: &MTableStatus{
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

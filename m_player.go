// ˅
package nimar

// ˄

type MPlayer struct {
	// ˅

	// ˄

	name string

	id string

	hand []*MTile

	tsumoriTile *MTile

	openedTile1 *MOpenedTiles

	openedTile2 *MOpenedTiles

	openedTile3 *MOpenedTiles

	openedTile4 *MOpenedTiles

	nimarGameTableStreamServer *NimaR_GameTableStreamServer

	nimarOperatorStreamServer *NimaR_OperatorsStreamServer

	kawa []*MTile

	// ˅

	// ˄
}

func (m *MPlayer) GetName() string {
	// ˅
	return m.name
	// ˄
}

func (m *MPlayer) GetID() string {
	// ˅
	return m.id
	// ˄
}

func (m *MPlayer) GetNimaRTableStreamServer() *NimaR_GameTableStreamServer {
	// ˅
	return m.nimarGameTableStreamServer
	// ˄
}

func (m *MPlayer) GetNimaROperatorsStreamServer() *NimaR_OperatorsStreamServer {
	// ˅
	return m.nimarOperatorStreamServer
	// ˄
}

func (m *MPlayer) SetNimaROperatorsStreamServer(operatorStreamServer *NimaR_OperatorsStreamServer) {
	// ˅
	m.nimarOperatorStreamServer = m.nimarOperatorStreamServer
	// ˄
}

func (m *MPlayer) GetHand() []*MTile {
	// ˅
	return m.hand
	// ˄
}

func (m *MPlayer) SetHand(hand []*MTile) {
	// ˅
	m.hand = hand
	// ˄
}

func (m *MPlayer) GetKawa() []*MTile {
	// ˅
	return m.kawa
	// ˄
}

func (m *MPlayer) SetKawa(kawa []*MTile) {
	// ˅
	m.kawa = kawa
	// ˄
}

func (m *MPlayer) GetTsumoriTile() *MTile {
	// ˅
	return m.tsumoriTile
	// ˄
}

func (m *MPlayer) SetTsumoriTile(tile *MTile) {
	// ˅
	m.tsumoriTile = tile
	// ˄
}

func (m *MPlayer) ToPlayer() *Player {
	// ˅
	p := &Player{}
	p.Id = m.GetID()
	p.Name = m.GetName()

	hand := &Tiles{}
	for _, tile := range m.hand {
		hand.Tiles = append(hand.Tiles, tile.ToTile())
	}
	p.Hand = hand

	p.TsumoriTile = m.GetTsumoriTile().ToTile()

	p.OpenedTile1 = &OpenedTiles{
		OpenType: *m.openedTile1.openType,
	}
	for _, tile := range m.openedTile1.tiles {
		p.OpenedTile1.Tiles.Tiles = append(p.OpenedTile1.Tiles.Tiles, tile.ToTile())
	}
	p.OpenedTile2 = &OpenedTiles{
		OpenType: *m.openedTile2.openType,
	}
	for _, tile := range m.openedTile2.tiles {
		p.OpenedTile2.Tiles.Tiles = append(p.OpenedTile2.Tiles.Tiles, tile.ToTile())
	}
	p.OpenedTile3 = &OpenedTiles{
		OpenType: *m.openedTile3.openType,
	}
	for _, tile := range m.openedTile3.tiles {
		p.OpenedTile3.Tiles.Tiles = append(p.OpenedTile3.Tiles.Tiles, tile.ToTile())
	}
	p.OpenedTile4 = &OpenedTiles{
		OpenType: *m.openedTile4.openType,
	}
	for _, tile := range m.openedTile4.tiles {
		p.OpenedTile4.Tiles.Tiles = append(p.OpenedTile4.Tiles.Tiles, tile.ToTile())
	}

	for _, tile := range m.GetKawa() {
		p.Kawa.Tiles = append(p.Kawa.Tiles, tile.ToTile())
	}

	return p
	// ˄
}

// ˅
func NewPlayer(playerName string, playerID string, nimarGameTableStreaqmServer *NimaR_GameTableStreamServer) *MPlayer {
	return &MPlayer{
		name:                       playerName,
		id:                         playerID,
		nimarGameTableStreamServer: nimarGameTableStreaqmServer,
		hand:                       []*MTile{},
	}
}

// ˄

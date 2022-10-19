// ˅
package nimar

import (
	"sort"
)

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

	status *MPlayerStatus

	ronTile *MTile

	openedPe *MOpenedTiles

	nimarMessageStreamServer *NimaR_MessageStreamServer

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

func (m *MPlayer) GetNimaRMessageStreamServer() *NimaR_MessageStreamServer {
	// ˅
	return m.nimarMessageStreamServer
	// ˄
}

func (m *MPlayer) SetNimaRMessageStreamServer(nimaRMessageStreamServer *NimaR_MessageStreamServer) {
	// ˅
	m.nimarMessageStreamServer = nimaRMessageStreamServer
	// ˄
}

func (m *MPlayer) GetHand() []*MTile {
	// ˅
	return m.hand[:]
	// ˄
}

func (m *MPlayer) SetHand(hand []*MTile) {
	// ˅
	m.hand = hand[:]
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

func (m *MPlayer) GetRonTile() *MTile {
	// ˅
	return m.ronTile
	// ˄
}

func (m *MPlayer) SetRonTile(tile *MTile) {
	// ˅
	m.ronTile = tile
	// ˄
}

func (m *MPlayer) Rihai() {
	// ˅
	sort.Slice(m.hand, func(i, j int) bool {
		return m.hand[j].ThisIsBig(m.hand[i])
	})
	// ˄
}

func (m *MPlayer) IsMenzen() bool {
	// ˅
	for _, openedTile := range []*MOpenedTiles{
		m.openedTile1,
		m.openedTile2,
		m.openedTile3,
		m.openedTile4,
	} {
		if openedTile.IsNil() {
			continue
		}
		switch *openedTile.openType {
		case OpenType_OPEN_CHI:
			fallthrough
		case OpenType_OPEN_KAKAN:
			fallthrough
		case OpenType_OPEN_PON:
			return false
		}
	}
	return true
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

	p.OpenedTile1 = m.openedTile1.ToOpenedTiles()
	p.OpenedTile2 = m.openedTile2.ToOpenedTiles()
	p.OpenedTile3 = m.openedTile3.ToOpenedTiles()
	p.OpenedTile4 = m.openedTile4.ToOpenedTiles()

	return p
	// ˄
}

// ˅
func NewPlayer(playerName string, playerID string, nimarGameTableStreaqmServer *NimaR_GameTableStreamServer) *MPlayer {
	openTypeNull := OpenType_OPEN_NULL
	return &MPlayer{
		name:                       playerName,
		id:                         playerID,
		nimarGameTableStreamServer: nimarGameTableStreaqmServer,
		hand:                       []*MTile{},
		kawa:                       []*MTile{},
		tsumoriTile:                nil,
		openedTile1:                &MOpenedTiles{openType: &openTypeNull},
		openedTile2:                &MOpenedTiles{openType: &openTypeNull},
		openedTile3:                &MOpenedTiles{openType: &openTypeNull},
		openedTile4:                &MOpenedTiles{openType: &openTypeNull},
		status:                     &MPlayerStatus{},
	}
}

// ˄

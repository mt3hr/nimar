// ˅
package mahjong

import (
	"sort"
)

// ˄

type Player struct {
	// ˅

	// ˄

	name string

	id string

	hand []*Tile

	tsumoriTile *Tile

	openedTile1 *OpenedTiles

	openedTile2 *OpenedTiles

	openedTile3 *OpenedTiles

	openedTile4 *OpenedTiles

	kawa []*Tile

	status *PlayerStatus

	ronTile *Tile

	openedPe *OpenedTiles

	// ˅

	// ˄
}

func (p *Player) GetName() string {
	// ˅
	return p.name
	// ˄
}

func (p *Player) GetID() string {
	// ˅
	return p.id
	// ˄
}

func (p *Player) GetHand() []*Tile {
	// ˅
	return p.hand[:]
	// ˄
}

func (p *Player) SetHand(hand []*Tile) {
	// ˅
	p.hand = hand[:]
	// ˄
}

func (p *Player) GetKawa() []*Tile {
	// ˅
	return p.kawa
	// ˄
}

func (p *Player) SetKawa(kawa []*Tile) {
	// ˅
	p.kawa = kawa
	// ˄
}

func (p *Player) GetTsumoriTile() *Tile {
	// ˅
	return p.tsumoriTile
	// ˄
}

func (p *Player) SetTsumoriTile(tile *Tile) {
	// ˅
	p.tsumoriTile = tile
	// ˄
}

func (p *Player) GetRonTile() *Tile {
	// ˅
	return p.ronTile
	// ˄
}

func (p *Player) SetRonTile(tile *Tile) {
	// ˅
	p.ronTile = tile
	// ˄
}

func (p *Player) Rihai() {
	// ˅
	sort.Slice(p.hand, func(i, j int) bool {
		return p.hand[j].ThisIsBig(p.hand[i])
	})
	// ˄
}

func (p *Player) IsMenzen() bool {
	// ˅
	for _, openedTile := range []*OpenedTiles{
		p.openedTile1,
		p.openedTile2,
		p.openedTile3,
		p.openedTile4,
	} {
		if openedTile.IsNil() {
			continue
		}
		switch *openedTile.openType {
		case OPEN_CHI:
			fallthrough
		case OPEN_KAKAN:
			fallthrough
		case OPEN_PON:
			return false
		}
	}
	return true
	// ˄
}

// ˅
func NewPlayer(playerName string, playerID string) *Player {
	openTypeNull := OPEN_NULL
	return &Player{
		name:        playerName,
		id:          playerID,
		hand:        []*Tile{},
		kawa:        []*Tile{},
		tsumoriTile: nil,
		openedTile1: &OpenedTiles{openType: &openTypeNull},
		openedTile2: &OpenedTiles{openType: &openTypeNull},
		openedTile3: &OpenedTiles{openType: &openTypeNull},
		openedTile4: &OpenedTiles{openType: &openTypeNull},
		status:      &PlayerStatus{},
	}
}

// ˄

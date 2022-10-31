// ˅
package mahjong

import (
	"sort"

	"golang.org/x/net/websocket"
)

// ˄

type Player struct {
	// ˅

	// ˄

	Name string

	ID string

	Hand []*Tile

	TsumoriTile *Tile

	OpenedTile1 *OpenedTiles

	OpenedTile2 *OpenedTiles

	OpenedTile3 *OpenedTiles

	OpenedTile4 *OpenedTiles

	Kawa []*Tile

	Status *PlayerStatus

	RonTile *Tile

	OpenedPe *OpenedTiles

	// ˅

	GameTableWs *websocket.Conn
	OperatorWs  *websocket.Conn

	// ˄
}

func (p *Player) Rihai() {
	// ˅
	nilWithoutHand := []*Tile{}
	for _, tile := range p.Hand {
		if tile == nil {
			continue
		}
		nilWithoutHand = append(nilWithoutHand, tile)
	}
	p.Hand = nilWithoutHand
	sort.Slice(p.Hand, func(i, j int) bool {
		return p.Hand[j].ThisIsBig(p.Hand[i])
	})
	// ˄
}

func (p *Player) IsMenzen() bool {
	// ˅
	for _, OpenedTile := range []*OpenedTiles{
		p.OpenedTile1,
		p.OpenedTile2,
		p.OpenedTile3,
		p.OpenedTile4,
	} {
		if OpenedTile.IsNil() {
			continue
		}
		switch *OpenedTile.OpenType {
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
	OpenTypeNull := OPEN_NULL
	return &Player{
		Name:        playerName,
		ID:          playerID,
		Hand:        []*Tile{},
		Kawa:        []*Tile{},
		TsumoriTile: nil,
		OpenedTile1: &OpenedTiles{OpenType: &OpenTypeNull},
		OpenedTile2: &OpenedTiles{OpenType: &OpenTypeNull},
		OpenedTile3: &OpenedTiles{OpenType: &OpenTypeNull},
		OpenedTile4: &OpenedTiles{OpenType: &OpenTypeNull},
		OpenedPe:    &OpenedTiles{OpenType: &OpenTypeNull},
		Status:      &PlayerStatus{},
	}
}

// ˄

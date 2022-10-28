package mahjong

import (
	"fmt"
)

type TileIDs [39]int

func (t *TileIDs) Reset() {
	for i := range t {
		t[i] = 0
	}
}

func (t *TileIDs) IsEmpty() bool {
	if t == nil {
		return true
	}
	for _, cnt := range t {
		if cnt != 0 {
			return false
		}
	}
	return true
}
func (t *TileIDs) String() string {
	if t == nil {
		return ""
	}

	alltiles := NewTable("", "").GameManager.GenerateTiles()
	tiles := []*Tile{}

	for tileid, count := range t {
	loop:
		for _, tile := range alltiles {
			if tile.ID == tileid {
				for i := 0; i < count; i++ {
					tiles = append(tiles, tile)
					break loop
				}
			}
		}
	}
	str := ""
	for _, tile := range tiles {
		str += fmt.Sprintf("[%+v]", tile.Name)
	}
	return str
}
func (t *TileIDs) Clone() *TileIDs {
	tileIDs := &TileIDs{}
	if t == nil {
		return tileIDs
	}
	for i, id := range t {
		tileIDs[i] = id
	}
	return tileIDs
}

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
	alltiles := NewTable("", "").GameManager.GenerateTiles()
	tiles := []*Tile{}

	for tileid, count := range t {
		for _, tile := range alltiles {
			if tile.ID == tileid {
				for i := 0; i < count; i++ {
					tiles = append(tiles, tile)
				}
			}
		}
	}

	if t == nil {
		return ""
	}
	str := ""
	for i, cnt := range t {
		for j := 0; j < cnt; j++ {
			str += fmt.Sprintf("[%+v]", i)
		}
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

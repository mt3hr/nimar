// ˅
package mahjong

// ˄

type Tsumo struct {
	// ˅

	// ˄

	Tiles []*Tile

	// ˅

	// ˄
}

func (t *Tsumo) RemainTilesCount() int {
	// ˅
	cnt := 0
	for _, tile := range t.Tiles {
		if tile != nil {
			cnt++
		}
	}
	return cnt
	// ˄
}

func (t *Tsumo) CanPop() bool {
	// ˅
	return t.RemainTilesCount() >= 18
	// ˄
}

func (t *Tsumo) Pop() *Tile {
	// ˅
	tile := t.Tiles[0]
	t.Tiles = t.Tiles[1:]
	return tile
	// ˄
}

func (t *Tsumo) OpenNextKandora() bool {
	// ˅
	//TODO
	return false
	// ˄
}

func (t *Tsumo) PopFromWanpai() *Tile {
	// ˅
	for i := len(t.Tiles) - 18; i <= len(t.Tiles)-8; i++ {
		if t.Tiles[i] != nil {
			tile := t.Tiles[i]
			t.Tiles[i] = nil
			return tile
		}
	}
	return nil
	// ˄
}

// ˅

// ˄

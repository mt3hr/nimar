// ˅
package mahjong

// ˄

type Tsumo struct {
	// ˅

	// ˄

	tiles []*Tile

	// ˅

	// ˄
}

func (t *Tsumo) InitTiles() {
	// ˅
	//TODO
	// ˄
}

func (t *Tsumo) RemainTilesCount() int {
	// ˅
	cnt := 0
	for _, tile := range t.tiles {
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
	if t.CanPop() {
		defer func() { t.tiles = t.tiles[1:] }()
		return t.tiles[0]
	}
	return nil
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
	//TODO
	return nil
	// ˄
}

// ˅

// ˄

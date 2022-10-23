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
	if t.CanPop() {
		defer func() { t.Tiles = t.Tiles[1:] }()
		return t.Tiles[0]
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

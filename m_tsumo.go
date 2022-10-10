// ˅
package nimar

// ˄

type MTsumo struct {
	// ˅

	// ˄

	tiles []*MTile

	// ˅

	// ˄
}

func (m *MTsumo) InitTiles() {
	// ˅
	//TODO
	// ˄
}

func (m *MTsumo) RemainTilesCount() int {
	// ˅
	cnt := 0
	for _, tile := range m.tiles {
		if tile != nil {
			cnt++
		}
	}
	return cnt
	// ˄
}

func (m *MTsumo) CanPop() bool {
	// ˅
	return m.RemainTilesCount() >= 18
	// ˄
}

func (m *MTsumo) Pop() *MTile {
	// ˅
	tile := m.tiles[0]
	m.tiles = m.tiles[1:]
	return tile
	// ˄
}

func (m *MTsumo) OpenNextKandora() bool {
	// ˅
	//TODO
	return false
	// ˄
}

func (m *MTsumo) PopFromWanpai() *MTile {
	// ˅
	//TODO
	return nil
	// ˄
}

// ˅

// ˄

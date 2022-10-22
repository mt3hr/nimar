// ˅
package mahjong

// ˄

type OpenedTiles struct {
	// ˅

	// ˄

	tiles []*Tile

	openType *OpenType

	// ˅

	// ˄
}

// ˅

func (m *OpenedTiles) IsNil() bool {
	if m == nil || m.openType == nil {
		return true
	}
	return false
}

func (m *OpenedTiles) ToOpenedTiles() *OpenedTiles {
	if m.IsNil() {
		return nil
	}
	openedTile := &OpenedTiles{}
	if *m.openType != OPEN_NULL {
		openedTile.openType = m.openType
		for _, tile := range m.tiles {
			openedTile.tiles = append(openedTile.tiles, tile)
		}
	}
	return openedTile
}

// ˄

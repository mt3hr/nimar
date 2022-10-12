// ˅
package nimar

// ˄

type MOpenedTiles struct {
	// ˅

	// ˄

	tiles []*MTile

	openType *OpenType

	// ˅

	// ˄
}

// ˅

func (m *MOpenedTiles) IsNil() bool {
	if m == nil || m.openType == nil {
		return true
	}
	return false
}

func (m *MOpenedTiles) ToOpenedTiles() *OpenedTiles {
	if m.IsNil() {
		return nil
	}
	openedTile := &OpenedTiles{}
	if *m.openType != OpenType_OPEN_NULL {
		openedTile.OpenType = *m.openType
		for _, tile := range m.tiles {
			openedTile.Tiles.Tiles = append(openedTile.Tiles.Tiles, tile.ToTile())
		}
	}
	return openedTile
}

// ˄

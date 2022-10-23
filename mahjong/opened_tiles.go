// ˅
package mahjong

// ˄

type OpenedTiles struct {
	// ˅

	// ˄

	Tiles []*Tile

	OpenType *OpenType

	// ˅

	// ˄
}

// ˅

func (m *OpenedTiles) IsNil() bool {
	if m == nil || m.OpenType == nil {
		return true
	}
	return false
}

func (m *OpenedTiles) ToOpenedTiles() *OpenedTiles {
	if m.IsNil() {
		return nil
	}
	OpenedTile := &OpenedTiles{}
	if *m.OpenType != OPEN_NULL {
		OpenedTile.OpenType = m.OpenType
		for _, tile := range m.Tiles {
			OpenedTile.Tiles = append(OpenedTile.Tiles, tile)
		}
	}
	return OpenedTile
}

// ˄

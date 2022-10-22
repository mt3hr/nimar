package mahjong

type TileIDs [38]int

func (t TileIDs) Reset() {
	for i := range t {
		t[i] = 0
	}
}

func (t TileIDs) IsEmpty() bool {
	for _, cnt := range t {
		if cnt != 0 {
			return false
		}
	}
	return true
}

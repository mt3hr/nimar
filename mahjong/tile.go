// ˅
package mahjong

// ˄

type Tile struct {
	// ˅

	// ˄

	ID int

	Name string

	Num int

	Dora int

	Akadora bool

	Suit *Suit

	// ˅

	DoraHyoujiHai    bool
	UraDoraHyoujiHai bool
	UraDora          int

	// ˄
}

func (t *Tile) ThisIsBig(tile *Tile) bool {
	// ˅
	if int(*t.Suit) == int(*tile.Suit) {
		return t.Num > tile.Num
	}
	return int(*t.Suit) > int(*tile.Suit)
	// ˄
}

// ˅

// ˄

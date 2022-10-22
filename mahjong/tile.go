// ˅
package mahjong

// ˄

type Tile struct {
	// ˅

	// ˄

	id int

	name string

	num int

	dora bool

	akadora bool

	suit *Suit

	// ˅

	// ˄
}

func (t *Tile) GetID() int {
	// ˅
	return t.id
	// ˄
}

func (t *Tile) GetName() string {
	// ˅
	return t.name
	// ˄
}

func (t *Tile) GetSuit() Suit {
	// ˅
	return *t.suit
	// ˄
}

func (t *Tile) GetNum() int {
	// ˅
	return t.num
	// ˄
}

func (t *Tile) IsDora() bool {
	// ˅
	return t.dora
	// ˄
}

func (t *Tile) IsAkadora() bool {
	// ˅
	return t.akadora
	// ˄
}

func (t *Tile) ThisIsBig(tile *Tile) bool {
	// ˅
	if int(t.GetSuit()) == int(tile.GetSuit()) {
		return t.GetNum() > tile.GetNum()
	}
	return int(t.GetSuit()) > int(tile.GetSuit())
	// ˄
}

// ˅

// ˄

// ˅
package nimar

// ˄

type MTile struct {
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

func (m *MTile) GetID() int {
	// ˅
	return m.id
	// ˄
}

func (m *MTile) GetName() string {
	// ˅
	return m.name
	// ˄
}

func (m *MTile) GetSuit() Suit {
	// ˅
	return *m.suit
	// ˄
}

func (m *MTile) GetNum() int {
	// ˅
	return m.num
	// ˄
}

func (m *MTile) IsDora() bool {
	// ˅
	return m.dora
	// ˄
}

func (m *MTile) IsAkadora() bool {
	// ˅
	return m.akadora
	// ˄
}

func (m *MTile) ToTile() *Tile {
	// ˅
	if m == nil {
		return nil
	}
	return &Tile{
		Name:    m.GetName(),
		Id:      int64(m.GetID()),
		Num:     int64(m.GetNum()),
		Suit:    m.GetSuit(),
		Dora:    m.IsDora(),
		Akadora: m.IsAkadora(),
	}
	// ˄
}

func (m *MTile) ThisIsBig(tile *MTile) bool {
	// ˅
	if m.GetSuit() == tile.GetSuit() {
		return m.GetNum() > tile.GetNum()
	}
	return m.GetSuit() > tile.GetSuit()
	// ˄
}

// ˅

func (t *Tile) ToMTile() *MTile {
	if t == nil {
		return nil
	}
	return &MTile{
		name:    t.GetName(),
		id:      int(t.GetId()),
		num:     int(t.GetNum()),
		suit:    t.GetSuit().Enum(),
		dora:    t.GetDora(),
		akadora: t.Akadora,
	}
}

// ˄

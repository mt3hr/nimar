// ˅
package nimar

// ˄

type MTile struct {
	// ˅

	// ˄

	id string

	num int

	dora bool

	akadora bool

	suit *Suit

	// ˅

	// ˄
}

func (m *MTile) GetID() string {
	// ˅
	return m.id
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
	return &Tile{
		Id:      m.GetID(),
		Num:     int64(m.GetNum()),
		Suit:    m.GetSuit(),
		Dora:    m.IsDora(),
		Akadora: m.IsAkadora(),
	}
	// ˄
}

// ˅

// ˄

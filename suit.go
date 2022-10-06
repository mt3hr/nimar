package nimar

func (s Suit) ToString() string {
	switch s {
	case Suit_NONE:
		return ""
	case Suit_MANZU:
		return "萬"
	case Suit_SOZU:
		return "索"
	case Suit_PINZU:
		return "筒"
	case Suit_TON:
		return "東"
	case Suit_NAN:
		return "南"
	case Suit_SHA:
		return "西"
	case Suit_PE:
		return "北"
	case Suit_HAKU:
		return "白"
	case Suit_HATSU:
		return "發"
	case Suit_CHUN:
		return "中"
	}
	return ""
}

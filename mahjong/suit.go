// ˅
package mahjong

// ˄

type Suit int

const (
	NONE Suit = iota + 1

	MANZU

	SOZU

	PINZU

	TON

	NAN

	SHA

	PE

	HAKU

	HATSU

	CHUN

	// ˅

	// ˄
)

// ˅

func (s Suit) ToString() string {
	switch s {
	case MANZU:
		return "萬"
	case SOZU:
		return "索"
	case PINZU:
		return "筒"
	case TON:
		return "東"
	case NAN:
		return "南"
	case SHA:
		return "西"
	case PE:
		return "北"
	case HAKU:
		return "白"
	case HATSU:
		return "發"
	case CHUN:
		return "中"
	case NONE:
		fallthrough
	default:
		return ""
	}
}

// ˄

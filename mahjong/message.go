// ˅
package mahjong

// ˄

type Message struct {
	// ˅

	// ˄

	MessageType *MessageType

	Agari *Agari

	Ryukyoku *Ryukyoku

	MatchResult *MatchResult

	// ˅
	NagashiMangan *NagashiMangan

	// ˄
}

// ˅
type NagashiMangan struct {
	Player1IsNagashiMangan bool
	Player2IsNagashiMangan bool
	Player1Bappu           int
	Player2Bappu           int
}

// ˄

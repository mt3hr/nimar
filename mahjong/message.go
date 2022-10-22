// ˅
package mahjong

// ˄

type Message struct {
	// ˅

	// ˄

	messageType *MessageType

	agari *Agari

	// ˅

	// ˄
}

func (m *Message) GetMessageType() *MessageType {
	// ˅
	return m.messageType
	// ˄
}

func (m *Message) GetAgari() *Agari {
	// ˅
	return m.agari
	// ˄
}

// ˅

// ˄

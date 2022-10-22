// ˅
package mahjong

// ˄

type Operator struct {
	// ˅

	// ˄

	roomID string

	playerID string

	targetTiles []*Tile

	operatorType *OperatorType

	// ˅

	// ˄
}

func (o *Operator) GetRoomID() string {
	// ˅
	return o.roomID
	// ˄
}

func (o *Operator) GetPlayerID() string {
	// ˅
	return o.playerID
	// ˄
}

func (o *Operator) GetOperatorType() *OperatorType {
	// ˅
	return o.operatorType
	// ˄
}

func (o *Operator) GetTargetTiles() []*Tile {
	// ˅
	return o.targetTiles
	// ˄
}

// ˅

// ˄

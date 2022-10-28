// ˅
package mahjong

// ˄

type Agarikei struct {
	// ˅

	// ˄

	Janto *TileIDs

	Mentsu1 *TileIDs

	Mentsu2 *TileIDs

	Mentsu3 *TileIDs

	Mentsu4 *TileIDs

	Mentsu1Type MentsuType

	Mentsu2Type MentsuType

	Mentsu3Type MentsuType

	Mentsu4Type MentsuType

	MachiHai map[int]interface{}

	Machi *Machi

	// ˅

	JantoType MentsuType
	// ˄
}

// ˅

func (a Agarikei) String() string {
	str := ""
	for _, mentsus := range []*TileIDs{
		a.Janto,
		a.Mentsu1,
		a.Mentsu2,
		a.Mentsu3,
		a.Mentsu4,
	} {
		if !mentsus.IsEmpty() {
			str += mentsus.String() + " "
		}
	}
	return str
}

func (a Agarikei) Clone() *Agarikei {
	agarikei := &Agarikei{}
	agarikei.Janto = a.Janto.Clone()
	agarikei.Mentsu1 = a.Mentsu1.Clone()
	agarikei.Mentsu2 = a.Mentsu2.Clone()
	agarikei.Mentsu3 = a.Mentsu3.Clone()
	agarikei.Mentsu4 = a.Mentsu4.Clone()
	agarikei.Mentsu1Type = a.Mentsu1Type
	agarikei.Mentsu2Type = a.Mentsu2Type
	agarikei.Mentsu3Type = a.Mentsu3Type
	agarikei.Mentsu4Type = a.Mentsu4Type
	agarikei.JantoType = a.JantoType

	agarikei.MachiHai = map[int]interface{}{}
	for id := range a.MachiHai {
		agarikei.MachiHai[id] = struct{}{}
	}

	if a.Machi != nil {
		machi := *a.Machi
		agarikei.Machi = &machi
	}

	return agarikei
}

// ˄

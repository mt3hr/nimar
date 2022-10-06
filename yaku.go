// ˅
package nimar

// ˄

type Yaku interface {
	IsMatch(player *MPlayer) bool

	GetName() string

	NumberOfHan() int

	NumberOfHanWhenNaki() int

	// ˅

	// ˄
}

// ˅
type Tanyao struct{}
type Reach struct{}
type Ippatsu struct{}
type MenzenTsumo struct{}
type Chankan struct{}
type Rinshan struct{}
type Haitei struct{}
type Houtei struct{}
type DoubleReach struct{}
type Chitoitsu struct{}
type DabuTon struct{}
type DabuNan struct{}
type DabuSha struct{}
type DabuPe struct{}
type SanAnko struct{}
type SanKantsu struct{}
type SuankoTanki struct{}
type JunseiChuren struct{}
type Kokushi13 struct{}
type Pinhu struct{}
type Haku struct{}
type Hatsu struct{}
type Chun struct{}
type Ton struct{}
type Nan struct{}
type Sha struct{}
type Pe struct{}
type Toitoi struct{}
type SanshokuDoukou struct{}
type SanshokuDoujun struct{}
type Honroto struct{}
type Ikkitsuukan struct{}
type Chanta struct{}
type Shousangen struct{}
type Honitsu struct{}
type Junchan struct{}
type Chinitsu struct{}
type Ryuiso struct{}
type Daisangen struct{}
type Shosushi struct{}
type Tsuiso struct{}
type Kokushi struct{}
type Suanko struct{}
type Chinroto struct{}
type Sankantsu struct{}
type Daisushi struct{}
type Churenpoto struct{}
type Ryanpeko struct{}
type Ipeko struct{}
type Nagashimangan struct{}
type Tenho struct{}
type Chiho struct{}
type Renho struct{}

// ˄

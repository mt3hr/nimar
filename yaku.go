// ˅
package nimar

// ˄

type Yaku interface {
	IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool

	GetName() string

	NumberOfHan() int

	NumberOfHanWhenNaki() int

	// ˅

	// ˄
}

// ˅

func handAndTsumoriTile(player *MPlayer) TileIDs {
	tileIDs := [38]int{}
	for _, tile := range player.hand {
		tileIDs[tile.GetID()]++
	}
	if player.tsumoriTile != nil {
		tileIDs[player.GetTsumoriTile().GetID()]++
	}
	return tileIDs
}

type Tanyao struct {
	han     int
	nakihan int
}
type Reach struct {
	han     int
	nakihan int
}
type Ippatsu struct {
	han     int
	nakihan int
}
type MenzenTsumo struct {
	han     int
	nakihan int
}
type Chankan struct {
	han     int
	nakihan int
}
type Rinshan struct {
	han     int
	nakihan int
}
type Haitei struct {
	han     int
	nakihan int
}
type Houtei struct {
	han     int
	nakihan int
}
type DoubleReach struct {
	han     int
	nakihan int
}
type Chitoitsu struct {
	han     int
	nakihan int
}
type DabuTon struct {
	han     int
	nakihan int
}
type DabuNan struct {
	han     int
	nakihan int
}
type DabuSha struct {
	han     int
	nakihan int
}
type DabuPe struct {
	han     int
	nakihan int
}
type SanAnko struct {
	han     int
	nakihan int
}
type SanKantsu struct {
	han     int
	nakihan int
}
type SuankoTanki struct {
	han     int
	nakihan int
}
type JunseiChuren struct {
	han     int
	nakihan int
}
type Kokushi13 struct {
	han     int
	nakihan int
}
type Pinhu struct {
	han     int
	nakihan int
}
type Haku struct {
	han     int
	nakihan int
}
type Hatsu struct {
	han     int
	nakihan int
}
type Chun struct {
	han     int
	nakihan int
}
type Ton struct {
	han     int
	nakihan int
}
type Nan struct {
	han     int
	nakihan int
}
type Sha struct {
	han     int
	nakihan int
}
type Pe struct {
	han     int
	nakihan int
}
type Toitoi struct {
	han     int
	nakihan int
}
type SanshokuDoukou struct {
	han     int
	nakihan int
}
type SanshokuDoujun struct {
	han     int
	nakihan int
}
type Honroto struct {
	han     int
	nakihan int
}
type Ikkitsuukan struct {
	han     int
	nakihan int
}
type Chanta struct {
	han     int
	nakihan int
}
type Shousangen struct {
	han     int
	nakihan int
}
type Honitsu struct {
	han     int
	nakihan int
}
type Junchan struct {
	han     int
	nakihan int
}
type Chinitsu struct {
	han     int
	nakihan int
}
type Ryuiso struct {
	han     int
	nakihan int
}
type Daisangen struct {
	han     int
	nakihan int
}
type Shosushi struct {
	han     int
	nakihan int
}
type Tsuiso struct {
	han     int
	nakihan int
}
type Kokushi struct {
	han     int
	nakihan int
}
type Suanko struct {
	han     int
	nakihan int
}
type Chinroto struct {
	han     int
	nakihan int
}
type Sukantsu struct {
	han     int
	nakihan int
}
type Daisushi struct {
	han     int
	nakihan int
}
type Churenpoto struct {
	han     int
	nakihan int
}
type Ryanpeko struct {
	han     int
	nakihan int
}
type Ipeko struct {
	han     int
	nakihan int
}
type Nagashimangan struct {
	han     int
	nakihan int
}
type Tenho struct {
	han     int
	nakihan int
}
type Chiho struct {
	han     int
	nakihan int
}
type Renho struct {
	han     int
	nakihan int
}

func (t *Tanyao) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	tiles := handAndTsumoriTile(player)

	for i := range tiles {
		if tiles[i] != 0 && (i == 1 || i == 9 || i == 11 || i == 19 || i == 21 || tiles[i] == 29 || (i >= 31 && i <= 37)) {
			return false
		}
	}
	return true
}

func (t *Tanyao) GetName() string {
	return "断么九"
}

func (t *Tanyao) NumberOfHan() int {
	return t.han
}

func (t *Tanyao) NumberOfHanWhenNaki() int {
	return t.nakihan
}

func (r *Reach) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	return player.status.Reach
}

func (r *Reach) GetName() string {
	return "立直"
}

func (r *Reach) NumberOfHan() int {
	return r.han
}

func (r *Reach) NumberOfHanWhenNaki() int {
	return r.nakihan
}

func (i *Ippatsu) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	return player.status.Ippatsu && player.IsMenzen()
}

func (i *Ippatsu) GetName() string {
	return "一発"
}

func (i *Ippatsu) NumberOfHan() int {
	return i.han
}

func (i *Ippatsu) NumberOfHanWhenNaki() int {
	return i.nakihan
}

func (m *MenzenTsumo) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	return player.status.Ippatsu && player.IsMenzen()
}

func (m *MenzenTsumo) GetName() string {
	return "門前自摸"
}

func (m *MenzenTsumo) NumberOfHan() int {
	return m.han
}

func (m *MenzenTsumo) NumberOfHanWhenNaki() int {
	return m.nakihan
}

func (c *Chankan) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (c *Chankan) GetName() string {
	return "槍槓"
}

func (c *Chankan) NumberOfHan() int {
	return c.han
}

func (c *Chankan) NumberOfHanWhenNaki() int {
	return c.nakihan
}

func (r *Rinshan) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (r *Rinshan) GetName() string {
	return "嶺上開花"
}

func (r *Rinshan) NumberOfHan() int {
	return r.han
}

func (r *Rinshan) NumberOfHanWhenNaki() int {
	return r.nakihan
}

func (h *Haitei) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (h *Haitei) GetName() string {
	return "海底撈月"
}

func (h *Haitei) NumberOfHan() int {
	return h.han
}

func (h *Haitei) NumberOfHanWhenNaki() int {
	return h.nakihan
}

func (h *Houtei) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (h *Houtei) GetName() string {
	return "河底撈魚"
}

func (h *Houtei) NumberOfHan() int {
	return h.han
}

func (h *Houtei) NumberOfHanWhenNaki() int {
	return h.nakihan
}

func (d *DoubleReach) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (d *DoubleReach) GetName() string {
	return "ダブルリーチ"
}

func (d *DoubleReach) NumberOfHan() int {
	return d.han
}

func (d *DoubleReach) NumberOfHanWhenNaki() int {
	return d.nakihan
}

func (c *Chitoitsu) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (c *Chitoitsu) GetName() string {
	return "七対子"
}

func (c *Chitoitsu) NumberOfHan() int {
	return c.han
}

func (c *Chitoitsu) NumberOfHanWhenNaki() int {
	return c.nakihan
}

func (d *DabuTon) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (d *DabuTon) GetName() string {
	return "連風東"
}

func (d *DabuTon) NumberOfHan() int {
	return d.han
}

func (d *DabuTon) NumberOfHanWhenNaki() int {
	return d.nakihan
}

func (d *DabuNan) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (d *DabuNan) GetName() string {
	return "連風南"
}

func (d *DabuNan) NumberOfHan() int {
	return d.han
}

func (d *DabuNan) NumberOfHanWhenNaki() int {
	return d.nakihan
}

func (d *DabuSha) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (d *DabuSha) GetName() string {
	return "連風西"
}

func (d *DabuSha) NumberOfHan() int {
	return d.han
}

func (d *DabuSha) NumberOfHanWhenNaki() int {
	return d.nakihan
}

func (d *DabuPe) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (d *DabuPe) GetName() string {
	return "連風北"
}

func (d *DabuPe) NumberOfHan() int {
	return d.han
}

func (d *DabuPe) NumberOfHanWhenNaki() int {
	return d.nakihan
}

func (s *SanAnko) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (s *SanAnko) GetName() string {
	return "三暗刻"
}

func (s *SanAnko) NumberOfHan() int {
	return s.han
}

func (s *SanAnko) NumberOfHanWhenNaki() int {
	return s.nakihan
}

func (s *SanKantsu) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (s *SanKantsu) GetName() string {
	panic("not implemented") // TODO: Implement
	return "三槓子"
}

func (s *SanKantsu) NumberOfHan() int {
	return s.han
}

func (s *SanKantsu) NumberOfHanWhenNaki() int {
	return s.nakihan
}

func (s *SuankoTanki) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (s *SuankoTanki) GetName() string {
	return "四暗刻単騎"
}

func (s *SuankoTanki) NumberOfHan() int {
	return s.han
}

func (s *SuankoTanki) NumberOfHanWhenNaki() int {
	return s.nakihan
}

func (j *JunseiChuren) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (j *JunseiChuren) GetName() string {
	return "純正九蓮宝燈"
}

func (j *JunseiChuren) NumberOfHan() int {
	return j.han
}

func (j *JunseiChuren) NumberOfHanWhenNaki() int {
	return j.nakihan
}

func (k *Kokushi13) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (k *Kokushi13) GetName() string {
	return "国士無双十三面待ち"
}

func (k *Kokushi13) NumberOfHan() int {
	return k.han
}

func (k *Kokushi13) NumberOfHanWhenNaki() int {
	return k.nakihan
}

func (p *Pinhu) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (p *Pinhu) GetName() string {
	return "平和"
}

func (p *Pinhu) NumberOfHan() int {
	return p.han
}

func (p *Pinhu) NumberOfHanWhenNaki() int {
	return p.nakihan
}

func (h *Haku) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (h *Haku) GetName() string {
	return "白"
}

func (h *Haku) NumberOfHan() int {
	return h.han
}

func (h *Haku) NumberOfHanWhenNaki() int {
	return h.nakihan
}

func (h *Hatsu) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (h *Hatsu) GetName() string {
	return "白"
}

func (h *Hatsu) NumberOfHan() int {
	return h.han
}

func (h *Hatsu) NumberOfHanWhenNaki() int {
	return h.nakihan
}

func (c *Chun) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (c *Chun) GetName() string {
	return "中"
}

func (c *Chun) NumberOfHan() int {
	return c.han
}

func (c *Chun) NumberOfHanWhenNaki() int {
	return c.nakihan
}

func (t *Ton) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (t *Ton) GetName() string {
	return "東"
}

func (t *Ton) NumberOfHan() int {
	return t.han
}

func (t *Ton) NumberOfHanWhenNaki() int {
	return t.nakihan
}

func (n *Nan) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (n *Nan) GetName() string {
	return "南"
}

func (n *Nan) NumberOfHan() int {
	return n.han
}

func (n *Nan) NumberOfHanWhenNaki() int {
	return n.nakihan
}

func (s *Sha) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (s *Sha) GetName() string {
	return "西"
}

func (s *Sha) NumberOfHan() int {
	return s.han
}

func (s *Sha) NumberOfHanWhenNaki() int {
	return s.nakihan
}

func (p *Pe) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (p *Pe) GetName() string {
	return "北"
}

func (p *Pe) NumberOfHan() int {
	return p.han
}

func (p *Pe) NumberOfHanWhenNaki() int {
	return p.nakihan
}

func (t *Toitoi) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (t *Toitoi) GetName() string {
	return "対々和"
}

func (t *Toitoi) NumberOfHan() int {
	return t.han
}

func (t *Toitoi) NumberOfHanWhenNaki() int {
	return t.nakihan
}

func (s *SanshokuDoukou) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (s *SanshokuDoukou) GetName() string {
	return "三色同刻"
}

func (s *SanshokuDoukou) NumberOfHan() int {
	return s.han
}

func (s *SanshokuDoukou) NumberOfHanWhenNaki() int {
	return s.nakihan
}

func (s *SanshokuDoujun) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (s *SanshokuDoujun) GetName() string {
	return "三色同順"
}

func (s *SanshokuDoujun) NumberOfHan() int {
	return s.han
}

func (s *SanshokuDoujun) NumberOfHanWhenNaki() int {
	return s.nakihan
}

func (h *Honroto) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (h *Honroto) GetName() string {
	return "混老頭"
}

func (h *Honroto) NumberOfHan() int {
	return h.han
}

func (h *Honroto) NumberOfHanWhenNaki() int {
	return h.nakihan
}

func (i *Ikkitsuukan) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (i *Ikkitsuukan) GetName() string {
	return "一気通貫"
}

func (i *Ikkitsuukan) NumberOfHan() int {
	return i.han
}

func (i *Ikkitsuukan) NumberOfHanWhenNaki() int {
	return i.nakihan
}

func (c *Chanta) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (c *Chanta) GetName() string {
	return "混全帯么九"
}

func (c *Chanta) NumberOfHan() int {
	return c.han
}

func (c *Chanta) NumberOfHanWhenNaki() int {
	return c.nakihan
}

func (s *Shousangen) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (s *Shousangen) GetName() string {
	return "小三元"
}

func (s *Shousangen) NumberOfHan() int {
	return s.han
}

func (s *Shousangen) NumberOfHanWhenNaki() int {
	return s.nakihan
}

func (h *Honitsu) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (h *Honitsu) GetName() string {
	return "混一色"
}

func (h *Honitsu) NumberOfHan() int {
	return h.han
}

func (h *Honitsu) NumberOfHanWhenNaki() int {
	return h.nakihan
}

func (j *Junchan) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (j *Junchan) GetName() string {
	return "純全帯么九"
}

func (j *Junchan) NumberOfHan() int {
	return j.han
}

func (j *Junchan) NumberOfHanWhenNaki() int {
	return j.nakihan
}

func (c *Chinitsu) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (c *Chinitsu) GetName() string {
	return "清一色"
}

func (c *Chinitsu) NumberOfHan() int {
	return c.han
}

func (c *Chinitsu) NumberOfHanWhenNaki() int {
	return c.nakihan
}

func (r *Ryuiso) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (r *Ryuiso) GetName() string {
	return "緑一色"
}

func (r *Ryuiso) NumberOfHan() int {
	return r.han
}

func (r *Ryuiso) NumberOfHanWhenNaki() int {
	return r.nakihan
}

func (d *Daisangen) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (d *Daisangen) GetName() string {
	return "大三元"
}

func (d *Daisangen) NumberOfHan() int {
	return d.han
}

func (d *Daisangen) NumberOfHanWhenNaki() int {
	return d.nakihan
}

func (s *Shosushi) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (s *Shosushi) GetName() string {
	return "小四喜"
}

func (s *Shosushi) NumberOfHan() int {
	return s.han
}

func (s *Shosushi) NumberOfHanWhenNaki() int {
	return s.nakihan
}

func (t *Tsuiso) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (t *Tsuiso) GetName() string {
	return "字一色"
}

func (t *Tsuiso) NumberOfHan() int {
	return t.han
}

func (t *Tsuiso) NumberOfHanWhenNaki() int {
	return t.nakihan
}

func (k *Kokushi) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (k *Kokushi) GetName() string {
	return "国士無双"
}

func (k *Kokushi) NumberOfHan() int {
	return k.han
}

func (k *Kokushi) NumberOfHanWhenNaki() int {
	return k.nakihan
}

func (s *Suanko) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (s *Suanko) GetName() string {
	return "四暗刻"
}

func (s *Suanko) NumberOfHan() int {
	return s.han
}

func (s *Suanko) NumberOfHanWhenNaki() int {
	return s.nakihan
}

func (c *Chinroto) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (c *Chinroto) GetName() string {
	return "清老頭"
}

func (c *Chinroto) NumberOfHan() int {
	return c.han
}

func (c *Chinroto) NumberOfHanWhenNaki() int {
	return c.nakihan
}

func (s *Sukantsu) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (s *Sukantsu) GetName() string {
	panic("not implemented") // TODO: Implement
	return "四槓子"
}

func (s *Sukantsu) NumberOfHan() int {
	return s.han
}

func (s *Sukantsu) NumberOfHanWhenNaki() int {
	return s.nakihan
}

func (d *Daisushi) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (d *Daisushi) GetName() string {
	return "大四喜"
}

func (d *Daisushi) NumberOfHan() int {
	return d.han
}

func (d *Daisushi) NumberOfHanWhenNaki() int {
	return d.nakihan
}

func (c *Churenpoto) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (c *Churenpoto) GetName() string {
	return "九蓮宝燈"
}

func (c *Churenpoto) NumberOfHan() int {
	return c.han
}

func (c *Churenpoto) NumberOfHanWhenNaki() int {
	return c.nakihan
}

func (r *Ryanpeko) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (r *Ryanpeko) GetName() string {
	return "二盃口"
}

func (r *Ryanpeko) NumberOfHan() int {
	return r.han
}

func (r *Ryanpeko) NumberOfHanWhenNaki() int {
	return r.nakihan
}

func (i *Ipeko) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (i *Ipeko) GetName() string {
	return "一盃口"
}

func (i *Ipeko) NumberOfHan() int {
	return i.han
}

func (i *Ipeko) NumberOfHanWhenNaki() int {
	return i.nakihan
}

func (n *Nagashimangan) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (n *Nagashimangan) GetName() string {
	return "流し満貫"
}

func (n *Nagashimangan) NumberOfHan() int {
	return n.han
}

func (n *Nagashimangan) NumberOfHanWhenNaki() int {
	return n.nakihan
}

func (t *Tenho) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (t *Tenho) GetName() string {
	return "天和"
}

func (t *Tenho) NumberOfHan() int {
	return t.han
}

func (t *Tenho) NumberOfHanWhenNaki() int {
	return t.nakihan
}

func (c *Chiho) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (c *Chiho) GetName() string {
	return "地和"
}

func (c *Chiho) NumberOfHan() int {
	return c.han
}

func (c *Chiho) NumberOfHanWhenNaki() int {
	return c.nakihan
}

func (r *Renho) IsMatch(player *MPlayer, agarikei *CountOfShantenAndAgarikei) bool {
	panic("not implemented") // TODO: Implement
}

func (r *Renho) GetName() string {
	return "人和"
}

func (r *Renho) NumberOfHan() int {
	return r.han
}

func (r *Renho) NumberOfHanWhenNaki() int {
	return r.nakihan
}

func NewTanyao(han, nakihan int) *Tanyao {
	return &Tanyao{
		han:     han,
		nakihan: nakihan,
	}
}
func NewReach(han, nakihan int) *Reach {
	return &Reach{
		han:     han,
		nakihan: nakihan,
	}
}
func NewIppatsu(han, nakihan int) *Ippatsu {
	return &Ippatsu{
		han:     han,
		nakihan: nakihan,
	}
}
func NewMenzenTsumo(han, nakihan int) *MenzenTsumo {
	return &MenzenTsumo{
		han:     han,
		nakihan: nakihan,
	}
}
func NewChankan(han, nakihan int) *Chankan {
	return &Chankan{
		han:     han,
		nakihan: nakihan,
	}
}
func NewRinshan(han, nakihan int) *Rinshan {
	return &Rinshan{
		han:     han,
		nakihan: nakihan,
	}
}
func NewHaitei(han, nakihan int) *Haitei {
	return &Haitei{
		han:     han,
		nakihan: nakihan,
	}
}
func NewHoutei(han, nakihan int) *Houtei {
	return &Houtei{
		han:     han,
		nakihan: nakihan,
	}
}
func NewDoubleReach(han, nakihan int) *DoubleReach {
	return &DoubleReach{
		han:     han,
		nakihan: nakihan,
	}
}
func NewChitoitsu(han, nakihan int) *Chitoitsu {
	return &Chitoitsu{
		han:     han,
		nakihan: nakihan,
	}
}
func NewDabuTon(han, nakihan int) *DabuTon {
	return &DabuTon{
		han:     han,
		nakihan: nakihan,
	}
}
func NewDabuNan(han, nakihan int) *DabuNan {
	return &DabuNan{
		han:     han,
		nakihan: nakihan,
	}
}
func NewDabuSha(han, nakihan int) *DabuSha {
	return &DabuSha{
		han:     han,
		nakihan: nakihan,
	}
}
func NewDabuPe(han, nakihan int) *DabuPe {
	return &DabuPe{
		han:     han,
		nakihan: nakihan,
	}
}
func NewSanAnko(han, nakihan int) *SanAnko {
	return &SanAnko{
		han:     han,
		nakihan: nakihan,
	}
}
func NewSanKantsu(han, nakihan int) *SanKantsu {
	return &SanKantsu{
		han:     han,
		nakihan: nakihan,
	}
}
func NewSuankoTanki(han, nakihan int) *SuankoTanki {
	return &SuankoTanki{
		han:     han,
		nakihan: nakihan,
	}
}
func NewJunseiChuren(han, nakihan int) *JunseiChuren {
	return &JunseiChuren{
		han:     han,
		nakihan: nakihan,
	}
}
func NewKokushi13(han, nakihan int) *Kokushi13 {
	return &Kokushi13{
		han:     han,
		nakihan: nakihan,
	}
}
func NewPinhu(han, nakihan int) *Pinhu {
	return &Pinhu{
		han:     han,
		nakihan: nakihan,
	}
}
func NewHaku(han, nakihan int) *Haku {
	return &Haku{
		han:     han,
		nakihan: nakihan,
	}
}
func NewHatsu(han, nakihan int) *Hatsu {
	return &Hatsu{
		han:     han,
		nakihan: nakihan,
	}
}
func NewChun(han, nakihan int) *Chun {
	return &Chun{
		han:     han,
		nakihan: nakihan,
	}
}
func NewTon(han, nakihan int) *Ton {
	return &Ton{
		han:     han,
		nakihan: nakihan,
	}
}
func NewNan(han, nakihan int) *Nan {
	return &Nan{
		han:     han,
		nakihan: nakihan,
	}
}
func NewSha(han, nakihan int) *Sha {
	return &Sha{
		han:     han,
		nakihan: nakihan,
	}
}
func NewPe(han, nakihan int) *Pe {
	return &Pe{
		han:     han,
		nakihan: nakihan,
	}
}
func NewToitoi(han, nakihan int) *Toitoi {
	return &Toitoi{
		han:     han,
		nakihan: nakihan,
	}
}
func NewSanshokuDoukou(han, nakihan int) *SanshokuDoukou {
	return &SanshokuDoukou{
		han:     han,
		nakihan: nakihan,
	}
}
func NewSanshokuDoujun(han, nakihan int) *SanshokuDoujun {
	return &SanshokuDoujun{
		han:     han,
		nakihan: nakihan,
	}
}
func NewHonroto(han, nakihan int) *Honroto {
	return &Honroto{
		han:     han,
		nakihan: nakihan,
	}
}
func NewIkkitsuukan(han, nakihan int) *Ikkitsuukan {
	return &Ikkitsuukan{
		han:     han,
		nakihan: nakihan,
	}
}
func NewChanta(han, nakihan int) *Chanta {
	return &Chanta{
		han:     han,
		nakihan: nakihan,
	}
}
func NewShousangen(han, nakihan int) *Shousangen {
	return &Shousangen{
		han:     han,
		nakihan: nakihan,
	}
}
func NewHonitsu(han, nakihan int) *Honitsu {
	return &Honitsu{
		han:     han,
		nakihan: nakihan,
	}
}
func NewJunchan(han, nakihan int) *Junchan {
	return &Junchan{
		han:     han,
		nakihan: nakihan,
	}
}
func NewChinitsu(han, nakihan int) *Chinitsu {
	return &Chinitsu{
		han:     han,
		nakihan: nakihan,
	}
}
func NewRyuiso(han, nakihan int) *Ryuiso {
	return &Ryuiso{
		han:     han,
		nakihan: nakihan,
	}
}
func NewDaisangen(han, nakihan int) *Daisangen {
	return &Daisangen{
		han:     han,
		nakihan: nakihan,
	}
}
func NewShosushi(han, nakihan int) *Shosushi {
	return &Shosushi{
		han:     han,
		nakihan: nakihan,
	}
}
func NewTsuiso(han, nakihan int) *Tsuiso {
	return &Tsuiso{
		han:     han,
		nakihan: nakihan,
	}
}
func NewKokushi(han, nakihan int) *Kokushi {
	return &Kokushi{
		han:     han,
		nakihan: nakihan,
	}
}
func NewSuanko(han, nakihan int) *Suanko {
	return &Suanko{
		han:     han,
		nakihan: nakihan,
	}
}
func NewChinroto(han, nakihan int) *Chinroto {
	return &Chinroto{
		han:     han,
		nakihan: nakihan,
	}
}
func NewSukantsu(han, nakihan int) *Sukantsu {
	return &Sukantsu{
		han:     han,
		nakihan: nakihan,
	}
}
func NewDaisushi(han, nakihan int) *Daisushi {
	return &Daisushi{
		han:     han,
		nakihan: nakihan,
	}
}
func NewChurenpoto(han, nakihan int) *Churenpoto {
	return &Churenpoto{
		han:     han,
		nakihan: nakihan,
	}
}
func NewRyanpeko(han, nakihan int) *Ryanpeko {
	return &Ryanpeko{
		han:     han,
		nakihan: nakihan,
	}
}
func NewIpeko(han, nakihan int) *Ipeko {
	return &Ipeko{
		han:     han,
		nakihan: nakihan,
	}
}
func NewNagashimangan(han, nakihan int) *Nagashimangan {
	return &Nagashimangan{
		han:     han,
		nakihan: nakihan,
	}
}
func NewTenho(han, nakihan int) *Tenho {
	return &Tenho{
		han:     han,
		nakihan: nakihan,
	}
}
func NewChiho(han, nakihan int) *Chiho {
	return &Chiho{
		han:     han,
		nakihan: nakihan,
	}
}
func NewRenho(han, nakihan int) *Renho {
	return &Renho{
		han:     han,
		nakihan: nakihan,
	}
}

func GenerateYakusDefault() []Yaku {
	//TODO
	panic("notImplemented")
}

func GenerateYakusNimar() []Yaku {
	//TODO
	panic("notImplemented")
}

// ˄

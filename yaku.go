// ˅
package nimar

// ˄

type Yaku interface {
	IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool

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
	if player.ronTile != nil {
		tileIDs[player.GetRonTile().GetID()]++
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
type KyushuKyuhai struct {
	han     int
	nakihan int
}

func (t *Tanyao) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
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

func (r *Reach) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
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

func (i *Ippatsu) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
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

func (m *MenzenTsumo) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
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

func (c *Chankan) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	return player.status.Chankan
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

func (r *Rinshan) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	return player.status.Rinshan
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

func (h *Haitei) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	return player.status.Haitei
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

func (h *Houtei) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	return player.status.Hotei
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

func (d *DoubleReach) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	return player.status.DoubleReach
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

func (c *Chitoitsu) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	return table.GetGameManager().shantenChecker.checkChitoitsu(player) == -1
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

func (d *DabuTon) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {

		if mentsu.IsEmpty() {
			continue
		}
		if *player.status.Kaze == Kaze_KAZE_TON && *table.GetStatus().Kaze == Kaze_KAZE_TON && mentsu[31] >= 3 {
			return true
		}
	}
	return false
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

func (d *DabuNan) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {

		if mentsu.IsEmpty() {
			continue
		}
		if *player.status.Kaze == Kaze_KAZE_NAN && *table.GetStatus().Kaze == Kaze_KAZE_NAN && mentsu[32] >= 3 {
			return true
		}
	}
	return false
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

func (d *DabuSha) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {

		if mentsu.IsEmpty() {
			continue
		}
		if *player.status.Kaze == Kaze_KAZE_SHA && *table.GetStatus().Kaze == Kaze_KAZE_SHA && mentsu[33] >= 3 {
			return true
		}
	}
	return false
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

func (d *DabuPe) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {

		if mentsu.IsEmpty() {
			continue
		}
		if *player.status.Kaze == Kaze_KAZE_PE && *table.GetStatus().Kaze == Kaze_KAZE_PE && mentsu[34] >= 3 {
			return true
		}
	}
	return false
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

func (s *SanAnko) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	ankoCount := 0
	for _, mentsuType := range []MentsuType{
		agarikei.Agarikei.Mentsu1Type,
		agarikei.Agarikei.Mentsu2Type,
		agarikei.Agarikei.Mentsu3Type,
		agarikei.Agarikei.Mentsu4Type,
	} {
		switch mentsuType {
		case Ankan:
		case Anko:
			ankoCount++
		}
	}
	return ankoCount >= 3
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

func (s *SanKantsu) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	kantsuCount := 0
	for _, mentsuType := range []MentsuType{
		agarikei.Agarikei.Mentsu1Type,
		agarikei.Agarikei.Mentsu2Type,
		agarikei.Agarikei.Mentsu3Type,
		agarikei.Agarikei.Mentsu4Type,
	} {
		switch mentsuType {
		case Ankan:
		case Minkan:
			kantsuCount++
		}
	}
	return kantsuCount >= 3
}

func (s *SanKantsu) GetName() string {
	return "三槓子"
}

func (s *SanKantsu) NumberOfHan() int {
	return s.han
}

func (s *SanKantsu) NumberOfHanWhenNaki() int {
	return s.nakihan
}

func (s *SuankoTanki) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	if agarikei.Agarikei.Janto.IsEmpty() {
		return false
	}
	if player.GetTsumoriTile() != nil && agarikei.Agarikei.Janto[player.GetTsumoriTile().GetID()] != 2 {
		return false
	}
	if player.GetRonTile() != nil && agarikei.Agarikei.Janto[player.GetRonTile().GetID()] != 2 {
		return false
	}

	ankoCount := 0
	for _, mentsuType := range []MentsuType{
		agarikei.Agarikei.Mentsu1Type,
		agarikei.Agarikei.Mentsu2Type,
		agarikei.Agarikei.Mentsu3Type,
		agarikei.Agarikei.Mentsu4Type,
	} {
		switch mentsuType {
		case Ankan:
		case Anko:
			ankoCount++
		}
	}
	return ankoCount >= 4
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

func (j *JunseiChuren) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	isJunseiTyuuren := false
	tileIDs := handAndTsumoriTile(player)
	if player.GetTsumoriTile() != nil {
		tileIDs[player.GetTsumoriTile().GetID()]--
	}
	if player.GetRonTile() != nil {
		tileIDs[player.GetRonTile().GetID()]--
	}

	for i := 0; i < 3; i++ {
		if tileIDs[i*10+1] == 3 && tileIDs[i*10+2] == 1 && tileIDs[i*10+3] == 1 && tileIDs[i*10+4] == 1 && tileIDs[i*10+5] == 1 && tileIDs[i*10+6] == 1 && tileIDs[i*10+7] == 1 && tileIDs[i*10+8] == 1 && tileIDs[i*10+9] == 3 {
			isJunseiTyuuren = true
		}
	}

	if player.GetTsumoriTile() != nil {
		tileIDs[player.GetTsumoriTile().GetID()]++
	}
	if player.GetRonTile() != nil {
		tileIDs[player.GetRonTile().GetID()]++
	}
	return isJunseiTyuuren
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

func (k *Kokushi13) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	isKokusi13 := false
	tileIDs := handAndTsumoriTile(player)
	if player.GetTsumoriTile() != nil {
		tileIDs[player.GetTsumoriTile().GetID()]--
	}
	if player.GetRonTile() != nil {
		tileIDs[player.GetRonTile().GetID()]--
	}

	if tileIDs[1] == 1 && tileIDs[9] == 1 && tileIDs[11] == 1 && tileIDs[19] == 1 && tileIDs[21] == 1 && tileIDs[29] == 1 && tileIDs[31] == 1 && tileIDs[32] == 1 && tileIDs[33] == 1 && tileIDs[34] == 1 && tileIDs[35] == 1 && tileIDs[36] == 1 && tileIDs[37] == 1 {
		isKokusi13 = true
	}

	if player.GetTsumoriTile() != nil {
		tileIDs[player.GetTsumoriTile().GetID()]++
	}
	if player.GetRonTile() != nil {
		tileIDs[player.GetRonTile().GetID()]++
	}

	return isKokusi13
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

func (p *Pinhu) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	tileIDs := handAndTsumoriTile(player)
	if tileIDs[35] != 0 {
		return false
	}
	if tileIDs[36] != 0 {
		return false
	}
	if tileIDs[37] != 0 {
		return false
	}
	switch *player.status.Kaze {
	case Kaze_KAZE_TON:
		if tileIDs[31] != 0 {
			return false
		}
		break
	case Kaze_KAZE_NAN:
		if tileIDs[32] != 0 {
			return false
		}
		break
	case Kaze_KAZE_SHA:
		if tileIDs[33] != 0 {
			return false
		}
		break
	case Kaze_KAZE_PE:
		if tileIDs[34] != 0 {
			return false
		}
		break
	}
	switch *table.GetStatus().Kaze {
	case Kaze_KAZE_TON:
		if tileIDs[31] != 0 {
			return false
		}
		break
	case Kaze_KAZE_NAN:
		if tileIDs[32] != 0 {
			return false
		}
		break
	case Kaze_KAZE_SHA:
		if tileIDs[33] != 0 {
			return false
		}
		break
	case Kaze_KAZE_PE:
		if tileIDs[34] != 0 {
			return false
		}
		break
	}
	for _, mentsuType := range []MentsuType{
		agarikei.Agarikei.Mentsu1Type,
		agarikei.Agarikei.Mentsu2Type,
		agarikei.Agarikei.Mentsu3Type,
		agarikei.Agarikei.Mentsu4Type,
	} {
		if mentsuType != MenzenShuntsu {
			return false
		}
	}

	// 両面待ち
	playerTemp := *player
	agariTileID := 0
	playerTemp.SetTsumoriTile(nil)
	playerTemp.SetRonTile(nil)

	tileIDs[agariTileID] -= 1
	matihai := table.GetGameManager().shantenChecker.CheckCountOfShanten(&playerTemp).Machihai
	tileIDs[agariTileID] += 1
	for i := 0; i <= 2; i++ {
		for j := 1; j <= 6; j++ {
			var small = false
			var big = false
			for mati := range matihai {
				if mati == i*10+j {
					small = true
				}
				if mati == i*10+j+3 {
					big = true
				}
			}
			if big && small {
				return true
			}
		}
	}
	return false
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

func (h *Haku) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		if mentsu.IsEmpty() {
			return false
		}
		if mentsu[35] >= 3 {
			return true
		}
	}
	return false
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

func (h *Hatsu) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		if mentsu.IsEmpty() {
			return false
		}
		if mentsu[36] >= 3 {
			return true
		}
	}
	return false
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

func (c *Chun) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		if mentsu.IsEmpty() {
			return false
		}
		if mentsu[37] >= 3 {
			return true
		}
	}
	return false
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

func (t *Ton) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		if mentsu.IsEmpty() {
			return false
		}
		if *table.tableStatus.Kaze == Kaze_KAZE_TON || *player.status.Kaze == Kaze_KAZE_TON {
			return mentsu[31] >= 3
		}
	}
	return false
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

func (n *Nan) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		if mentsu.IsEmpty() {
			return false
		}
		if *table.tableStatus.Kaze == Kaze_KAZE_NAN || *player.status.Kaze == Kaze_KAZE_NAN {
			return mentsu[32] >= 3
		}
	}
	return false
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

func (s *Sha) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		if mentsu.IsEmpty() {
			return false
		}
		if *table.tableStatus.Kaze == Kaze_KAZE_SHA || *player.status.Kaze == Kaze_KAZE_SHA {
			return mentsu[33] >= 3
		}
	}
	return false
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

func (p *Pe) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		if mentsu.IsEmpty() {
			return false
		}
		if *table.tableStatus.Kaze == Kaze_KAZE_PE || *player.status.Kaze == Kaze_KAZE_PE {
			return mentsu[34] >= 3
		}
	}
	return false
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

func (t *Toitoi) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsuType := range []MentsuType{
		agarikei.Agarikei.Mentsu1Type,
		agarikei.Agarikei.Mentsu2Type,
		agarikei.Agarikei.Mentsu3Type,
		agarikei.Agarikei.Mentsu4Type,
	} {
		switch mentsuType {
		case Anko:
			fallthrough
		case Ankan:
			continue
		default:
			return false
		}
	}
	return true
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

func (s *SanshokuDoukou) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	tileIDs := handAndTsumoriTile(player)
	for i := 1; i <= 9; i++ {
		if tileIDs[i] > 3 && tileIDs[i+10] > 3 && tileIDs[i+20] > 3 {
			return true
		}
	}
	return false
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

func (s *SanshokuDoujun) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	tileIDs := handAndTsumoriTile(player)
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		if mentsu.IsEmpty() {
			return false
		}
		//TODO 索子しか使われないから使われません。このまんまでいいや
		for j := 1; j <= 7; j++ {
			if tileIDs[j] > 1 && tileIDs[j+1] > 1 && tileIDs[j+2] > 1 &&
				tileIDs[j+10] > 1 && tileIDs[j+11] > 1 && tileIDs[j+12] > 1 &&
				tileIDs[j+20] > 1 && tileIDs[j+21] > 1 && tileIDs[j+22] > 1 {
				return true
			}
		}
	}
	return false
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

func (h *Honroto) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	tileIDs := handAndTsumoriTile(player)
	for i := 2; i <= 8; i++ {
		if tileIDs[i] != 0 {
			return false
		}
	}
	for i := 12; i <= 18; i++ {
		if tileIDs[i] != 0 {
			return false
		}
	}
	for i := 22; i <= 28; i++ {
		if tileIDs[i] != 0 {
			return false
		}
	}
	return true
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

func (i *Ikkitsuukan) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for i := 0; i < 3; i++ {
		m123 := false
		m456 := false
		m789 := false

		for _, mentsu := range []TileIDs{
			agarikei.Agarikei.Mentsu1,
			agarikei.Agarikei.Mentsu2,
			agarikei.Agarikei.Mentsu3,
			agarikei.Agarikei.Mentsu4,
		} {
			if mentsu.IsEmpty() {
				return false
			}
			m123 = m123 || mentsu[i*10+1] == 1 && mentsu[i*10+2] == 1 && mentsu[i*10+3] == 1
			m456 = m456 || mentsu[i*10+4] == 1 && mentsu[i*10+5] == 1 && mentsu[i*10+6] == 1
			m789 = m789 || mentsu[i*10+7] == 1 && mentsu[i*10+8] == 1 && mentsu[i*10+9] == 1
		}
		if m123 && m456 && m789 {
			return true
		}
	}
	return false
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

func (c *Chanta) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		var hasYaochu = false
		if mentsu.IsEmpty() {
			return false
		}
		for j := 0; j < len(mentsu); j++ {
			if (j == 1 || j == 9 || j == 11 || j == 19 || j == 21 || j == 29 || j == 31 || j == 32 || j == 33 || j == 34 || j == 35 || j == 36 || j == 37) && mentsu[j] != 0 {
				hasYaochu = true
				break
			}
		}
		if !hasYaochu {
			return false
		}
	}
	return true
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

func (s *Shousangen) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	tileIDs := handAndTsumoriTile(player)
	if tileIDs[35] >= 2 && tileIDs[36] >= 3 && tileIDs[37] >= 3 {
		return true
	}
	if tileIDs[35] >= 3 && tileIDs[36] >= 2 && tileIDs[37] >= 3 {
		return true
	}
	if tileIDs[35] >= 3 && tileIDs[36] >= 3 && tileIDs[37] >= 2 {
		return true
	}
	return false
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

func (h *Honitsu) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	var man = false
	var pin = false
	var sou = false

	tileIDs := handAndTsumoriTile(player)

	for i := 0; i < len(tileIDs); i++ {
		for j := 1; j <= 9; j++ {
			man = man || tileIDs[j] > 0
			pin = pin || tileIDs[j+10] > 0
			sou = sou || tileIDs[j+20] > 0
		}
	}
	return (man && !pin && !sou) || (!man && pin && !sou) || (!man && !pin && sou)
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

func (j *Junchan) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	for _, mentsu := range []TileIDs{
		agarikei.Agarikei.Mentsu1,
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		has19 := false
		if mentsu.IsEmpty() {
			return false
		}
		for j := 0; j < len(mentsu); j++ {
			if (j == 1 || j == 9 || j == 11 || j == 19 || j == 21 || j == 29) && mentsu[j] != 0 {
				has19 = true
				break
			}
		}
		if !has19 {
			return false
		}
	}
	return true
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

func (c *Chinitsu) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	tileIDs := handAndTsumoriTile(player)
	var man = false
	var pin = false
	var sou = false
	for i := 0; i < len(tileIDs); i++ {
		for j := 31; j <= 37; j++ {
			if tileIDs[j] != 0 {
				return false
			}
		}

		for j := 1; j <= 9; j++ {
			man = man || tileIDs[j] != 0
			pin = pin || tileIDs[j+10] != 0
			sou = sou || tileIDs[j+20] != 0
		}

	}
	return (man && !pin && !sou) || (!man && pin && !sou) || (!man && !pin && sou)
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

func (r *Ryuiso) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	tileIDs := handAndTsumoriTile(player)
	var hasHatsu = false
	for j := 0; j < len(tileIDs); j++ {
		if ((j >= 1 && j <= 20) || (j >= 31 && j <= 35) || (j == 37) || (j == 21 || j == 25 || j == 27 || j == 29)) || tileIDs[j] != 0 {
			return false
		}
		if j == 36 && tileIDs[36] != 0 {
			hasHatsu = true
		}
	}
	return hasHatsu
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

func (d *Daisangen) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	tileIDs := handAndTsumoriTile(player)
	return tileIDs[35] >= 3 && tileIDs[36] >= 3 && tileIDs[37] >= 3
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

func (s *Shosushi) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	tileIDs := handAndTsumoriTile(player)

	if tileIDs[31] >= 2 && tileIDs[32] >= 3 && tileIDs[33] >= 3 && tileIDs[34] >= 3 {
		return true
	}
	if tileIDs[31] >= 3 && tileIDs[32] >= 2 && tileIDs[33] >= 3 && tileIDs[34] >= 3 {
		return true
	}
	if tileIDs[31] >= 3 && tileIDs[32] >= 3 && tileIDs[33] >= 2 && tileIDs[34] >= 3 {
		return true
	}
	if tileIDs[31] >= 3 && tileIDs[32] >= 3 && tileIDs[33] >= 3 && tileIDs[34] >= 2 {
		return true
	}
	return false
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

func (t *Tsuiso) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	tileIDs := handAndTsumoriTile(player)
	for j := 0; j < len(tileIDs); j++ {
		if !(j == 31 || j == 32 || j == 33 || j == 34 || j == 35 || j == 36 || j == 37) && tileIDs[j] != 0 {
			return false
		}
	}
	return true
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

func (k *Kokushi) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	return table.GetGameManager().shantenChecker.checkKokushi(player) == -1
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

func (s *Suanko) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	ankoCount := 0
	for _, mentsuType := range []MentsuType{
		agarikei.Agarikei.Mentsu1Type,
		agarikei.Agarikei.Mentsu2Type,
		agarikei.Agarikei.Mentsu3Type,
		agarikei.Agarikei.Mentsu4Type,
	} {
		switch mentsuType {
		case Ankan:
		case Anko:
			ankoCount++
		}
	}
	return ankoCount >= 4
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

func (c *Chinroto) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	tileIDs := handAndTsumoriTile(player)
	for i := 2; i <= 8; i++ {
		if tileIDs[i] > 0 {
			return false
		}
	}
	for i := 12; i <= 18; i++ {
		if tileIDs[i] > 0 {
			return false
		}
	}
	for i := 22; i <= 28; i++ {
		if tileIDs[i] > 0 {
			return false
		}
	}
	for i := 31; i <= 37; i++ {
		if tileIDs[i] > 0 {
			return false
		}
	}
	return true
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

func (s *Sukantsu) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	kantsuCount := 0
	for _, mentsuType := range []MentsuType{
		agarikei.Agarikei.Mentsu1Type,
		agarikei.Agarikei.Mentsu2Type,
		agarikei.Agarikei.Mentsu3Type,
		agarikei.Agarikei.Mentsu4Type,
	} {
		switch mentsuType {
		case Ankan:
		case Minkan:
			kantsuCount++
		}
	}
	return kantsuCount >= 4
}

func (s *Sukantsu) GetName() string {
	return "四槓子"
}

func (s *Sukantsu) NumberOfHan() int {
	return s.han
}

func (s *Sukantsu) NumberOfHanWhenNaki() int {
	return s.nakihan
}

func (d *Daisushi) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	tileIDs := handAndTsumoriTile(player)
	return tileIDs[31] >= 3 && tileIDs[32] >= 3 && tileIDs[33] >= 3 && tileIDs[34] >= 3
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

func (c *Churenpoto) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	isChuren := false
	tileIDs := handAndTsumoriTile(player)

	for i := 0; i < 3; i++ {
		if tileIDs[i*10+1] >= 3 && tileIDs[i*10+2] >= 1 && tileIDs[i*10+3] >= 1 && tileIDs[i*10+4] >= 1 && tileIDs[i*10+5] >= 1 && tileIDs[i*10+6] >= 1 && tileIDs[i*10+7] >= 1 && tileIDs[i*10+8] >= 1 && tileIDs[i*10+9] >= 3 {
			tileIDs[i*10+1] -= 3
			tileIDs[i*10+2] -= 1
			tileIDs[i*10+3] -= 1
			tileIDs[i*10+4] -= 1
			tileIDs[i*10+5] -= 1
			tileIDs[i*10+6] -= 1
			tileIDs[i*10+7] -= 1
			tileIDs[i*10+8] -= 1
			tileIDs[i*10+9] -= 3

			if tileIDs[i*10+1] == 1 && tileIDs[i*10+2] == 1 && tileIDs[i*10+3] == 1 && tileIDs[i*10+4] == 1 && tileIDs[i*10+5] == 1 && tileIDs[i*10+6] == 1 && tileIDs[i*10+7] == 1 && tileIDs[i*10+8] == 1 && tileIDs[i*10+9] == 1 {
				isChuren = true
			}

			tileIDs[i*10+1] += 3
			tileIDs[i*10+2] += 1
			tileIDs[i*10+3] += 1
			tileIDs[i*10+4] += 1
			tileIDs[i*10+5] += 1
			tileIDs[i*10+6] += 1
			tileIDs[i*10+7] += 1
			tileIDs[i*10+8] += 1
			tileIDs[i*10+9] += 3
		}
	}
	return isChuren
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

func (r *Ryanpeko) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	if !player.IsMenzen() {
		return false
	}
	peko := 0

	// 1つ目の面子と同じ面子があるかどうか。234
	for _, hikakuTaisyou := range []TileIDs{
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		if agarikei.Agarikei.Mentsu1.IsEmpty() {
			return false
		}
		mentsu := agarikei.Agarikei.Mentsu1
		for j := 1; j <= 7; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
		for j := 11; j <= 17; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
		for j := 21; j <= 27; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
	}
	// 2つ目の面子と同じ面子があるかどうか。34
	for _, hikakuTaisyou := range []TileIDs{
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		if agarikei.Agarikei.Mentsu2.IsEmpty() {
			return false
		}
		mentsu := agarikei.Agarikei.Mentsu2
		for j := 1; j <= 7; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
		for j := 11; j <= 17; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
		for j := 21; j <= 27; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
	}
	// 3つ目の面子と同じ面子があるかどうか。4。
	for _, hikakuTaisyou := range []TileIDs{
		agarikei.Agarikei.Mentsu4,
	} {
		if agarikei.Agarikei.Mentsu3.IsEmpty() {
			return false
		}
		mentsu := agarikei.Agarikei.Mentsu3
		for j := 1; j <= 7; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
		for j := 11; j <= 17; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
		for j := 21; j <= 27; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
	}
	return peko >= 2
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

func (i *Ipeko) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	if !player.IsMenzen() {
		return false
	}
	peko := 0

	// 1つ目の面子と同じ面子があるかどうか。234
	for _, hikakuTaisyou := range []TileIDs{
		agarikei.Agarikei.Mentsu2,
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		if agarikei.Agarikei.Mentsu1.IsEmpty() {
			return false
		}
		mentsu := agarikei.Agarikei.Mentsu1
		for j := 1; j <= 7; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
		for j := 11; j <= 17; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
		for j := 21; j <= 27; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
	}
	// 2つ目の面子と同じ面子があるかどうか。34
	for _, hikakuTaisyou := range []TileIDs{
		agarikei.Agarikei.Mentsu3,
		agarikei.Agarikei.Mentsu4,
	} {
		if agarikei.Agarikei.Mentsu2.IsEmpty() {
			return false
		}
		mentsu := agarikei.Agarikei.Mentsu2
		for j := 1; j <= 7; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
		for j := 11; j <= 17; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
		for j := 21; j <= 27; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
	}
	// 3つ目の面子と同じ面子があるかどうか。4。
	for _, hikakuTaisyou := range []TileIDs{
		agarikei.Agarikei.Mentsu4,
	} {
		if agarikei.Agarikei.Mentsu3.IsEmpty() {
			return false
		}
		mentsu := agarikei.Agarikei.Mentsu3
		for j := 1; j <= 7; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
		for j := 11; j <= 17; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
		for j := 21; j <= 27; j++ {
			if mentsu[j] == 1 && hikakuTaisyou[j] == 1 && mentsu[j+1] == 1 && hikakuTaisyou[j+1] == 1 && mentsu[j+2] == 1 && hikakuTaisyou[j+2] == 1 {
				peko++
			}
		}
	}
	return peko >= 1
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

func (n *Nagashimangan) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	if player.status.Nakare {
		return false
	}
	for i := 0; i < len(player.GetKawa()); i++ {
		for j := 2; j <= 8; j++ {
			if player.GetKawa()[i].GetID() == j {
				return false
			}
		}
		for j := 12; j <= 18; j++ {
			if player.GetKawa()[i].GetID() == j {
				return false
			}
		}
		for j := 22; j <= 28; j++ {
			if player.GetKawa()[i].GetID() == j {
				return false
			}
		}
	}
	return true
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

func (t *Tenho) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	if len(player.GetKawa()) == 0 && !player.status.Nakare && !player.status.NakareWhenAround && *player.status.Kaze == Kaze_KAZE_TON && player.GetTsumoriTile() != nil && len(player.openedPe.tiles) == 0 {
		return true
	}
	return false
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

func (c *Chiho) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	if len(player.GetKawa()) == 0 && !player.status.Nakare && !player.status.NakareWhenAround && *player.status.Kaze != Kaze_KAZE_TON && player.GetTsumoriTile() != nil && len(player.openedPe.tiles) == 0 {
		return true
	}
	return false
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

func (r *Renho) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	if len(player.GetKawa()) == 0 && !player.status.Nakare && !player.status.NakareWhenAround && *player.status.Kaze != Kaze_KAZE_TON && player.GetRonTile() != nil && len(player.openedPe.tiles) == 0 {
		return true
	}
	return false
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

func (r *KyushuKyuhai) IsMatch(player *MPlayer, table *MTable, agarikei *CountOfShantenAndAgarikei) bool {
	if !(len(player.openedPe.tiles) == 0 && len(player.GetKawa()) == 0) {
		return false
	}
	hasMan := 0
	hasPin := 0
	hasSou := 0
	hasTon := 0
	hasNan := 0
	hasSya := 0
	hasPe := 0
	hasHaku := 0
	hasHatsu := 0
	hasChun := 0
	for _, tileNum := range handAndTsumoriTile(player) {
		if tileNum >= 0 && tileNum <= 9 {
			hasMan = 1
		}
		if tileNum >= 10 && tileNum <= 19 {
			hasPin = 1
		}
		if tileNum >= 20 && tileNum <= 29 {
			hasSou = 1
		}
		if tileNum == 31 {
			hasTon = 1
		}
		if tileNum == 32 {
			hasNan = 1
		}
		if tileNum == 33 {
			hasSya = 1
		}
		if tileNum == 34 {
			hasPe = 1
		}
		if tileNum == 35 {
			hasHaku = 1
		}
		if tileNum == 36 {
			hasHatsu = 1
		}
		if tileNum == 37 {
			hasChun = 1
		}
	}
	return 9 <= hasMan+hasPin+hasSou+hasTon+hasNan+hasSya+hasPe+hasHaku+hasHatsu+hasChun
}

func (r *KyushuKyuhai) GetName() string {
	return "九種九牌"
}

func (r *KyushuKyuhai) NumberOfHan() int {
	return r.han
}

func (r *KyushuKyuhai) NumberOfHanWhenNaki() int {
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
func NewKyushuKyuhai() *KyushuKyuhai {
	return &KyushuKyuhai{
		han:     0,
		nakihan: 0,
	}
}

func GenerateYakusDefault() map[string]Yaku {
	yakus := map[string]Yaku{}
	var yaku Yaku
	yaku = NewTanyao(1, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewReach(1, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewIppatsu(1, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewMenzenTsumo(1, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewChankan(1, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewRinshan(1, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewHaitei(1, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewHoutei(1, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewDoubleReach(2, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewChitoitsu(2, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewDabuTon(2, 2)
	yakus[yaku.GetName()] = yaku
	yaku = NewDabuNan(2, 2)
	yakus[yaku.GetName()] = yaku
	yaku = NewDabuSha(2, 2)
	yakus[yaku.GetName()] = yaku
	yaku = NewDabuPe(2, 2)
	yakus[yaku.GetName()] = yaku
	yaku = NewSanAnko(2, 2)
	yakus[yaku.GetName()] = yaku
	yaku = NewSanKantsu(2, 2)
	yakus[yaku.GetName()] = yaku
	yaku = NewSuankoTanki(26, 26)
	yakus[yaku.GetName()] = yaku
	yaku = NewJunseiChuren(26, 26)
	yakus[yaku.GetName()] = yaku
	yaku = NewKokushi13(26, 26)
	yakus[yaku.GetName()] = yaku
	yaku = NewPinhu(1, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewHaku(1, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewHatsu(1, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewChun(1, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewTon(1, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewNan(1, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewSha(1, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewPe(1, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewToitoi(2, 2)
	yakus[yaku.GetName()] = yaku
	yaku = NewSanshokuDoukou(2, 2)
	yakus[yaku.GetName()] = yaku
	yaku = NewSanshokuDoujun(2, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewHonroto(13, 13)
	yakus[yaku.GetName()] = yaku
	yaku = NewIkkitsuukan(2, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewChanta(2, 1)
	yakus[yaku.GetName()] = yaku
	yaku = NewShousangen(2, 2)
	yakus[yaku.GetName()] = yaku
	yaku = NewHonitsu(3, 2)
	yakus[yaku.GetName()] = yaku
	yaku = NewJunchan(3, 2)
	yakus[yaku.GetName()] = yaku
	yaku = NewRyuiso(13, 13)
	yakus[yaku.GetName()] = yaku
	yaku = NewDaisangen(13, 13)
	yakus[yaku.GetName()] = yaku
	yaku = NewShosushi(13, 13)
	yakus[yaku.GetName()] = yaku
	yaku = NewTsuiso(13, 13)
	yakus[yaku.GetName()] = yaku
	yaku = NewKokushi(13, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewSuanko(13, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewChinroto(13, 13)
	yakus[yaku.GetName()] = yaku
	yaku = NewSukantsu(13, 13)
	yakus[yaku.GetName()] = yaku
	yaku = NewDaisushi(13, 13)
	yakus[yaku.GetName()] = yaku
	yaku = NewChurenpoto(13, 13)
	yakus[yaku.GetName()] = yaku
	yaku = NewRyanpeko(3, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewIpeko(1, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewNagashimangan(4, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewTenho(13, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewChiho(13, 0)
	yakus[yaku.GetName()] = yaku
	yaku = NewRenho(13, 13)
	yakus[yaku.GetName()] = yaku
	yaku = NewKyushuKyuhai()
	yakus[yaku.GetName()] = yaku
	return yakus
}

func GenerateYakusNimar() map[string]Yaku {
	//TODO
	panic("notImplemented")
}

// ˄

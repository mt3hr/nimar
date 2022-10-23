// ˅
package mahjong

import "math"

// ˄

type PointCalcrator struct {
	// ˅

	// ˄

	// ˅

	// ˄
}

func (p *PointCalcrator) CalcratePoint(player *Player, agarikei *CountOfShantenAndAgarikei, table *Table, yakus Yakus) *Point {
	// ˅
	point := &Point{}

	point.MatchYakus = yakus.MatchYakus(player, table, agarikei)

	dora := &Dora{}
	pe := &PeNuki{}
	if dora.IsMatch(player, table, agarikei) {
		yakus[dora.GetName()] = dora
	}
	if pe.IsMatch(player, table, agarikei) {
		yakus[pe.GetName()] = pe
	}

	for _, yaku := range yakus {
		if player.IsMenzen() {
			point.Han += yaku.NumberOfHan()
		} else {
			point.Han += yaku.NumberOfHanWhenNaki()
		}
	}

	if yakus["七対子"].IsMatch(player, table, agarikei) {
		point.Hu = 25
	} else if yakus["平和"].IsMatch(player, table, agarikei) {
		point.Hu = 20
	} else if yakus["国士無双十三面待ち"].IsMatch(player, table, agarikei) || yakus["国士無双"].IsMatch(player, table, agarikei) {
		point.Hu = 0
	} else {
		point.Hu = 20

		if player.IsMenzen() && player.TsumoriTile != nil {
			point.Hu += 2
		} else if player.IsMenzen() && player.RonTile != nil {
			point.Hu += 10
		}

		mentsus := []TileIDs{
			agarikei.Agarikei.Janto,
			agarikei.Agarikei.Mentsu1,
			agarikei.Agarikei.Mentsu2,
			agarikei.Agarikei.Mentsu3,
			agarikei.Agarikei.Mentsu4,
		}
		mentsuTypes := []MentsuType{
			Janto,
			agarikei.Agarikei.Mentsu1Type,
			agarikei.Agarikei.Mentsu2Type,
			agarikei.Agarikei.Mentsu3Type,
			agarikei.Agarikei.Mentsu4Type,
		}
		for i := range mentsus {
			switch mentsuTypes[i] {
			case Janto:

				for tileID, tileCount := range mentsus[i] {
					if tileCount == 2 {
						if tileID == 34 ||
							tileID == 35 ||
							tileID == 36 ||
							((tileID == 31 && *table.Status.Kaze == KAZE_TON) || (tileID == 31 && *player.Status.Kaze == KAZE_TON)) ||
							((tileID == 31 && *table.Status.Kaze == KAZE_NAN) || (tileID == 31 && *player.Status.Kaze == KAZE_NAN)) ||
							((tileID == 31 && *table.Status.Kaze == KAZE_SHA) || (tileID == 31 && *player.Status.Kaze == KAZE_SHA)) ||
							((tileID == 31 && *table.Status.Kaze == KAZE_PE) || (tileID == 31 && *player.Status.Kaze == KAZE_PE)) {
							point.Hu += 2
						}
					}
				}
			case Ankan:
				for tileID, tileCount := range mentsus[i] {
					if tileCount == 4 {
						if tileID == 1 ||
							tileID == 9 ||
							tileID == 11 ||
							tileID == 19 ||
							tileID == 21 ||
							tileID == 29 ||
							tileID == 31 ||
							tileID == 32 ||
							tileID == 33 ||
							tileID == 34 ||
							tileID == 35 ||
							tileID == 36 ||
							tileID == 37 {
							point.Hu += 32
						} else {
							point.Hu += 16
						}
					}
				}
			case Minkan:
				for tileID, tileCount := range mentsus[i] {
					if tileCount == 4 {
						if tileID == 1 ||
							tileID == 9 ||
							tileID == 11 ||
							tileID == 19 ||
							tileID == 21 ||
							tileID == 29 ||
							tileID == 31 ||
							tileID == 32 ||
							tileID == 33 ||
							tileID == 34 ||
							tileID == 35 ||
							tileID == 36 ||
							tileID == 37 {
							point.Hu += 16
						} else {
							point.Hu += 8
						}
					}
				}
			case Anko:
				for tileID, tileCount := range mentsus[i] {
					if tileCount == 3 {
						if tileID == 1 ||
							tileID == 9 ||
							tileID == 11 ||
							tileID == 19 ||
							tileID == 21 ||
							tileID == 29 ||
							tileID == 31 ||
							tileID == 32 ||
							tileID == 33 ||
							tileID == 34 ||
							tileID == 35 ||
							tileID == 36 ||
							tileID == 37 {
							point.Hu += 8
						} else {
							point.Hu += 4
						}
					}
				}
			case Minko:
				for tileID, tileCount := range mentsus[i] {
					if tileCount == 3 {
						if tileID == 1 ||
							tileID == 9 ||
							tileID == 11 ||
							tileID == 19 ||
							tileID == 21 ||
							tileID == 29 ||
							tileID == 31 ||
							tileID == 32 ||
							tileID == 33 ||
							tileID == 34 ||
							tileID == 35 ||
							tileID == 36 ||
							tileID == 37 {
							point.Hu += 4
						} else {
							point.Hu += 2
						}
					}
				}
			case MenzenShuntsu:
			case NakiShuntsu:
			default:
			}
		}
	}

	switch *agarikei.Agarikei.Machi {
	case TANKI:
		fallthrough
	case PENCHAN:
		fallthrough
	case KANCHAN:
		point.Hu += 2
	}

	if point.Hu%10 != 0 {
		var huTemp = 0
		huTemp += ((int)(point.Hu / 100)) * 100
		huTemp += ((int)(point.Hu / 10)) * 10
		huTemp += 10
		point.Hu = huTemp
	}

	// 満願未満
	if point.Han <= 4 {
		switch point.Hu {
		case 20:
		case 25:
		case 30:
			break
		case 40:
		case 50:
		case 60:
			if point.Han >= 4 {
				if player.ID == table.GameManager.dealerPlayer.ID {
					point.Point = 12000 + table.Status.NumberOfHonba*300
				} else {
					point.Point = 8000 + table.Status.NumberOfHonba*300
				}
			}
			break
		case 70:
		case 80:
		case 90:
		case 100:
		case 110:
			if point.Han >= 3 {
				if player.ID == table.GameManager.dealerPlayer.ID {
					point.Point = 12000 + table.Status.NumberOfHonba*300
				} else {
					point.Point = 8000 + table.Status.NumberOfHonba*300
				}
			}
			break
		}
		kihonten := point.Hu * int(math.Pow(float64(2), float64(point.Han+2)))
		tensuu := 0
		if player.ID == table.GameManager.dealerPlayer.ID {
			tensuu = kihonten * 6
		} else {
			tensuu = kihonten * 4
		}

		var fixedTensuu = 0
		fixedTensuu += ((int)(tensuu / 100)) * 100
		if ((int)(tensuu/10))*10 != 0 {
			fixedTensuu += 100
		}
		point.Point = fixedTensuu + table.Status.NumberOfHonba*300
	}

	// 満願以上
	if point.Han == 5 {
		if player.ID == table.GameManager.dealerPlayer.ID {
			point.Point = 12000 + table.Status.NumberOfHonba*300
		} else {
			point.Point = 8000 + table.Status.NumberOfHonba*300
		}
	}
	if point.Han >= 6 && point.Han <= 7 {
		if player.ID == table.GameManager.dealerPlayer.ID {
			point.Point = 18000 + table.Status.NumberOfHonba*300
		} else {
			point.Point = 12000 + table.Status.NumberOfHonba*300
		}
	}
	if point.Han >= 8 && point.Han <= 10 {
		if player.ID == table.GameManager.dealerPlayer.ID {
			point.Point = 24000 + table.Status.NumberOfHonba*300
		} else {
			point.Point = 16000 + table.Status.NumberOfHonba*300
		}
	}
	if point.Han >= 11 && point.Han <= 12 {
		if player.ID == table.GameManager.dealerPlayer.ID {
			point.Point = 36000 + table.Status.NumberOfHonba*300
		} else {
			point.Point = 24000 + table.Status.NumberOfHonba*300
		}
	}
	if player.ID == table.GameManager.dealerPlayer.ID {
		point.Point = point.Han/13*48000 + table.Status.NumberOfHonba*300
	} else {
		point.Point = point.Han/13*32000 + table.Status.NumberOfHonba*300
	}

	return point
	// ˄
}

// ˅

// ˄

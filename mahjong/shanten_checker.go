// ˅
package mahjong

// ˄

type ShantenChecker struct {
	// ˅

	// ˄

	yakuList map[string]Yaku

	// ˅
	countOfToitsu        int // トイツ数
	countOfKotsu         int // コーツ数
	countOfShuntsu       int // シュンツ数
	countOfTatsu         int // ターツ数
	countOfMentsu        int // メンツ数
	countOfAnkan         int // 暗槓数
	countOfMinkan        int // 明槓数
	shantenTemp          int // シャンテン数（計算用）
	shantenNormal        int // シャンテン数（結果用）
	countOfKanzenKotsu   int // 完全コーツ数
	countOfKanzenShuntsu int // 完全シュンツ数
	countOfKanzenKoritsu int // 完全孤立牌数

	// 開かれていない牌ここから
	menzenTileIDs     TileIDs
	tempMenzenTileIDs TileIDs
	// kanzenKoritsu     TileIDs
	// 開かれていない牌ここまで
	// 開かれている牌ここから
	ankanTileIDs       []TileIDs
	minkanTileIDs      []TileIDs
	minkoTileIDs       []TileIDs
	nakishuntsuTileIDs []TileIDs
	// 開かれている牌ここまで

	agarikei Agarikei
	machihai map[int]interface{}
	// ˄
}

func (s *ShantenChecker) CheckCountOfShanten(player *Player) *CountOfShantenAndAgarikei {
	// ˅
	s.machihai = map[int]interface{}{}
	shanten := 8
	shantenTemp := 8

	shantenAndAgarikei := &CountOfShantenAndAgarikei{
		Shanten:  shanten,
		Agarikei: &Agarikei{},
	}

	shantenTemp = s.checkNormal(player)
	if shantenTemp < shanten {
		shanten = shantenTemp
		shantenAndAgarikei.Shanten = shantenTemp

		if shanten <= 0 {
			shantenAndAgarikei.Agarikei.MachiHai = s.calcMachihai()
			shantenAndAgarikei.Agarikei = &Agarikei{}
			if shanten == -1 {
				machi := s.checkMachi(player, &s.agarikei)
				shantenAndAgarikei.Agarikei.Machi = &machi
			}
		}
	}

	shantenTemp = s.checkKokushi(player)
	if shantenTemp < shanten {
		shanten = shantenTemp
		shantenAndAgarikei.Shanten = shantenTemp

		if shanten <= 0 {
			shantenAndAgarikei.Agarikei.MachiHai = s.calcMachihai()
			shantenAndAgarikei.Agarikei = &Agarikei{}
			if shanten == -1 {
				machi := s.checkMachi(player, &s.agarikei)
				shantenAndAgarikei.Agarikei.Machi = &machi
			}
		}
	}

	shantenTemp = s.checkChitoitsu(player)
	if shantenTemp < shanten {
		shanten = shantenTemp
		shantenAndAgarikei.Shanten = shantenTemp

		if shanten <= 0 {
			shantenAndAgarikei.Agarikei.MachiHai = s.calcMachihai()
			shantenAndAgarikei.Agarikei = &Agarikei{}
			if shanten == -1 {
				machi := s.checkMachi(player, &s.agarikei)
				shantenAndAgarikei.Agarikei.Machi = &machi
			}
		}
	}

	return shantenAndAgarikei
	// ˄
}

func (s *ShantenChecker) GetYakuList() Yakus {
	// ˅
	return s.yakuList
	// ˄
}

func (s *ShantenChecker) SetYakuList(yakuList map[string]Yaku) {
	// ˅
	s.yakuList = yakuList
	// ˄
}

func (s *ShantenChecker) checkMachihai(player *Player) map[int]interface{} {
	// ˅

	h := 0
	i := 0
	j := 0
	k := 0
	l := 0
	m := 0
	n := 0
	// 通常の形
	for h = 1; h <= 37; h++ {
		// 対子があれば取り出して
		if s.tempMenzenTileIDs[h] >= 2 {
			// 刻子刻子刻子 余り
			s.tempMenzenTileIDs[h] -= 2
			for i = 1; i <= 37; i++ {
				if s.cutKotsu(i) {
					for j = 1; j <= 37; j++ {
						if s.cutKotsu(j) {
							for k = 1; k <= 37; k++ {
								if s.cutKotsu(k) {
									for machihaiID := range s.calcMachihai() {
										s.machihai[machihaiID] = struct{}{}
									}
									s.addKotsu(k)
								}
							}
							s.addKotsu(j)
						}
					}
					s.addKotsu(i)
				}
			}

			// 順子刻子刻子 余り
			for i = 1; i <= 27; i++ {
				if !((i >= 1 && i <= 7) || (i >= 11 && i <= 17) || (i >= 21 && i <= 27)) {
					continue
				}
				if s.cutShuntsu(i) {
					for j = 1; j <= 37; j++ {
						if s.cutKotsu(j) {
							for k = 1; k <= 37; k++ {
								if s.cutKotsu(k) {
									for machihaiID := range s.calcMachihai() {
										s.machihai[machihaiID] = struct{}{}
									}
									s.addKotsu(k)
								}
							}
							s.addKotsu(j)
						}
					}
					s.addShuntsu(i)
				}
			}

			// 順子順子刻子 余り
			for i = 1; i <= 27; i++ {
				if !((i >= 1 && i <= 7) || (i >= 11 && i <= 17) || (i >= 21 && i <= 27)) {
					continue
				}
				if s.cutShuntsu(i) {
					for j = 1; j <= 27; j++ {
						if !((j >= 1 && j <= 7) || (j >= 11 && j <= 17) || (j >= 21 && j <= 27)) {
							continue
						}
						if s.cutShuntsu(j) {
							for k = 1; k <= 37; k++ {
								if s.cutKotsu(k) {
									for machihaiID := range s.calcMachihai() {
										s.machihai[machihaiID] = struct{}{}
									}
									s.addKotsu(k)
								}
							}
							s.addShuntsu(j)
						}
					}
					s.addShuntsu(i)
				}
			}

			// 順子順子順子 余り
			for i = 1; i <= 27; i++ {
				if !((i >= 1 && i <= 7) || (i >= 11 && i <= 17) || (i >= 21 && i <= 27)) {
					continue
				}
				if s.cutShuntsu(i) {
					for j = 1; j <= 27; j++ {
						if !((j >= 1 && j <= 7) || (j >= 11 && j <= 17) || (j >= 21 && j <= 27)) {
							continue
						}
						if s.cutShuntsu(j) {
							for k = 1; k <= 27; k++ {
								if !((k >= 1 && k <= 7) || (k >= 11 && k <= 17) || (k >= 21 && k <= 27)) {
									continue
								}
								if s.cutShuntsu(k) {
									for machihaiID := range s.calcMachihai() {
										s.machihai[machihaiID] = struct{}{}
									}
									s.addShuntsu(k)
								}
							}
							s.addShuntsu(j)
						}
					}
					s.addShuntsu(i)
				}
			}

			// 七対子
			for i = 1; i <= 37; i++ {
				if s.cutToitsu(i) {
					for j = 1; j <= 37; j++ {
						if s.cutToitsu(j) {
							for k = 1; k <= 37; k++ {
								if s.cutToitsu(k) {
									for l = 1; l <= 37; l++ {
										if s.cutToitsu(l) {
											for m = 1; m <= 37; m++ {
												if s.cutToitsu(m) {
													// 残りの牌で単騎待ち.
													for n = 1; n <= 37; n++ {
														if s.tempMenzenTileIDs[n] == 1 {
															s.machihai[n] = struct{}{}
														}
													}
													s.addToitsu(m)
												}
											}
											s.addToitsu(l)
										}
									}
									s.addToitsu(k)
								}
							}
							s.addToitsu(j)
						}
					}
					s.addToitsu(i)
				}
			}
			s.tempMenzenTileIDs[h] += 2
		}
	}

	// 国士
	if s.checkKokushi(player) == 0 {
		var jantou = -1

		if s.tempMenzenTileIDs[1] == 2 {
			jantou = 1
		}
		if s.tempMenzenTileIDs[9] == 2 {
			jantou = 2
		}
		if s.tempMenzenTileIDs[11] == 2 {
			jantou = 11
		}
		if s.tempMenzenTileIDs[19] == 2 {
			jantou = 19
		}
		if s.tempMenzenTileIDs[21] == 2 {
			jantou = 21
		}
		if s.tempMenzenTileIDs[29] == 2 {
			jantou = 29
		}
		if s.tempMenzenTileIDs[31] == 2 {
			jantou = 31
		}
		if s.tempMenzenTileIDs[32] == 2 {
			jantou = 32
		}
		if s.tempMenzenTileIDs[33] == 2 {
			jantou = 33
		}
		if s.tempMenzenTileIDs[34] == 2 {
			jantou = 34
		}
		if s.tempMenzenTileIDs[35] == 2 {
			jantou = 35
		}
		if s.tempMenzenTileIDs[36] == 2 {
			jantou = 36
		}
		if s.tempMenzenTileIDs[37] == 2 {
			jantou = 37
		}

		if jantou == -1 {
			s.machihai[1] = struct{}{}
			s.machihai[9] = struct{}{}
			s.machihai[11] = struct{}{}
			s.machihai[19] = struct{}{}
			s.machihai[21] = struct{}{}
			s.machihai[29] = struct{}{}
			s.machihai[31] = struct{}{}
			s.machihai[32] = struct{}{}
			s.machihai[33] = struct{}{}
			s.machihai[34] = struct{}{}
			s.machihai[35] = struct{}{}
			s.machihai[38] = struct{}{}
			s.machihai[37] = struct{}{}
		} else {
			if s.tempMenzenTileIDs[1] == 0 {
				s.machihai[1] = struct{}{}
			}
			if s.tempMenzenTileIDs[9] == 0 {
				s.machihai[9] = struct{}{}
			}
			if s.tempMenzenTileIDs[11] == 0 {
				s.machihai[11] = struct{}{}
			}
			if s.tempMenzenTileIDs[19] == 0 {
				s.machihai[19] = struct{}{}
			}
			if s.tempMenzenTileIDs[21] == 0 {
				s.machihai[21] = struct{}{}
			}
			if s.tempMenzenTileIDs[29] == 0 {
				s.machihai[29] = struct{}{}
			}
			if s.tempMenzenTileIDs[31] == 0 {
				s.machihai[31] = struct{}{}
			}
			if s.tempMenzenTileIDs[32] == 0 {
				s.machihai[32] = struct{}{}
			}
			if s.tempMenzenTileIDs[33] == 0 {
				s.machihai[33] = struct{}{}
			}
			if s.tempMenzenTileIDs[34] == 0 {
				s.machihai[34] = struct{}{}
			}
			if s.tempMenzenTileIDs[35] == 0 {
				s.machihai[35] = struct{}{}
			}
			if s.tempMenzenTileIDs[36] == 0 {
				s.machihai[36] = struct{}{}
			}
			if s.tempMenzenTileIDs[37] == 0 {
				s.machihai[37] = struct{}{}
			}
		}
	}
	return s.machihai

	// ˄
}

func (s *ShantenChecker) checkNormal(player *Player) int {
	// ˅
	s.preparation(player)
	i := 0

	// 完全な順子、刻子、孤立牌を抜き、数を数える
	s.countOfKanzenKotsu = s.cutKanzenKotsuAndGetCount()
	s.countOfKanzenShuntsu = s.cutKanzenShuntsuAndGetCount()
	s.countOfKanzenKoritsu = s.cutKanzenKoritsuAndGetCount()

	resetAgarikei := func() {
		s.agarikei.Janto.Reset()
		s.agarikei.Mentsu1.Reset()
		s.agarikei.Mentsu2.Reset()
		s.agarikei.Mentsu3.Reset()
		s.agarikei.Mentsu4.Reset()
	}

	//雀頭抜き出し→コーツ抜き出し→シュンツ抜き出し→ターツ候補抜き出し
	for i = 1; i < 38; i++ {
		shantenTemp := func() int {
			tempMenzenTileIDs := s.tempMenzenTileIDs
			resetTempMenzenTile := func() {
				s.tempMenzenTileIDs = tempMenzenTileIDs
			}

			defer resetAgarikei()
			defer resetTempMenzenTile()

			s.shantenTemp = 8
			s.countOfToitsu = 0
			if s.tempMenzenTileIDs[i] >= 2 {
				s.agarikei.Janto[i] = 2
				s.countOfToitsu++
				s.tempMenzenTileIDs[i] -= 2
			}
			s.mentsuCut(1)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return -1
			}
			s.tempMenzenTileIDs[i] += 2
			s.countOfToitsu--
			return s.shantenTemp
		}()

		if shantenTemp < s.shantenTemp {
			s.shantenTemp = shantenTemp
		}

		if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
			return -1
		}
	}

	if s.agarikei.Janto.IsEmpty() || s.agarikei.Mentsu4.IsEmpty() {
		shantenTemp := func() int {
			tempMenzenTileIDs := s.tempMenzenTileIDs
			resetTempMenzenTile := func() {
				s.tempMenzenTileIDs = tempMenzenTileIDs
			}

			defer resetTempMenzenTile()
			defer resetAgarikei()
			s.countOfToitsu = 0
			s.shantenTemp = 8
			//【雀頭が無い場合の処理】コーツ抜き出し→シュンツ抜き出し→ターツ候補抜き出し
			s.mentsuCut(1)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return -1
			}
			return s.shantenTemp
		}()
		if shantenTemp < s.shantenTemp {
			s.shantenTemp = shantenTemp
		}
	}

	if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
		return -1
	}

	if s.shantenTemp < s.shantenNormal {
		s.shantenNormal = s.shantenTemp
	}

	return s.shantenNormal
	// ˄
}

func (s *ShantenChecker) checkKokushi(player *Player) int {
	// ˅
	s.preparation(player)
	shantenKokusi := 13
	toitsu_suu := 0 //雀頭
	i := 0

	//19牌をチェックする処理
	for i = 1; i < 30; i++ {
		//10で割った余りが1または9の場合に実行する
		if i%10 == 1 || i%10 == 9 {
			if s.menzenTileIDs[i] != 0 {
				shantenKokusi--
			}
			////余った19牌を雀頭としてカウント。1個でOK
			if s.menzenTileIDs[i] >= 2 && toitsu_suu == 0 {
				toitsu_suu = 1
			}
		}
	}

	//字牌をチェックする処理
	for i = 31; i < 38; i++ {
		if s.menzenTileIDs[i] != 0 {
			shantenKokusi--
		}
		////余った字牌を雀頭としてカウント。1個でOK
		if s.menzenTileIDs[i] >= 2 && toitsu_suu == 0 {
			toitsu_suu = 1
		}
	}

	//雀頭がある場合の処理
	shantenKokusi -= toitsu_suu
	return shantenKokusi
	// ˄
}

func (s *ShantenChecker) checkChitoitsu(player *Player) int {
	// ˅
	s.preparation(player)
	countOfChitoitsu := 0          //対子数
	countOfShuruiForChitoitsu := 0 //牌の種類
	shantenChitoitsu := 6          //七対子のシャンテン数
	i := 0
	countOfKantsuForChitoitsu := 0

	for i = 1; i < 38; i++ {
		if s.tempMenzenTileIDs[i] == 0 {
			continue
		} //牌が無い時は以降の処理を中断して、ループの最初に戻る
		if s.tempMenzenTileIDs[i] == 4 {
			countOfKantsuForChitoitsu++
		}
		countOfShuruiForChitoitsu++ //4枚チートイツを回避するために牌種をカウントしておく
		if s.tempMenzenTileIDs[i] >= 2 {
			countOfChitoitsu++
		}
	}

	if countOfShuruiForChitoitsu == 7 && countOfChitoitsu == 7 && countOfKantsuForChitoitsu == 0 {
		return -1 //アガリ判定
	}

	if countOfShuruiForChitoitsu >= 7 && countOfChitoitsu == 6 && countOfKantsuForChitoitsu == 0 {
		return 0 //テンパイ判定
	}

	if countOfShuruiForChitoitsu == 6 && countOfChitoitsu == 6 {
		return 1 //1シャンテン判定
	}

	shantenChitoitsu = 6 - countOfChitoitsu //チートイツのシャンテン数を求める計算式

	return shantenChitoitsu

	// ˄
}

func (s *ShantenChecker) checkMachi(player *Player, agarikei *Agarikei) Machi {
	// ˅
	i := 0
	agarihaiID := 0

	deletedAgarihai := false
	for _, tileIDs := range []TileIDs{
		agarikei.Janto,
		agarikei.Mentsu1,
		agarikei.Mentsu2,
		agarikei.Mentsu3,
		agarikei.Mentsu4,
	} {
		for tileID, tileCount := range tileIDs {
			if !deletedAgarihai && (player.TsumoriTile != nil && player.TsumoriTile.ID == tileID && tileCount > 1) {
				tileIDs[tileID]--
				agarihaiID = tileID
				deletedAgarihai = true
			}

			if !deletedAgarihai && (player.RonTile != nil && player.RonTile.ID == tileID && tileCount > 1) {
				tileIDs[tileID]--
				agarihaiID = tileID
				deletedAgarihai = true
			}
		}
		if deletedAgarihai {
			// 単騎待ち
			for i = 1; i < 38; i++ {
				if tileIDs[i] == 2 {
					if i == agarihaiID {
						return TANKI
					}
				}
			}

			// 辺張待ち
			for i = 0; i <= 2; i++ {
				if tileIDs[i*10+1] == 1 && tileIDs[i*10+2] == 1 {
					if i == agarihaiID {
						return PENCHAN
					}
				}
				if tileIDs[i*10+8] == 1 && tileIDs[i*10+9] == 1 {
					if i == agarihaiID {
						return PENCHAN
					}
				}
			}
			// 両面待ち
			for i = 0; i <= 29; i++ {
				if !((i >= 2 && i <= 7) || (i >= 12 && i <= 17) || (i >= 22 && i <= 27)) {
					continue
				}
				if tileIDs[i] == 1 && tileIDs[i+1] == 1 {
					if i-1 == agarihaiID || i+2 == agarihaiID {
						return RYANMEN
					}
				}
			}
			// 嵌張待ち
			for i = 0; i <= 29; i++ {
				if !((i >= 2 && i <= 7) || (i >= 12 && i <= 17) || (i >= 22 && i <= 27)) {
					continue
				}
				if tileIDs[i] == 1 && tileIDs[i+2] == 1 {
					return KANCHAN
				}
			}
			break
		}
	}
	return TANKI
	// ˄
}

func (s *ShantenChecker) mentsuCut(i int) {
	// ˅
	j := 0
	if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
		s.shantenTemp = -1
		s.shantenNormal = -1
		return
	}
	//※字牌のコーツは完全コーツ処理で抜いているの数牌だけで良い
	for j = i; j < 30; j++ {
		//コーツ抜き出し

		if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
			s.shantenTemp = -1
			s.shantenNormal = -1
			return
		}
		if s.tempMenzenTileIDs[j] >= 3 {
			s.countOfMentsu++
			s.countOfKotsu++
			mentsu := TileIDs{}
			mentsu[j] = 3
			s.addMentsu(mentsu, Anko)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				s.shantenTemp = -1
				s.shantenNormal = -1
				return
			}
			s.tempMenzenTileIDs[j] -= 3
			s.mentsuCut(j)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				s.shantenTemp = -1
				s.shantenNormal = -1
				return
			}
			s.tempMenzenTileIDs[j] += 3
			s.undoMentsu()
			s.countOfKotsu--
			s.countOfMentsu--
		}
		//シュンツ抜き出し
		if s.tempMenzenTileIDs[j] != 0 && s.tempMenzenTileIDs[j+1] != 0 && s.tempMenzenTileIDs[j+2] != 0 && j < 28 {
			s.countOfMentsu++
			s.countOfShuntsu++
			mentsu := TileIDs{}
			mentsu[j] = 1
			mentsu[j+1] = 1
			mentsu[j+2] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				s.shantenTemp = -1
				s.shantenNormal = -1
				return
			}
			s.tempMenzenTileIDs[j]--
			s.tempMenzenTileIDs[j+1]--
			s.tempMenzenTileIDs[j+2]--

			s.mentsuCut(j)

			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				s.shantenTemp = -1
				s.shantenNormal = -1
				return
			}
			s.tempMenzenTileIDs[j]++
			s.tempMenzenTileIDs[j+1]++
			s.tempMenzenTileIDs[j+2]++
			s.undoMentsu()
			s.countOfShuntsu--
			s.countOfMentsu--
		}
	}
	s.tatsuCut(1) //ターツ抜きへ
	return
	// ˄
}

func (s *ShantenChecker) tatsuCut(i int) {
	// ˅
	j := 0

	for j = i; j < 38; j++ {
		s.countOfMentsu = s.countOfAnkan + s.countOfMinkan + s.countOfKanzenKotsu + s.countOfKotsu + s.countOfKanzenShuntsu + s.countOfShuntsu

		if s.countOfMentsu+s.countOfTatsu < 4 {
			//メンツとターツの合計は4まで
			//トイツ抜き出し
			if s.tempMenzenTileIDs[j] == 2 {
				s.countOfTatsu++
				s.tempMenzenTileIDs[j] -= 2
				s.tatsuCut(j)
				s.tempMenzenTileIDs[j] += 2
				s.countOfTatsu--
			}

			//リャンメン・ペンチャン抜き出し
			if j < 29 && j%10 < 9 {
				if s.tempMenzenTileIDs[j] != 0 && s.tempMenzenTileIDs[j+1] != 0 {
					s.countOfTatsu++
					s.tempMenzenTileIDs[j]--
					s.tempMenzenTileIDs[j+1]--
					s.tatsuCut(j)
					s.tempMenzenTileIDs[j]++
					s.tempMenzenTileIDs[j+1]++
					s.countOfTatsu--
				}
			}

			//カンチャン抜き出し
			if j < 28 && j%10 < 8 {
				if s.tempMenzenTileIDs[j] != 0 && s.tempMenzenTileIDs[j+1] == 0 && s.tempMenzenTileIDs[j+2] != 0 {
					s.countOfTatsu++
					s.tempMenzenTileIDs[j]--
					s.tempMenzenTileIDs[j+2]--
					s.tatsuCut(j)
					s.tempMenzenTileIDs[j]++
					s.tempMenzenTileIDs[j+2]++
					s.countOfTatsu--
				}
			}
		}
	}

	s.shantenTemp = 8 - s.countOfMentsu*2 - s.countOfTatsu - s.countOfToitsu
	if s.shantenTemp <= s.shantenNormal {
		s.shantenNormal = s.shantenTemp
	}
	return
	// ˄
}

func (s *ShantenChecker) kantsuCut() int {
	// ˅
	countOfKanzenKantsu := 0
	i := 0
	j := 0
	//字牌の完全コーツを抜き出す
	for i = 31; i < 38; i++ {
		if s.tempMenzenTileIDs[i] >= 4 {
			s.tempMenzenTileIDs[i] -= 4
			countOfKanzenKantsu++

			mentsu := TileIDs{}
			mentsu[i] = 4
			s.addMentsu(mentsu, Ankan)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenKantsu
			}
		}
	}

	//数牌の完全コーツを抜き出す
	for i = 0; i < 30; i += 10 {
		if s.tempMenzenTileIDs[i+1] >= 4 && s.tempMenzenTileIDs[i+2] == 0 && s.tempMenzenTileIDs[i+3] == 0 {
			s.tempMenzenTileIDs[i+1] -= 4
			countOfKanzenKantsu++

			mentsu := TileIDs{}
			mentsu[i+1] = 4
			s.addMentsu(mentsu, Ankan)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenKantsu
			}
		}
		if s.tempMenzenTileIDs[i+1] == 0 && s.tempMenzenTileIDs[i+2] >= 4 && s.tempMenzenTileIDs[i+3] == 0 && s.tempMenzenTileIDs[i+4] == 0 {
			s.tempMenzenTileIDs[i+2] -= 4
			countOfKanzenKantsu++

			mentsu := TileIDs{}
			mentsu[i+2] = 4
			s.addMentsu(mentsu, Ankan)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenKantsu
			}
		}

		//3~7の完全コーツを抜く
		for j = 0; j < 5; j++ {
			if s.tempMenzenTileIDs[i+j+1] == 0 && s.tempMenzenTileIDs[i+j+2] == 0 && s.tempMenzenTileIDs[i+j+3] >= 4 && s.tempMenzenTileIDs[i+j+4] == 0 && s.tempMenzenTileIDs[i+j+5] == 0 {
				s.tempMenzenTileIDs[i+j+3] -= 4
				countOfKanzenKantsu++

				mentsu := TileIDs{}
				mentsu[i+j+3] = 4
				s.addMentsu(mentsu, Ankan)
				if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
					return countOfKanzenKantsu
				}
			}
		}
		if s.tempMenzenTileIDs[i+6] == 0 && s.tempMenzenTileIDs[i+7] == 0 && s.tempMenzenTileIDs[i+8] >= 4 && s.tempMenzenTileIDs[i+9] == 0 {
			s.tempMenzenTileIDs[i+8] -= 4
			countOfKanzenKantsu++

			mentsu := TileIDs{}
			mentsu[i+8] = 4
			s.addMentsu(mentsu, Ankan)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenKantsu
			}
		}

		if s.tempMenzenTileIDs[i+7] == 0 && s.tempMenzenTileIDs[i+8] == 0 && s.tempMenzenTileIDs[i+9] >= 4 {
			s.tempMenzenTileIDs[i+9] -= 4
			countOfKanzenKantsu++

			mentsu := TileIDs{}
			mentsu[i+9] = 4
			s.addMentsu(mentsu, Ankan)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenKantsu
			}
		}
	}
	return countOfKanzenKantsu
	// ˄
}

func (s *ShantenChecker) cutKanzenKotsuAndGetCount() int {
	// ˅
	countOfKanzenKotsu := 0
	i := 0
	j := 0
	//字牌の完全コーツを抜き出す
	for i = 31; i < 38; i++ {
		if s.tempMenzenTileIDs[i] >= 3 {
			s.tempMenzenTileIDs[i] -= 3
			countOfKanzenKotsu++

			mentsu := TileIDs{}
			mentsu[i] = 3
			s.addMentsu(mentsu, Anko)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenKotsu
			}
		}
	}

	//数牌の完全コーツを抜き出す
	for i = 0; i < 30; i += 10 {
		if s.tempMenzenTileIDs[i+1] >= 3 && s.tempMenzenTileIDs[i+2] == 0 && s.tempMenzenTileIDs[i+3] == 0 {
			s.tempMenzenTileIDs[i+1] -= 3
			countOfKanzenKotsu++

			mentsu := TileIDs{}
			mentsu[i+1] = 3
			s.addMentsu(mentsu, Anko)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenKotsu
			}
		}
		if s.tempMenzenTileIDs[i+1] == 0 && s.tempMenzenTileIDs[i+2] >= 3 && s.tempMenzenTileIDs[i+3] == 0 && s.tempMenzenTileIDs[i+4] == 0 {
			s.tempMenzenTileIDs[i+2] -= 3
			countOfKanzenKotsu++

			mentsu := TileIDs{}
			mentsu[i+2] = 3
			s.addMentsu(mentsu, Anko)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenKotsu
			}
		}

		//3~7の完全コーツを抜く
		for j = 0; j < 5; j++ {
			if s.tempMenzenTileIDs[i+j+1] == 0 && s.tempMenzenTileIDs[i+j+2] == 0 && s.tempMenzenTileIDs[i+j+3] >= 3 && s.tempMenzenTileIDs[i+j+4] == 0 && s.tempMenzenTileIDs[i+j+5] == 0 {
				s.tempMenzenTileIDs[i+j+3] -= 3
				countOfKanzenKotsu++

				mentsu := TileIDs{}
				mentsu[i+j+3] = 3
				s.addMentsu(mentsu, Anko)
				if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
					return countOfKanzenKotsu
				}
			}

		}
		if s.tempMenzenTileIDs[i+6] == 0 && s.tempMenzenTileIDs[i+7] == 0 && s.tempMenzenTileIDs[i+8] >= 3 && s.tempMenzenTileIDs[i+9] == 0 {
			s.tempMenzenTileIDs[i+8] -= 3
			countOfKanzenKotsu++

			mentsu := TileIDs{}
			mentsu[i+8] = 3
			s.addMentsu(mentsu, Anko)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenKotsu
			}
		}

		if s.tempMenzenTileIDs[i+7] == 0 && s.tempMenzenTileIDs[i+8] == 0 && s.tempMenzenTileIDs[i+9] >= 3 {
			s.tempMenzenTileIDs[i+9] -= 3
			countOfKanzenKotsu++

			mentsu := TileIDs{}
			mentsu[i+9] = 3
			s.addMentsu(mentsu, Anko)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenKotsu
			}
		}
	}
	return countOfKanzenKotsu
	// ˄
}

func (s *ShantenChecker) cutKanzenShuntsuAndGetCount() int {
	// ˅
	countOfKanzenShuntsu := 0
	i := 0
	//123,456のような完全に独立したシュンツを抜き出すための処理
	////【注意】番地0，10，20，30が「0」の必要あり。事前に赤ドラを移動させる処理をしておく。
	for i = 0; i < 30; i += 10 {
		//マンズ→ピンズ→ソーズ
		//123▲▲
		if s.tempMenzenTileIDs[i+1] == 2 && s.tempMenzenTileIDs[i+2] == 2 && s.tempMenzenTileIDs[i+3] == 2 && s.tempMenzenTileIDs[i+4] == 0 && s.tempMenzenTileIDs[i+5] == 0 {
			s.tempMenzenTileIDs[i+1] -= 2
			s.tempMenzenTileIDs[i+2] -= 2
			s.tempMenzenTileIDs[i+3] -= 2
			countOfKanzenShuntsu += 2

			mentsu := TileIDs{}
			mentsu[i+1] = 1
			mentsu[i+2] = 1
			mentsu[i+3] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}

		//▲234▲▲
		if s.tempMenzenTileIDs[i+1] == 0 && s.tempMenzenTileIDs[i+2] == 2 && s.tempMenzenTileIDs[i+3] == 2 && s.tempMenzenTileIDs[i+4] == 2 && s.tempMenzenTileIDs[i+5] == 0 && s.tempMenzenTileIDs[i+6] == 0 {
			s.tempMenzenTileIDs[i+2] -= 2
			s.tempMenzenTileIDs[i+3] -= 2
			s.tempMenzenTileIDs[i+4] -= 2
			countOfKanzenShuntsu += 2

			mentsu := TileIDs{}
			mentsu[i+2] = 1
			mentsu[i+3] = 1
			mentsu[i+4] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}

		//▲▲345▲▲
		if s.tempMenzenTileIDs[i+1] == 0 && s.tempMenzenTileIDs[i+2] == 0 && s.tempMenzenTileIDs[i+3] == 2 && s.tempMenzenTileIDs[i+4] == 2 && s.tempMenzenTileIDs[i+5] == 2 && s.tempMenzenTileIDs[i+6] == 0 && s.tempMenzenTileIDs[i+7] == 0 {
			s.tempMenzenTileIDs[i+3] -= 2
			s.tempMenzenTileIDs[i+4] -= 2
			s.tempMenzenTileIDs[i+5] -= 2
			countOfKanzenShuntsu += 2

			mentsu := TileIDs{}
			mentsu[i+3] = 1
			mentsu[i+4] = 1
			mentsu[i+5] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}

		//▲▲456▲▲
		if s.tempMenzenTileIDs[i+2] == 0 && s.tempMenzenTileIDs[i+3] == 0 && s.tempMenzenTileIDs[i+4] == 2 && s.tempMenzenTileIDs[i+5] == 2 && s.tempMenzenTileIDs[i+6] == 2 && s.tempMenzenTileIDs[i+7] == 0 && s.tempMenzenTileIDs[i+8] == 0 {
			s.tempMenzenTileIDs[i+4] -= 2
			s.tempMenzenTileIDs[i+5] -= 2
			s.tempMenzenTileIDs[i+6] -= 2
			countOfKanzenShuntsu += 2

			mentsu := TileIDs{}
			mentsu[i+4] = 1
			mentsu[i+5] = 1
			mentsu[i+6] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}

		//▲▲567▲▲
		if s.tempMenzenTileIDs[i+3] == 0 && s.tempMenzenTileIDs[i+4] == 0 && s.tempMenzenTileIDs[i+5] == 2 && s.tempMenzenTileIDs[i+6] == 2 && s.tempMenzenTileIDs[i+7] == 2 && s.tempMenzenTileIDs[i+8] == 0 && s.tempMenzenTileIDs[i+9] == 0 {
			s.tempMenzenTileIDs[i+5] -= 2
			s.tempMenzenTileIDs[i+6] -= 2
			s.tempMenzenTileIDs[i+7] -= 2
			countOfKanzenShuntsu += 2

			mentsu := TileIDs{}
			mentsu[i+5] = 1
			mentsu[i+6] = 1
			mentsu[i+7] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}

		//▲▲678▲
		if s.tempMenzenTileIDs[i+4] == 0 && s.tempMenzenTileIDs[i+5] == 0 && s.tempMenzenTileIDs[i+6] == 2 && s.tempMenzenTileIDs[i+7] == 2 && s.tempMenzenTileIDs[i+8] == 2 && s.tempMenzenTileIDs[i+9] == 0 {
			s.tempMenzenTileIDs[i+6] -= 2
			s.tempMenzenTileIDs[i+7] -= 2
			s.tempMenzenTileIDs[i+8] -= 2
			countOfKanzenShuntsu += 2

			mentsu := TileIDs{}
			mentsu[i+6] = 1
			mentsu[i+7] = 1
			mentsu[i+8] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}

		//▲▲789
		if s.tempMenzenTileIDs[i+5] == 0 && s.tempMenzenTileIDs[i+6] == 0 && s.tempMenzenTileIDs[i+7] == 2 && s.tempMenzenTileIDs[i+8] == 2 && s.tempMenzenTileIDs[i+9] == 2 {
			s.tempMenzenTileIDs[i+7] -= 2
			s.tempMenzenTileIDs[i+8] -= 2
			s.tempMenzenTileIDs[i+9] -= 2
			countOfKanzenShuntsu += 2

			mentsu := TileIDs{}
			mentsu[i+7] = 1
			mentsu[i+8] = 1
			mentsu[i+9] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}
	}

	for i = 0; i < 30; i += 10 {
		//マンズ→ピンズ→ソーズ
		//123▲▲
		if s.tempMenzenTileIDs[i+1] == 1 && s.tempMenzenTileIDs[i+2] == 1 && s.tempMenzenTileIDs[i+3] == 1 && s.tempMenzenTileIDs[i+4] == 0 && s.tempMenzenTileIDs[i+5] == 0 {
			s.tempMenzenTileIDs[i+1]--
			s.tempMenzenTileIDs[i+2]--
			s.tempMenzenTileIDs[i+3]--
			countOfKanzenShuntsu++

			mentsu := TileIDs{}
			mentsu[i+1] = 1
			mentsu[i+2] = 1
			mentsu[i+3] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}

		//▲234▲▲
		if s.tempMenzenTileIDs[i+1] == 0 && s.tempMenzenTileIDs[i+2] == 1 && s.tempMenzenTileIDs[i+3] == 1 && s.tempMenzenTileIDs[i+4] == 1 && s.tempMenzenTileIDs[i+5] == 0 && s.tempMenzenTileIDs[i+6] == 0 {
			s.tempMenzenTileIDs[i+2]--
			s.tempMenzenTileIDs[i+3]--
			s.tempMenzenTileIDs[i+4]--
			countOfKanzenShuntsu++

			mentsu := TileIDs{}
			mentsu[i+2] = 1
			mentsu[i+3] = 1
			mentsu[i+4] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}

		//▲▲345▲▲
		if s.tempMenzenTileIDs[i+1] == 0 && s.tempMenzenTileIDs[i+2] == 0 && s.tempMenzenTileIDs[i+3] == 1 && s.tempMenzenTileIDs[i+4] == 1 && s.tempMenzenTileIDs[i+5] == 1 && s.tempMenzenTileIDs[i+6] == 0 && s.tempMenzenTileIDs[i+7] == 0 {
			s.tempMenzenTileIDs[i+3]--
			s.tempMenzenTileIDs[i+4]--
			s.tempMenzenTileIDs[i+5]--
			countOfKanzenShuntsu++

			mentsu := TileIDs{}
			mentsu[i+3] = 1
			mentsu[i+4] = 1
			mentsu[i+5] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}

		//▲▲456▲▲
		if s.tempMenzenTileIDs[i+2] == 0 && s.tempMenzenTileIDs[i+3] == 0 && s.tempMenzenTileIDs[i+4] == 1 && s.tempMenzenTileIDs[i+5] == 1 && s.tempMenzenTileIDs[i+6] == 1 && s.tempMenzenTileIDs[i+7] == 0 && s.tempMenzenTileIDs[i+8] == 0 {
			s.tempMenzenTileIDs[i+4]--
			s.tempMenzenTileIDs[i+5]--
			s.tempMenzenTileIDs[i+6]--
			countOfKanzenShuntsu++

			mentsu := TileIDs{}
			mentsu[i+4] = 1
			mentsu[i+5] = 1
			mentsu[i+6] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}

		//▲▲567▲▲
		if s.tempMenzenTileIDs[i+3] == 0 && s.tempMenzenTileIDs[i+4] == 0 && s.tempMenzenTileIDs[i+5] == 1 && s.tempMenzenTileIDs[i+6] == 1 && s.tempMenzenTileIDs[i+7] == 1 && s.tempMenzenTileIDs[i+8] == 0 && s.tempMenzenTileIDs[i+9] == 0 {
			s.tempMenzenTileIDs[i+5]--
			s.tempMenzenTileIDs[i+6]--
			s.tempMenzenTileIDs[i+7]--
			countOfKanzenShuntsu++

			mentsu := TileIDs{}
			mentsu[i+5] = 1
			mentsu[i+6] = 1
			mentsu[i+7] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}

		//▲▲678▲
		if s.tempMenzenTileIDs[i+4] == 0 && s.tempMenzenTileIDs[i+5] == 0 && s.tempMenzenTileIDs[i+6] == 1 && s.tempMenzenTileIDs[i+7] == 1 && s.tempMenzenTileIDs[i+8] == 1 && s.tempMenzenTileIDs[i+9] == 0 {
			s.tempMenzenTileIDs[i+6]--
			s.tempMenzenTileIDs[i+7]--
			s.tempMenzenTileIDs[i+8]--
			countOfKanzenShuntsu++

			mentsu := TileIDs{}
			mentsu[i+6] = 1
			mentsu[i+7] = 1
			mentsu[i+8] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}

		//▲▲789
		if s.tempMenzenTileIDs[i+5] == 0 && s.tempMenzenTileIDs[i+6] == 0 && s.tempMenzenTileIDs[i+7] == 1 && s.tempMenzenTileIDs[i+8] == 1 && s.tempMenzenTileIDs[i+9] == 1 {
			s.tempMenzenTileIDs[i+7]--
			s.tempMenzenTileIDs[i+8]--
			s.tempMenzenTileIDs[i+9]--
			countOfKanzenShuntsu++

			mentsu := TileIDs{}
			mentsu[i+7] = 1
			mentsu[i+8] = 1
			mentsu[i+9] = 1
			s.addMentsu(mentsu, MenzenShuntsu)
			if !s.agarikei.Janto.IsEmpty() && !s.agarikei.Mentsu4.IsEmpty() {
				return countOfKanzenShuntsu
			}
		}
	}
	return countOfKanzenShuntsu

	// ˄
}

func (s *ShantenChecker) cutKanzenKoritsuAndGetCount() int {
	// ˅
	countOfKanzenKoritsu := 0
	i := 0
	j := 0
	//字牌の完全孤立牌を抜き出す
	for i = 31; i < 38; i++ {
		if s.tempMenzenTileIDs[i] == 1 {
			// s.kanzenKoritsu[i]++
			s.tempMenzenTileIDs[i]--
			countOfKanzenKoritsu++
		}
	}

	//数牌の完全孤立牌を抜き出す
	for i = 0; i < 30; i += 10 {
		//マンズ→ピンズ→ソーズ
		//1の孤立牌を抜く
		if s.tempMenzenTileIDs[i+1] == 1 && s.tempMenzenTileIDs[i+2] == 0 && s.tempMenzenTileIDs[i+3] == 0 {
			// s.kanzenKoritsu[i+1]++
			s.tempMenzenTileIDs[i+1]--
			countOfKanzenKoritsu++
		}
		//2の完全孤立牌を抜く
		if s.tempMenzenTileIDs[i+1] == 0 && s.tempMenzenTileIDs[i+2] == 1 && s.tempMenzenTileIDs[i+3] == 0 && s.tempMenzenTileIDs[i+4] == 0 {
			// s.kanzenKoritsu[i+2]++
			s.tempMenzenTileIDs[i+2]--
			countOfKanzenKoritsu++
		}

		//3~7の完全孤立牌を抜く
		for j = 0; j < 5; j++ {
			if s.tempMenzenTileIDs[i+j+1] == 0 && s.tempMenzenTileIDs[i+j+2] == 0 && s.tempMenzenTileIDs[i+j+3] == 1 && s.tempMenzenTileIDs[i+j+4] == 0 && s.tempMenzenTileIDs[i+j+5] == 0 {
				// s.kanzenKoritsu[i+j+3]++
				s.tempMenzenTileIDs[i+j+3]--
				countOfKanzenKoritsu++
			}
		}
		//8の完全孤立牌を抜く
		if s.tempMenzenTileIDs[i+6] == 0 && s.tempMenzenTileIDs[i+7] == 0 && s.tempMenzenTileIDs[i+8] == 1 && s.tempMenzenTileIDs[i+9] == 0 {
			// s.kanzenKoritsu[i+8]++
			s.tempMenzenTileIDs[i+8]--
			countOfKanzenKoritsu++
		}
		//9の完全孤立牌を抜く
		if s.tempMenzenTileIDs[i+7] == 0 && s.tempMenzenTileIDs[i+8] == 0 && s.tempMenzenTileIDs[i+9] == 1 {
			// s.kanzenKoritsu[i+9]++
			s.tempMenzenTileIDs[i+9]--
			countOfKanzenKoritsu++
		}
	}
	return countOfKanzenKoritsu

	// ˄
}

func (s *ShantenChecker) cutKotsu(tileID int) bool {
	// ˅
	if s.tempMenzenTileIDs[tileID] >= 3 {
		s.tempMenzenTileIDs[tileID] -= 3
		return true
	}
	return false
	// ˄
}

func (s *ShantenChecker) addKotsu(tileID int) {
	// ˅
	s.tempMenzenTileIDs[tileID] += 3
	// ˄
}

func (s *ShantenChecker) cutShuntsu(firstTileID int) bool {
	// ˅
	if s.tempMenzenTileIDs[firstTileID] >= 1 && s.tempMenzenTileIDs[firstTileID+1] >= 1 && s.tempMenzenTileIDs[firstTileID+2] >= 1 {
		s.tempMenzenTileIDs[firstTileID] -= 1
		s.tempMenzenTileIDs[firstTileID+1] -= 1
		s.tempMenzenTileIDs[firstTileID+2] -= 1
		return true
	}
	return false
	// ˄
}

func (s *ShantenChecker) addShuntsu(firstTileID int) {
	// ˅
	s.tempMenzenTileIDs[firstTileID] += 1
	s.tempMenzenTileIDs[firstTileID+1] += 1
	s.tempMenzenTileIDs[firstTileID+2] += 1
	// ˄
}

func (s *ShantenChecker) cutToitsu(tileID int) bool {
	// ˅
	if s.tempMenzenTileIDs[tileID] >= 2 {
		s.tempMenzenTileIDs[tileID] -= 2
		return true
	}
	return false
	// ˄
}

func (s *ShantenChecker) addToitsu(tileID int) {
	// ˅
	s.tempMenzenTileIDs[tileID] += 2
	// ˄
}

// ˅

func (s *ShantenChecker) addMentsu(mentsu TileIDs, mentsuType MentsuType) {
	if s.agarikei.Mentsu1.IsEmpty() && s.agarikei.Mentsu1Type == Null {
		s.agarikei.Mentsu1 = mentsu
		s.agarikei.Mentsu1Type = mentsuType
		return
	}
	if s.agarikei.Mentsu2.IsEmpty() && s.agarikei.Mentsu2Type == Null {
		s.agarikei.Mentsu2 = mentsu
		s.agarikei.Mentsu2Type = mentsuType
		return
	}
	if s.agarikei.Mentsu3.IsEmpty() && s.agarikei.Mentsu3Type == Null {
		s.agarikei.Mentsu3 = mentsu
		s.agarikei.Mentsu3Type = mentsuType
		return
	}
	if s.agarikei.Mentsu4.IsEmpty() && s.agarikei.Mentsu4Type == Null {
		s.agarikei.Mentsu4 = mentsu
		s.agarikei.Mentsu4Type = mentsuType
		return
	}
}

func (s *ShantenChecker) undoMentsu() {
	if !s.agarikei.Mentsu4.IsEmpty() && s.agarikei.Mentsu4Type != Null {
		s.agarikei.Mentsu4.Reset()
		s.agarikei.Mentsu4Type = Null
		return
	}
	if !s.agarikei.Mentsu3.IsEmpty() && s.agarikei.Mentsu3Type != Null {
		s.agarikei.Mentsu3.Reset()
		s.agarikei.Mentsu3Type = Null
		return
	}
	if !s.agarikei.Mentsu2.IsEmpty() && s.agarikei.Mentsu2Type != Null {
		s.agarikei.Mentsu2.Reset()
		s.agarikei.Mentsu2Type = Null
		return
	}
	if !s.agarikei.Mentsu1.IsEmpty() && s.agarikei.Mentsu1Type != Null {
		s.agarikei.Mentsu1.Reset()
		s.agarikei.Mentsu1Type = Null
		return
	}
}

func (s *ShantenChecker) preparation(player *Player) {
	// 直接はいじらないようにしよう
	p := player
	s.menzenTileIDs = TileIDs{}
	s.tempMenzenTileIDs = TileIDs{}

	s.agarikei = Agarikei{}

	s.countOfMentsu = 0
	s.countOfToitsu = 0
	s.countOfKotsu = 0
	s.countOfShuntsu = 0
	s.countOfTatsu = 0
	s.countOfAnkan = 0
	s.countOfMinkan = 0

	s.shantenTemp = 8
	s.shantenNormal = 8
	s.countOfKanzenKotsu = 0
	s.countOfKanzenShuntsu = 0
	s.countOfKanzenKoritsu = 0

	for _, tileID := range s.minkoTileIDs {
		tileID.Reset()
	}
	for _, tileID := range s.nakishuntsuTileIDs {
		tileID.Reset()
	}
	for _, tileID := range s.ankanTileIDs {
		tileID.Reset()
	}
	for _, tileID := range s.minkanTileIDs {
		tileID.Reset()
	}
	s.agarikei.Janto.Reset()
	s.agarikei.Mentsu1.Reset()
	s.agarikei.Mentsu2.Reset()
	s.agarikei.Mentsu3.Reset()
	s.agarikei.Mentsu4.Reset()

	s.menzenTileIDs.Reset()
	s.tempMenzenTileIDs.Reset()
	// s.kanzenKoritsu.Reset()

	// 開かれた牌の読み取り（暗槓とかポンとか）
	for _, OpenedTiles := range []*OpenedTiles{
		p.OpenedTile1,
		p.OpenedTile2,
		p.OpenedTile3,
		p.OpenedTile4,
	} {
		if OpenedTiles != nil && len(OpenedTiles.Tiles) != 0 {
			switch *OpenedTiles.OpenType {
			case OPEN_PON:
				minko := TileIDs{}
				for _, tile := range OpenedTiles.Tiles {
					minko[int(tile.ID)]++
				}
				s.minkoTileIDs = append(s.minkoTileIDs, minko)
				s.countOfKanzenKotsu++
			case OPEN_CHI:
				nakishuntsu := TileIDs{}
				for _, tile := range OpenedTiles.Tiles {
					nakishuntsu[int(tile.ID)]++
				}
				s.nakishuntsuTileIDs = append(s.nakishuntsuTileIDs, nakishuntsu)
				s.countOfKanzenShuntsu++
			case OPEN_ANKAN:
				ankan := TileIDs{}
				for _, tile := range OpenedTiles.Tiles {
					ankan[int(tile.ID)]++
				}
				s.ankanTileIDs = append(s.ankanTileIDs, ankan)
				s.countOfAnkan++
			case OPEN_DAIMINKAN:
				fallthrough
			case OPEN_KAKAN:
				minkan := TileIDs{}
				for _, tile := range OpenedTiles.Tiles {
					minkan[int(tile.ID)]++
				}
				s.minkanTileIDs = append(s.minkanTileIDs, minkan)
				s.countOfMinkan++
			}
		}
	}

	for _, mentsu := range s.ankanTileIDs {
		s.addMentsu(mentsu, Ankan)
	}

	for _, mentsu := range s.minkanTileIDs {
		s.addMentsu(mentsu, Minkan)
	}

	for _, mentsu := range s.minkoTileIDs {
		s.addMentsu(mentsu, Minko)
	}

	for _, mentsu := range s.nakishuntsuTileIDs {
		s.addMentsu(mentsu, NakiShuntsu)
	}

	// 開かれていない牌の読み取り
	for _, tile := range append(p.Hand, p.TsumoriTile) {
		if tile == nil {
			continue
		}
		s.menzenTileIDs[int(tile.ID)]++
		s.tempMenzenTileIDs[int(tile.ID)]++
	}
}

func (s *ShantenChecker) calcMachihai() map[int]interface{} {
	haiCount := 0
	set := map[int]interface{}{}
	i := 0

	for i = 0; i < len(s.tempMenzenTileIDs); i++ {
		haiCount += s.tempMenzenTileIDs[i]
	}

	if haiCount == 2 {
		return set
	}

	// 単騎待ち
	for i = 1; i < 38; i++ {
		if s.tempMenzenTileIDs[i] == 2 {
			set[i] = struct{}{}
		}
	}

	// 辺張待ち
	for i = 0; i <= 2; i++ {
		if s.tempMenzenTileIDs[i*10+1] == 1 && s.tempMenzenTileIDs[i*10+2] == 1 {
			set[i*10+3] = struct{}{}
		}
		if s.tempMenzenTileIDs[i*10+8] == 1 && s.tempMenzenTileIDs[i*10+9] == 1 {
			set[i*10+7] = struct{}{}
		}
	}
	// 両面待ち
	for i = 0; i <= 29; i++ {
		if !((i >= 2 && i <= 7) || (i >= 12 && i <= 17) || (i >= 22 && i <= 27)) {
			continue
		}
		if s.tempMenzenTileIDs[i] == 1 && s.tempMenzenTileIDs[i+1] == 1 {
			set[i-1] = struct{}{}
			set[i+2] = struct{}{}
		}
	}
	// 嵌張待ち
	for i = 0; i <= 29; i++ {
		if !((i >= 2 && i <= 7) || (i >= 12 && i <= 17) || (i >= 22 && i <= 27)) {
			continue
		}
		if s.tempMenzenTileIDs[i] == 1 && s.tempMenzenTileIDs[i+2] == 1 {
			set[i+1] = struct{}{}
		}
	}
	return set
}

func NewShantenChecker() *ShantenChecker {
	return &ShantenChecker{
		yakuList: GenerateYakusDefault(),
	}
}

// ˄

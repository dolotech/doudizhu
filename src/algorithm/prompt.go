package algorithm

func (this *AnalyseCards) Prompt(discards []byte, length uint8) []AnalyseCards {
	var cardType AnalyseCards
	cardType.Analyse(discards)
	result := make([]AnalyseCards, 0, length)
	if cardType.GetKind() == WANG_ZHA { // 上手出王炸，直接要不起
		return result
	}
	switch cardType.GetKind() {
	case FEI_JI_DAI_DAN: // 飞机带单
		if length >= 8 {
			this.PromptFeijiDaiDan(cardType.Len(), cardType.GetValue(), &result)
		}
	case FEI_JI_DAI_DUI: // 飞机带对
		if length >= 10 {
			this.PromptFeijiDaiDui(cardType.Len(), cardType.GetValue(), &result)
		}
	case SI_DAI_DAN: // 四带单
		if length >= 6 {
			this.PromptSiDaiDan(cardType.GetValue(), &result)
		}
	case SI_DAI_DUI: // 四带对
		if length >= 8 {
			this.PromptSiDaiDui(cardType.GetValue(), &result)
		}
	case SAN_SHUN_ZI: // 三顺
		if length >= 6 {
			this.ShunZiCommon(cardType.Len(), cardType.GetValue(), 3, SAN_SHUN_ZI, &result)
		}
	case DAN_SHUN_ZI: // 单顺
		if length >= 5 {
			this.ShunZiCommon(cardType.Len(), cardType.GetValue(), 1, DAN_SHUN_ZI, &result)
		}
	case LIAN_DUI: // 连对
		if length >= 6 {
			this.ShunZiCommon(cardType.Len(), cardType.GetValue(), 2, LIAN_DUI, &result)
		}
	case ZHA_DAN: // 炸弹
		if length >= 4 {
			this.PromptZhaDan(cardType.GetValue(), &result)
		}
	case DAN_TIAO: //单条
		this.PromptDanTiao(cardType.GetValue(), &result)
	case DUI_ZI: // 对子
		if length >= 2 {
			this.PromptDuiZi(cardType.GetValue(), &result)
		}
	case SAN_TIAO: // 三条
		if length >= 3 {
			this.PromptSanTiao(cardType.GetValue(), &result)
		}
	case SAN_DAI_YI: // 三带单
		if length >= 4 {
			this.PromptSanDaiDan(cardType.GetValue(), &result)
		}
	case SAN_DAI_YI_DUI: // 三带对
		if length >= 5 {
			this.PromptSanDaiDui(cardType.GetValue(), &result)
		}
	}
	// 找手牌的炸弹对抗上手的普通牌
	if cardType.GetKind() != ZHA_DAN {
		this.PromptZhaDan(cardType.GetValue(), &result)
	}
	// 王炸
	this.PromptWangZha(&result)
	return result
}

func (this *AnalyseCards) PromptSanDaiDui(value byte, cardTypes *[]AnalyseCards) {
	if value == TWO {
		return
	}

	if value == 0 {
		value = byte(THREE - 1)
	}

	end:= this.End()
	start:= this.Start()
	if end > TWO{
		end =TWO
	}

	for i := end; i > value; i-- {
		if this.Get(i) < 3 {
			continue
		}

		for j := start; j <= end; j++ {
			if j != i {
				if this.Get(j) > 1 {
					var anCards AnalyseCards
					anCards.Incr(i, 3)
					anCards.Incr(j, 2)
					anCards.SetKind(SAN_DAI_YI_DUI)
					anCards.SetValue(i)
					*cardTypes = append(*cardTypes, anCards)
					break
				}
			}
		}
	}
	return
}
func (this *AnalyseCards) PromptSanDaiDan(value byte, cardTypes *[]AnalyseCards) {
	if value == TWO {
		return
	}
	if value == 0 {
		value = byte(THREE - 1)
	}
	end:= this.End()
	start:= this.Start()
	if end > TWO{
		end =TWO
	}
	for i := end; i > value; i-- {
		if this.Get(i) < 3 {
			continue
		}
		for j := start; j <= end; j++ {
			if j != i {
				if this.Get(j) > 0 {
					var anCards AnalyseCards
					anCards.Incr(i, 3)
					anCards.Incr(j, 1)
					anCards.SetKind(SAN_DAI_YI)
					anCards.SetValue(i)
					*cardTypes = append(*cardTypes, anCards)
					break
				}
			}
		}
	}
	return
}
func (this *AnalyseCards) PromptSanTiao(value byte, cardTypes *[]AnalyseCards) {
	if value == TWO {
		return
	}
	if value == 0 {
		value = byte(THREE - 1)
	}

	end:= this.End()
	if end > TWO{
		end =TWO
	}
	for i := end; i > value; i-- {
		if this.Get(i) < 3 {
			continue
		}
		var anCards AnalyseCards
		anCards.Incr(i, 3)

		anCards.SetKind(SAN_TIAO)
		anCards.SetValue(i)
		*cardTypes = append(*cardTypes, anCards)
	}
	return
}
func (this *AnalyseCards) PromptDuiZi(value byte, cardTypes *[]AnalyseCards) {
	if value == TWO {
		return
	}
	if value == 0 {
		value = byte(THREE - 1)
	}
	end:= this.End()
	if end > TWO{
		end =TWO
	}
	for i := end; i > value; i-- {
		if this.Get(i) > 1 {
			var anCards AnalyseCards
			anCards.Incr(i, 2)
			anCards.SetKind(DUI_ZI)
			anCards.SetValue(i)
			*cardTypes = append(*cardTypes, anCards)
		}
	}
	return
}
func (this *AnalyseCards) PromptDanTiao(value byte, cardTypes *[]AnalyseCards) {
	if value == DA_WANG {
		return
	}
	if value == 0 {
		value = byte(THREE - 1)
	}

	end:= this.End()

	for i := end; i > value; i-- {
		if this.Get(i) > 0 {
			var anCards AnalyseCards
			anCards.Incr(i, 1)
			anCards.SetKind(DAN_TIAO)
			anCards.SetValue(i)
			*cardTypes = append(*cardTypes, anCards)
		}
	}
	return
}

func (this *AnalyseCards) ShunZiCommon(length uint8, value byte, leng uint8, kind uint8, cardTypes *[]AnalyseCards) {
	discardsLength := length / leng

	end:= this.End()
	//start:= this.Start()
	s:= end
	if s > A{
		s = A
	}

	for i := s; i > value; i-- {
		for k := uint8(0); k < discardsLength; k++ {
			if this.Get(i-k) < leng {
				goto BR
			}
			if k+1 == discardsLength {
				var anCards AnalyseCards
				anCards.Padding(discardsLength, leng, i)
				anCards.SetKind(kind)
				anCards.SetValue(i)
				*cardTypes = append(*cardTypes, anCards)
				goto BR
			}
		}
	BR:
	}
	return
}

func (this *AnalyseCards) PromptSiDaiDan(value byte, cardTypes *[]AnalyseCards) {
	if value == TWO {
		return
	}
	if value == 0 {
		value = byte(THREE - 1)
	}

	end:= this.End()
	start:= this.Start()
	s:= end
	if s > TWO{
		s = TWO
	}

	for i := s; i > value; i-- {
		var lastCards byte
		if this.Get(i) != 4 {
			continue
		}

		for j := start; j <= s; j++ {
			if j != i {
				if this.Get(j) > 0 {
					if lastCards == 0 {
						lastCards = j
						continue
					}
					var anCards AnalyseCards
					anCards.Incr(i, 4)
					anCards.Incr(lastCards, 1)
					anCards.Incr(j, 1)
					anCards.SetKind(SI_DAI_DAN)
					anCards.SetValue(i)
					*cardTypes = append(*cardTypes, anCards)
					break
				}
			}
		}
	}
	return
}
func (this *AnalyseCards) PromptSiDaiDui(value byte, cardTypes *[]AnalyseCards) {
	if value == TWO {
		return
	}
	if value == 0 {
		value = byte(THREE - 1)
	}
	end:= this.End()
	start:= this.Start()
	s:= end
	if s > TWO{
		s = TWO
	}
	for i := s; i > value; i-- {
		count := uint8(0)
		var lastCards byte
		if this.Get(i) != 4 {
			continue
		}

		for j := start; j <= s; j++ {
			if j != i {
				if this.Get(j) > 1 {
					count ++
					if lastCards == 0 {
						lastCards = j
					}
				}
				if count == 2 {
					var anCards AnalyseCards
					anCards.Incr(i, 4)
					anCards.Incr(lastCards, 2)
					anCards.Incr(j, 2)

					anCards.SetKind(SI_DAI_DUI)
					anCards.SetValue(i)
					*cardTypes = append(*cardTypes, anCards)
					break
				}
			}
		}
	}
	return
}

func (this *AnalyseCards) PromptFeijiDaiDan(length uint8, value byte, cardTypes *[]AnalyseCards) {
	if value == A {
		return
	}
	start:= this.Start()
	end:= this.End()
	discardsLength := length / 4
	s:= end
	if s > A{
		s = A
	}
	for i := A; i > value; i-- {
		if this.Get(i) == 0 {
			continue
		}


		if end>TWO{
			end = TWO
		}
		count := uint8(0)
		var anCards AnalyseCards
		for k := uint8(0); k < discardsLength; k++ {
			if this.Get(i-k) < 3 {
				goto BR
			}
		}

		for j := start; j <= end; j++ {
			c := this.Get(j)
			if c > 0 {
				if c <= 3 && j <= i && j > i-discardsLength {
					continue
				}
				anCards.Incr(j, 1)
				count ++
				if count == discardsLength {
					anCards.Padding(discardsLength, 3, i)

					anCards.SetKind(FEI_JI_DAI_DAN)
					anCards.SetValue(i)
					*cardTypes = append(*cardTypes, anCards)
					goto BR
				}
			}
		}
	BR:
	}
	return
}
func (this *AnalyseCards) PromptZhaDan(value byte, cardTypes *[]AnalyseCards) {
	if value == TWO {
		return
	}
	if value == 0 {
		value = byte(THREE - 1)
	}
	end:= this.End()
	if end > TWO{
		end = TWO
	}
	for i := end; i > value; i-- {
		if this.Get(i) == 4 {
			var anCards AnalyseCards
			anCards.Incr(i, 4)

			anCards.SetKind(ZHA_DAN)
			anCards.SetValue(i)
			*cardTypes = append(*cardTypes, anCards)
		}
	}
	return
}
func (this *AnalyseCards) PromptWangZha(cardTypes *[]AnalyseCards) {
	if this.Get(DA_WANG) > 0 && this.Get(XIAO_WANG) > 0 {
		var anCards AnalyseCards
		anCards.Incr(DA_WANG, 1)
		anCards.Incr(XIAO_WANG, 1)
		anCards.SetKind(WANG_ZHA)
		anCards.SetValue(DA_WANG)
		*cardTypes = append(*cardTypes, anCards)
	}
	return
}

func (this *AnalyseCards) PromptFeijiDaiDui(length uint8, value byte, cardTypes *[]AnalyseCards) {
	if value == A {
		return
	}
	end:= this.End()
	start:= this.Start()
	s:= end
	if s > A{
		s = A
	}
	discardsLength := length / 5
	for i := s; i > value; i-- {
		if this.Get(i) == 0 {
			continue
		}
		if this.Get(i) == 1 {
			return
		}
		count := uint8(0)
		var anCards AnalyseCards = 0
		e := end
		if e > TWO{
			e = TWO
		}

		for k := uint8(0); k < discardsLength; k++ {
			if this.Get(i-k) < 3 {
				goto BR
			}
		}

		for j := start; j <= e; j++ {
			if this.Get(j) == 0 {
				continue
			}
			if j > i || j <= i-discardsLength {
				if this.Get(j) == 4 {
					anCards.Incr(j, 4)
					count += 2
				} else if this.Get(j) == 2 {
					anCards.Incr(j, 2)
					count ++
				}
				if count == discardsLength {
					anCards.Padding(discardsLength, 3, i)
					anCards.SetKind(FEI_JI_DAI_DUI)
					anCards.SetValue(i)
					*cardTypes = append(*cardTypes, anCards)
					goto BR
				}
			}
		}
	BR:
	}
	return
}

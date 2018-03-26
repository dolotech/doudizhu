package algorithm

// 找出手牌中所有的牌型组合
func (this *AnalyseCards) GetGroup() ([]AnalyseCards) {
	var length = this.Len()
	result := make([]AnalyseCards, 0, length*2)
	if length == 0 {
		return result
	}

	this.PromptDanTiao(0, &result)

	if length >= 4 {
		this.PromptZhaDan(0, &result)
	}

	if length >= 2 {
		this.PromptDuiZi(0, &result)
	}
	if length >= 3 {
		this.PromptSanTiao(0, &result)
	}
	if length >= 2 {
		this.PromptWangZha(&result)
	}
	if length >= 4 {
		this.AllSanDaiDan(&result)
	}
	if length >= 5 {
		this.AllSanDaiDui(&result)
	}

	if length >= 6 {
		this.AllSiDaiDan(&result)
	}
	if length >= 8 {
		this.AllSiDaiDui(&result)
	}

	if length >= 8 {
		this.AllFeijiDaiDan(&result)
	}
	if length >= 10 {
		this.AllFeijiDaiDui(&result)
	}
	if length >= 6 {
		this.AllLianDui(&result)
	}
	if length >= 6 {
		this.AllSanShun(&result)
	}
	if length >= 5 {
		this.AllShunZi(&result)
	}

	return result
}

// 找出手牌中所有的顺子
func (this *AnalyseCards) AllShunZi(cardTypes *[]AnalyseCards) {
	end := this.End()
	start := this.Start()
	gap := end - start + 1
	if gap < 5 {
		return
	}
	if end > A {
		end = A
	}
	for i := start; i <= end; i++ {
		if this.Get(i) >= 1 {
			for j := uint8(5); j <= gap; j++ {
				k := uint8(0)
				for k = uint8(1); k < j; k++ {
					if i+k > end || this.Get(i+k) == 0 {
						break
					}
				}
				if k == j {
					var anCards AnalyseCards
					for n := uint8(0); n < j; n++ {
						anCards.Incr(i+n, 1)
					}
					anCards.SetKind(DAN_SHUN_ZI)
					anCards.SetValue(i + j - 1)
					*cardTypes = append(*cardTypes, anCards)
				}
			}
		}
	}
	return
}

// 找出手牌中所有的连对
func (this *AnalyseCards) AllLianDui(cardTypes *[]AnalyseCards) {
	end := this.End()
	start := this.Start()
	if end > A {
		end = A
	}

	gap := end - start + 1
	if gap < 3 {
		return
	}
	for i := start; i <= end; i++ {
		if this.Get(i) >= 2 {
			for j := uint8(3); j <= gap; j++ {
				k := uint8(0)
				for k = uint8(1); k < j; k++ {
					if i+k > end || this.Get(i+k) < 2 {
						break
					}
				}
				if k == j {
					var anCards AnalyseCards
					for n := uint8(0); n < j; n++ {
						anCards.Incr(i+n, 2)
					}
					anCards.SetKind(LIAN_DUI)
					anCards.SetValue(i + j - 1)
					*cardTypes = append(*cardTypes, anCards)
				}
			}
		}
	}
	return
}

// 找出手牌中所有的三顺
func (this *AnalyseCards) AllSanShun(cardTypes *[]AnalyseCards) {
	end := this.End()
	start := this.Start()
	if end > A {
		end = A
	}

	gap := end - start + 1
	if gap < 2 {
		return
	}
	for i := start; i <= end; i++ {
		if this.Get(i) >= 3 {
			for j := uint8(2); j <= gap; j++ {
				k := uint8(0)
				for k = 1; k < j; k++ {
					if i+k > end || this.Get(i+k) < 3 {
						break
					}
				}
				if k == j {
					var anCards AnalyseCards
					for n := uint8(0); n < j; n++ {
						anCards.Incr(i+n, 3)
					}
					anCards.SetKind(SAN_SHUN_ZI)
					anCards.SetValue(i + j - 1)
					*cardTypes = append(*cardTypes, anCards)
				}
			}
		}
	}
	return
}

// 找出手牌中所有的飞机带单
func (this *AnalyseCards) AllFeijiDaiDan(cardTypes *[]AnalyseCards) {
	arr := make([]AnalyseCards, 0, this.Len())
	this.AllSanShun(&arr)
	for key := 0; key < len(arr); key++ {
		cards := arr[key]
		l := cards.Len()
		an := this.Sub(cards)
		anL := an.Len()
		if l/3 > anL {
			continue
		}
		anFlat := an.Flat()

		combo := CombineUnique(anFlat, int(l/3))
		//combo= combo[:l/3]
		for i := 0; i < len(combo); i++ {
			//glog.Errorln(combo[i],cards.Flat())
			c := cards
			for j := uint8(0); j < l/3; j++ {
				c.Incr(combo[i][j], 1)
			}
			c.SetKind(FEI_JI_DAI_DAN)
			c.SetValue(cards.GetValue())
			*cardTypes = append(*cardTypes, c)
		}
	}
	return
}

// 找出手牌中所有的飞机带对
func (this *AnalyseCards) AllFeijiDaiDui(cardTypes *[]AnalyseCards) {
	arr := make([]AnalyseCards, 0, this.Len())
	this.AllSanShun(&arr)

	for _, cards := range arr {
		l := cards.Len()
		an := this.Sub(cards)
		anL := an.Len()
		if l/3 > anL/2 {
			continue
		}

		anFlat := make([]byte, 0, anL/2)
		for i := an.Start(); i <= an.End(); i++ {
			if an.Get(i) >= 2 {
				anFlat = append(anFlat, i)
			}
		}

		combo := CombineUnique(anFlat, int(l/3))

		//combo= combo[:l/3]
		for i := 0; i < len(combo); i++ {
			c := cards
			for j := uint8(0); j < l/3; j++ {
				c.Incr(combo[i][j], 2)
			}
			c.SetKind(FEI_JI_DAI_DUI)
			c.SetValue(cards.GetValue())
			*cardTypes = append(*cardTypes, c)
		}
	}
	return
}

func (this *AnalyseCards) AllSiDaiDui(cardTypes *[]AnalyseCards) {

	end := this.End()
	if end > TWO {
		end = TWO
	}
	start := this.Start()
	for i := end; i >= start; i-- {
		if this.Get(i) != 4 {
			continue
		}
		an := *this
		an.Incr(i, -4)
		anFlat := make([]byte, 0, 2)
		for j := an.Start(); j <= an.End(); j++ {
			if an.Get(j) == 2 {
				anFlat = append(anFlat, j)
			} else if an.Get(j) == 4 {
				var anCards AnalyseCards
				anCards.Incr(i, 4)
				anCards.Incr(j, 4)
				anCards.SetKind(SI_DAI_DUI)
				anCards.SetValue(j)
				*cardTypes = append(*cardTypes, anCards)
			}
		}

		combo := CombineUnique(anFlat, 2)
		l := uint8(len(combo))
		for j := uint8(0); j < l; j++ {
			var anCards AnalyseCards
			anCards.Incr(i, 4)
			anCards.Incr(combo[j][0], 2)
			anCards.Incr(combo[j][1], 2)
			anCards.SetKind(SI_DAI_DUI)
			anCards.SetValue(i)
			*cardTypes = append(*cardTypes, anCards)
		}
	}
	return
}

func (this *AnalyseCards) AllSiDaiDan(cardTypes *[]AnalyseCards) {
	end := this.End()
	start := this.Start()
	if end > TWO {
		end = TWO
	}
	for i := end; i >= start; i-- {
		if this.Get(i) != 4 {
			continue
		}
		anFlat := make([]byte, 0, 2)
		for j := this.Start(); j <= this.End(); j++ {
			if i != j && this.Get(j) > 0 {
				anFlat = append(anFlat, j)
			}
		}

		combo := CombineUnique(anFlat, 2)
		for j := this.Start(); j <= this.End(); j++ {
			if i != j && this.Get(j) >= 2 {
				combo = append(combo, []byte{j, j})
			}
		}
		l := uint8(len(combo))
		for j := uint8(0); j < l; j++ {
			var anCards AnalyseCards
			anCards.Incr(i, 4)
			anCards.Incr(combo[j][0], 1)
			anCards.Incr(combo[j][1], 1)
			anCards.SetKind(SI_DAI_DAN)
			anCards.SetValue(i)
			*cardTypes = append(*cardTypes, anCards)
		}
	}
	return
}

func (this *AnalyseCards) AllSanDaiDui(cardTypes *[]AnalyseCards) {
	end := this.End()
	start := this.Start()
	if end > TWO {
		end = TWO
	}
	for i := end; i >= start; i-- {
		if this.Get(i) >=3 {
			for j := start; j <= end; j++ {
				if i != j && this.Get(j) > 1 {
					var anCards AnalyseCards
					anCards.Incr(i, 3)
					anCards.Incr(j, 2)
					anCards.SetKind(SAN_DAI_YI_DUI)
					anCards.SetValue(i)
					*cardTypes = append(*cardTypes, anCards)
				}
			}
		}
	}
	return
}
func (this *AnalyseCards) AllSanDaiDan(cardTypes *[]AnalyseCards) {
	end := this.End()
	start := this.Start()
	if end > TWO {
		end = TWO
	}
	for i := end; i >= start; i-- {
		if this.Get(i) >= 3 {
			for j := start; j <= end; j++ {
				if i != j && this.Get(j) > 0 {
					var anCards AnalyseCards
					anCards.Incr(i, 3)
					anCards.Incr(j, 1)
					anCards.SetKind(SAN_DAI_YI_DUI)
					anCards.SetValue(i)
					*cardTypes = append(*cardTypes, anCards)
				}
			}
		}
	}
	return
}

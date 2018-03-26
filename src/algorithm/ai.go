package algorithm

// 搜索当前手牌最优组合

func AnalyseWeightMax(anList *[]AnalyseCards, cards AnalyseCards) int {
	array := cards.GetGroup()
	if len(array) == 0 {
		return 0
	}

	for _, v := range array {
		t := cards.Sub(v)
		// 提前判断，一手牌能出完的时候
		if t.Len() == 0 {
			*anList = append(*anList, v)
			return v.Weight()
		}
	}

	v := SearchMaxWeight(array)// 拿出当前权
	t := cards.Sub(v)
	tweight := v.Weight() + AnalyseWeightMax(anList, t)
	*anList = append(*anList, v)
	return tweight
}


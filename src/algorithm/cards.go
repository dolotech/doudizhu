package algorithm

//发牌
// 最后一个返回为预留3张地主牌
func Deal() ([]byte, []byte, []byte, []byte) {
	cards := Shuffle()
	cards4 := make([]byte, 3)
	copy(cards4, cards[:3])
	cards3 := make([]byte, 17)
	copy(cards3, cards[3:20])
	cards2 := make([]byte, 17)
	copy(cards2, cards[20:37])
	cards1 := make([]byte, 17)
	copy(cards1, cards[37:])
	return cards1, cards2, cards3, cards4
}

type Cards struct {
	Cards        []byte //原始牌
	analyseCards AnalyseCards
}

func (this *Cards) Value() byte {
	return this.analyseCards.GetValue()
}

func (this *Cards) Weight() uint8 {
	return this.analyseCards.GetWeight()
}
func (this *Cards) Kind() uint8 {
	return this.analyseCards.GetKind()
}

func (this *Cards) Len() uint8 {
	return uint8(len(this.Cards))
}

func BuildCards(cards []byte) *Cards {
	c := &Cards{}
	c.AnalyseSort(cards)
	return c
}

func (c *Cards) AnalyseSort(cards []byte) {
	c.Cards = cards
	c.analyseCards.Analyse(cards)
}

// 分析选牌是否大于上手所出的牌
func (c *Cards) Follow(discard []byte) bool {
	return c.analyseCards.Follow(discard)
}

//分析牌型
func (this *Cards) AnalyseUnSort(cards []byte) {
	this.Cards = cards
	this.analyseCards.Set(cards)
}

// 获取所有可能的牌型组合
func (this *Cards) GetGroup() (result []AnalyseCards) {
	return this.analyseCards.GetGroup()
}

// 压牌提示
func (this *Cards) Prompt(discards []byte) (result []*Cards) {
	if len(this.Cards) == 0 || len(discards) == 0 {
		return
	}
	arr := this.analyseCards.Prompt(discards, this.Len())
	// 返回选中的牌，需要把牌转换为正常的牌面值
	for _, v := range arr {
		c := &Cards{analyseCards: v}
		c.Cards = v.ColorRecover(this.Cards)
		result = append(result, c)
	}
	return
}

package algorithm

// 分析选牌是否大于上手所出的牌
func (c *AnalyseCards) Follow(dType []byte) bool {
	var an AnalyseCards
	an.Analyse(dType)
	return c.FollowAnalyseCards(an)
}
func (c *AnalyseCards) FollowAnalyseCards(dType AnalyseCards) bool {
	if c.Len() != dType.Len() &&
		c.Len() != 4 &&
		c.Len() != 2 { //跟牌，但数量不符且不为炸弹
		return false
	}

	if c.GetKind() == UNKNOW { //所选牌不符合规定
		return false
	}

	if c.GetWeight() > dType.GetWeight() { // 牌型权重碾压
		return true
	}

	if c.GetWeight() < dType.GetWeight() { // 牌型权重没碾压
		return false
	}

	if dType.GetKind() != c.GetKind() { // 牌型不同
		return false
	}

	if dType.Len() != c.Len() { // 顺子牌型牌张要一致
		return false
	}
	return c.GetValue() > dType.GetValue() // 选牌不大于上家牌
}

func (c *AnalyseCards) Analyse(cards []byte) {
	kind, value := c.analyse(cards)
	c.SetValue(value)
	c.SetKind(kind)
}

func (c *AnalyseCards) analyse(cards [] byte) (uint8, byte) {
	c.Set(cards)

	length := c.Len()
	if length == 1 {
		kind, value := c.judgeDan()
		if kind > 0 {
			return kind, value
		}
	}
	if length == 2 {
		kind, value := c.judgePair()
		if kind > 0 {
			return kind, value
		}
		kind, value = c.judgeWangZha()
		if kind > 0 {
			return kind, value
		}
		return 0, 0
	}
	if length == 3 {
		return c.judge3Same()
	}
	if length == 4 {
		kind, value := c.judgeZhaDan()
		if kind > 0 {
			return kind, value
		}
		kind, value = c.judge3And1()
		if kind > 0 {
			return kind, value
		}
		return 0, 0
	}
	if length >= 5 && length <= 12 {
		kind, value := c.judgeShunZi(length)
		if kind > 0 {
			return kind, value
		}
	}
	if length >= 6 && length%2 == 0 {
		kind, value := c.judgeLianDui(length)
		if kind > 0 {
			return kind, value
		}
	}
	if length == 5 {
		kind, value := c.judge3AndPair()
		if kind > 0 {
			return kind, value
		}
	}
	if length == 6 {
		kind, value := c.judge4And2()
		if kind > 0 {
			return kind, value
		}
	}
	if length >= 6 && length%3 == 0 {
		kind, value := c.judgeFeiJi(length)
		if kind > 0 {
			return kind, value
		}
	}
	if length >= 8 && length%4 == 0 {
		kind, value := c.judgeFeiJiAndSingle(length)
		if kind > 0 {
			return kind, value
		}
	}
	if length >= 10 && length%5 == 0 {
		kind, value := c.judgeFeiJiAndPair(length)
		if kind > 0 {
			return kind, value
		}
	}
	if length == 8 {
		return c.judge4And2Pair()
	}
	return 0, 0
}

//单条
func (c *AnalyseCards) judgeDan() (uint8, byte) {
	if c.Get(c.End()) == 1 {
		return DAN_TIAO, c.End()
	}
	return 0, 0
}

// 对子
func (c *AnalyseCards) judgePair() (uint8, byte) {
	if c.Get(c.End()) == 2 {
		return DUI_ZI, c.End()
	}
	return 0, 0
}

func (c *AnalyseCards) judgeWangZha() (uint8, byte) {
	if c.Get(XIAO_WANG) == 1 && c.Get(DA_WANG) == 1 { // 王炸
		return WANG_ZHA, DA_WANG
	}
	return 0, 0
}

//三条
func (c *AnalyseCards) judge3Same() (uint8, byte) {
	if c.Get(c.End()) == 3 {
		return SAN_TIAO, c.End()
	}
	return 0, 0
}

//炸弹
func (c *AnalyseCards) judgeZhaDan() (uint8, byte) {
	if c.Get(c.End()) == 4 {
		return ZHA_DAN, c.End()
	}
	return 0, 0
}

//三带一
func (c *AnalyseCards) judge3And1() (uint8, byte) {
	start := c.Start()
	end := c.End()
	if end > TWO {
		end = TWO
	}
	for i := start; i <= end; i++ {
		count := c.Get(i)
		if count == 3 {
			return SAN_DAI_YI, i
		} else if count == 2 || count == 4 {
			break
		}
	}
	return 0, 0
}

//顺子
func (c *AnalyseCards) judgeShunZi(length uint8) (uint8, byte) {
	start := c.Start()
	end := c.End()
	if end > A {
		end = A
	}
	if ( end - start + 1) != length {
		return 0, 0
	}

	for i := start; i <= end; i++ {
		if c.Get(i) != 1 {
			return 0, 0
		}
	}

	return DAN_SHUN_ZI, end
}

//连对
func (c *AnalyseCards) judgeLianDui(length uint8) (uint8, byte) {
	start := c.Start()
	end := c.End()
	if end > A {
		end = A
	}
	if ( end-start+1)*2 != length {
		return 0, 0
	}

	for i := start; i <= end; i++ {
		if c.Get(i) != 2 {
			return 0, 0
		}
	}
	return LIAN_DUI, end
}

//三带一对 11122 22333
func (c *AnalyseCards) judge3AndPair() (uint8, byte) {
	var value byte
	var dui uint8
	start := c.Start()
	end := c.End()
	if end > TWO {
		end = TWO
	}
	for i := end; i >= start; i-- {
		count := c.Get(i)
		if count == 3 {
			value = i
		} else if count == 2 {
			dui ++
		} else if count > 0 {
			break
		}
		if dui == 1 && value > 0 {
			return SAN_DAI_YI_DUI, value
		}
	}
	return 0, 0
}

//四带两个 111123 233334 456666
func (c *AnalyseCards) judge4And2() (uint8, byte) {

	start := c.Start()
	end := c.End()
	if end > TWO {
		end = TWO
	}
	for i := end; i >= start; i-- {
		count := c.Get(i)
		if count == 3 {
			break
		} else if count == 4 {
			return SI_DAI_DAN, i
		}
	}
	return 0, 0
}

//三顺 333444555
func (c *AnalyseCards) judgeFeiJi(length uint8) (uint8, byte) {
	start := c.Start()
	end := c.End()
	if end > A {
		end = A
	}
	if ( end-start+1)*3 != length {
		return 0, 0
	}
	for i := start; i <= end; i++ {
		if c.Get(i) != 3 {
			return 0, 0
		}
	}

	return SAN_SHUN_ZI, end
}

//飞机带单
func (c *AnalyseCards) judgeFeiJiAndSingle(length uint8) (uint8, byte) {
	var num uint8
	var value byte
	start := c.Start()
	end := c.End()
	if end > A {
		end = A
	}
	for i := start; i <= end; i++ {
		count := c.Get(i)
		if count >= 3 {
			if value > 0 && value+1 != i {
				break
			}
			num ++
			value = i
		}
		if length == num*4 {
			return FEI_JI_DAI_DAN, value
		}
	}
	return 0, 0
}

//飞机带对
func (c *AnalyseCards) judgeFeiJiAndPair(length uint8) (uint8, byte) {
	var three uint8
	var dui uint8
	var value byte
	start := c.Start()
	end := c.End()
	if end > TWO {
		end = TWO
	}
	for i := start; i <= end; i++ {
		count := c.Get(i)
		if count == 3 && i != TWO {
			if value > 0 && value+1 != i {
				break
			}
			three ++
			value = i
		} else if count == 2 {
			dui ++
		} else if count == 4 {
			dui += 2
		}
		if length == three*3+dui*2 {
			return FEI_JI_DAI_DUI, value
		}
	}
	return 0, 0
}

//四带两对  满足 11222233(2-4-2) 11112233(4-2-2) 11223333(2-2-4) 11112222(2-2-4)
func (c *AnalyseCards) judge4And2Pair() (uint8, byte) {
	var value byte
	var dui uint8
	start := c.Start()
	end := c.End()
	if end > TWO {
		end = TWO
	}
	for i := end; i >= start; i-- {
		count := c.Get(i)
		if count == 4 && value == 0 {
			value = i
		} else if count == 4 {
			dui += 2
		} else if count == 2 {
			dui ++
		} else if count > 0 {
			break
		}
		if dui == 2 && value > 0 {
			return SI_DAI_DUI, value
		}
	}
	return 0, 0
}

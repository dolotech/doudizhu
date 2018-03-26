package algorithm

import (
	"testing"
)



func TestIsFollowSanTiao(t *testing.T) {
	// 三条
	var an AnalyseCards
	an.Analyse([]byte{0x02, 0x12, 0x22})
	an.judge3Same()
	t.Logf("三条 %v,%v，%v", an.GetKind() == SAN_TIAO, an.GetWeight() == 1, an.GetValue() == GetValue(0x02))

}
func TestIsFollow(t *testing.T) {
	//cardType := &Cards{}
	var an AnalyseCards
	an.Analyse([]byte{0xFF, 0xEF})
	t.Logf("能否跟上上手牌 %v ", an.Follow([]byte{0x02, 0x12, 0x22, 0x32}) == true)


	an.Analyse([]byte{0x02, 0x12, 0x22, 0x32})
	t.Logf("能否跟上上手牌 %v ", an.Follow([]byte{0xFF, 0xEF}) == false)


	an.Analyse([]byte{0x03, 0x13, 0x23, 0x33})

	t.Logf("能否跟上上手牌 %v ", an.Follow([]byte{0x02, 0x12, 0x22, 0x32}) == false)


	an.Analyse([]byte{0x04, 0x05, 0x06, 0x07, 0x08})
	t.Logf("能否跟上上手牌 %v ", an.Follow([]byte{0x03, 0x04, 0x05, 0x06, 0x07}) == true)

}


// 顺子
func TestAnalyseShunzi(t *testing.T) {
	cardType := BuildCards([]byte{0x0b, 0x0c, 0x2d, 0x1e, 0xFF})
	t.Logf("顺子 %v,%v，%v %v，%v %v", cardType.Kind() != DAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0, cardType.Kind(), cardType.Weight(), cardType.Value())

	//t.Log(a,b,c, ConvertValue([]byte{0x0b, 0x0c, 0x2d, 0x1e, 0xFF}))
	cardType = BuildCards([]byte{0x02, 0x03, 0x24, 0x15, 0x36})
	t.Logf("顺子 %v,%v，%v", cardType.Kind() != DAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x0b, 0x0c, 0x2d, 0x1e, 0x02})

	t.Logf("顺子 %v,%v，%v", cardType.Kind() != DAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)
}

// 三顺
func TestAnalyseSanshun(t *testing.T) {
	cardType := BuildCards([]byte{0x12, 0x02, 0x12, 0x28, 0x08, 0x18})
	t.Logf("三顺 %v,%v，%v", cardType.Kind() != SAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x12, 0x02, 0x12, 0x23, 0x03, 0x13})
	t.Logf("三顺 %v,%v，%v", cardType.Kind() != SAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x23, 0x03, 0x13, 0x14, 0x04, 0x14})
	t.Logf("三顺 %v,%v，%v", cardType.Kind() == SAN_SHUN_ZI, cardType.Weight() == 1, cardType.Value() == 4)

	cardType = BuildCards([]byte{0x23, 0x03, 0x14, 0x15, 0x05, 0x16})
	t.Logf("三顺 %v,%v，%v", cardType.Kind() != SAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x2d, 0x0d, 0x1d, 0x1e, 0x0e, 0x1e})
	t.Logf("三顺 %v,%v，%v", cardType.Kind() == SAN_SHUN_ZI, cardType.Weight() == 1, cardType.Value() == 0xe)

	cardType = BuildCards([]byte{0x1e, 0x1e, 0x1e, 0x02, 0x02, 0x02})
	t.Logf("三顺 %v,%v，%v", cardType.Kind() != SAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x14, 0x24, 0x04, 0x05, 0x15, 0x25, 0x03, 0x13, 0x23, 0x06, 0x16, 0x26})
	t.Logf("三顺 %v,%v，%v", cardType.Kind() == SAN_SHUN_ZI, cardType.Weight() == 1, cardType.Value() == 6)
}

//飞机带两个
func TestAnalyseFeijidaidan(t *testing.T) {

	cardType := BuildCards([]byte{0x02, 0x12, 0x22, 0x05, 0x15, 0x25, 0x26, 0x27})
	t.Logf("飞机带两个 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DAN, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x02, 0x12, 0x22, 0x03, 0x13, 0x23, 0x26, 0x27})
	t.Logf("飞机带两个 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DAN, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x26, 0x27})
	t.Logf("飞机带两个 %v,%v，%v", cardType.Kind() == FEI_JI_DAI_DAN, cardType.Weight() == 1, cardType.Value() == 0x5)

	cardType = BuildCards([]byte{0x1e, 0x1e, 0x1e, 0x02, 0x02, 0x02, 0x27, 0x27})
	t.Logf("飞机带两个 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DAN, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x1e, 0x0e, 0x2e, 0x3e, 0x0d, 0x0d, 0x0d, 0x0d})
	t.Logf("飞机带两个 %v,%v，%v", cardType.Kind() == FEI_JI_DAI_DAN, cardType.Weight() == 1, cardType.Value() == 0x0e)

	cardType = BuildCards([]byte{0x14, 0x24, 0x04, 0x05, 0x15, 0x25, 0x03, 0x13, 0x23, 0x06, 0x16, 0x26})
	t.Logf("飞机带两个 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DAN, cardType.Weight() == 1, cardType.Value() == 6)
}

//飞机带对
func TestAnalyseFeijidaidui(t *testing.T) {
	//飞机带对
	cardType := BuildCards([]byte{0x03, 0x13, 0x23, 0x05, 0x15, 0x25, 0x26, 0x36, 0x27, 0x37})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	//飞机带对
	cardType = BuildCards([]byte{0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x26, 0x36, 0x27, 0x37})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() == FEI_JI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 0x5)

	//飞机带对
	cardType = BuildCards([]byte{0x02, 0x12, 0x22, 0x03, 0x13, 0x23, 0x26, 0x36, 0x27, 0x37})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	//飞机带对
	cardType = BuildCards([]byte{0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x26, 0x36, 0xEF, 0xFF})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	//飞机带对
	cardType = BuildCards([]byte{0x02, 0x02, 0x03, 0x13, 0x23, 0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x26, 0x36, 0x27, 0x37})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() == FEI_JI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 5)

	cardType = BuildCards([]byte{0x1e, 0x1e, 0x1e, 0x02, 0x02, 0x02, 0x27, 0x27, 0x27, 0x27})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x1e, 0x1e, 0x1e, 0x02, 0x02, 0x02, 0x26, 0x26, 0x27, 0x27})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x1e, 0x1e, 0x1e, 0x0d, 0x0d, 0x0d, 0x27, 0x27, 0x27, 0x27})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() == FEI_JI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 0x0e)

	cardType = BuildCards([]byte{0x03, 0x03, 0x03, 0x04, 0x04, 0x04, 0x05, 0x05, 0x05, 0x05})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() == FEI_JI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 0x04)
}

//四带两对
func TestAnalyseSidaiDui(t *testing.T) {
	cardType := BuildCards([]byte{0x38, 0x08, 0x18, 0x28, 0xEF, 0xFF, 0x1a, 0x2a})
	t.Logf("四带两对 %v,%v，%v", cardType.Kind() != SI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0xFF, 0x08, 0x18, 0x28, 0x38, 0xEF, 0x1e, 0x2e})
	t.Logf("四带两对 %v,%v，%v", cardType.Kind() != SI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x1e, 0x0e, 0x2e, 0x3e, 0x0d, 0x0d, 0x0d, 0x0d})
	t.Logf("四带两对 %v,%v，%v", cardType.Kind() != SI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 0x0e)

	cardType = BuildCards([]byte{0x1e, 0x0e, 0x2e, 0x3e, 0x03, 0x03, 0x0d, 0x0d})
	t.Logf("四带两对 %v,%v，%v", cardType.Kind() == SI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 0x0e)

	cardType = BuildCards([]byte{0x18, 0x8, 0x38, 0x28, 0x5, 0x15, 0x6, 0x16})
	t.Logf("四带两对 %v,%v，%v", cardType.Kind() == SI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 0x08)

	cardType = BuildCards([]byte{0x23, 0x13, 0x3, 0x33, 0x24, 0x14, 0x29, 0x19})
	t.Logf("四带两对 %v,%v，%v", cardType.Kind() == SI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 0x03)
}

//三带一对
func TestAnalyseSanDaidui(t *testing.T) {

	cardType := BuildCards([]byte{0x1a, 0x08, 0x18, 0x28, 0x0a})
	t.Logf("三带一对 %v,%v，%v", cardType.Kind() == SAN_DAI_YI_DUI, cardType.Weight() == 1, cardType.Value() == 0x8)

	cardType = BuildCards([]byte{0xEF, 0x08, 0x18, 0x28, 0xFF})
	t.Logf("三带一对 %v,%v，%v", cardType.Kind() != SAN_DAI_YI_DUI, cardType.Weight() == 0, cardType.Value() == 0)
}

//连对
func TestAnalyseLiandui(t *testing.T) {
	cardType := BuildCards([]byte{0x23, 0x13, 0x14, 0x04, 0x05, 0x15, 0x06, 0x36})
	t.Logf("连对 %v,%v，%v", cardType.Kind() == LIAN_DUI, cardType.Weight() == 1, cardType.Value() == 0x6)

	cardType = BuildCards([]byte{0x23, 0x13, 0x14, 0x04, 0x05, 0x15, 0x07, 0x37})
	t.Logf("连对 %v,%v，%v", cardType.Kind() != LIAN_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x0b, 0x0c, 0x2d, 0x1e, 0xff, 0x0b, 0x0c, 0x2d, 0x1e, 0xff})
	t.Logf("连对 %v,%v，%v", cardType.Kind() != LIAN_DUI, cardType.Weight() == 0, cardType.Value() == 0)
}

func TestAnalyse(t *testing.T) {

	// 顺子
	cardType := BuildCards([]byte{0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E})
	t.Logf("顺子 %v,%v，%v", cardType.Kind() != DAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)

	// 顺子
	cardType = BuildCards([]byte{0x03, 0x04, 0x25, 0x16, 0x02})
	t.Logf("顺子 %v,%v，%v", cardType.Kind() != DAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)

	// 顺子
	cardType = BuildCards([]byte{0x0b, 0x0c, 0x2d, 0x1e, 0xFF})
	t.Logf("顺子 %v,%v，%v", cardType.Kind() != DAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)

	// 王炸
	cardType = BuildCards([]byte{0xFF, 0xEF})
	t.Logf("王炸 %v,%v，%v", cardType.Kind() == WANG_ZHA, cardType.Weight() == 3, cardType.Value() == GetValue(0xFF))

	// 单张
	cardType = BuildCards([]byte{0x02})
	t.Logf("单张 %v,%v，%v", cardType.Kind() == DAN_TIAO, cardType.Weight() == 1, cardType.Value() == GetValue(0x02))

	// 对子
	cardType = BuildCards([]byte{0x02, 0x12, 0x22})
	//t.Logf("对子 %v,%v，%v", cardType.Kind(),b,c)
	t.Logf("对子 %v,%v，%v", cardType.Kind() != DUI_ZI, cardType.Weight() == 1, cardType.Value() == GetValue(0x02))

	// 三条
	cardType = BuildCards([]byte{0x02, 0x12, 0x22})
	t.Logf("三条 %v,%v，%v", cardType.Kind() == SAN_TIAO, cardType.Weight() == 1, cardType.Value() == GetValue(0x02))

	// 三带单
	cardType = BuildCards([]byte{0x03, 0x02, 0x12, 0x22})
	t.Logf("三带单 %v,%v，%v", cardType.Kind() == SAN_DAI_YI, cardType.Weight() == 1, cardType.Value() == GetValue(0x02))

	// 炸弹
	cardType = BuildCards([]byte{0x02, 0x12, 0x22, 0x32})
	t.Logf("炸弹 %v,%v，%v", cardType.Kind() == ZHA_DAN, cardType.Weight() == 2, cardType.Value() == GetValue(0x02))

	// 炸弹
	cardType = BuildCards([]byte{0x03, 0x13, 0x23, 0x33})
	t.Logf("炸弹 %v,%v，%v", cardType.Kind() == ZHA_DAN, cardType.Weight() == 2, cardType.Value() == 3)

	//连对
	cardType = BuildCards([]byte{0x23, 0x04, 0x05, 0x06, 0x13, 0x14, 0x15, 0x36})
	t.Logf("连对 %v,%v，%v", cardType.Kind() == LIAN_DUI, cardType.Weight() == 1, cardType.Value() == 6)

	//连对
	cardType = BuildCards([]byte{0x0b, 0x0c, 0x2d, 0x1e, 0xff, 0x0b, 0x0c, 0x2d, 0x1e, 0xff})
	//t.Logf("连对 %v,%v，%v", cardType.Kind(), b, c)
	t.Logf("连对 %v,%v，%v", cardType.Kind() != LIAN_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	//三带一对
	cardType = BuildCards([]byte{0x17, 0x02, 0x12, 0x22, 0x07})
	t.Logf("三带一对 %v,%v，%v", cardType.Kind() == SAN_DAI_YI_DUI, cardType.Weight() == 1, cardType.Value() == GetValue(0x02))

	//飞机带单
	cardType = BuildCards([]byte{0x02, 0x12, 0x22, 0x05, 0x15, 0x25, 0x26, 0x27})
	t.Logf("飞机带单 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DAN, cardType.Weight() == 0, cardType.Value() == 0)

	//飞机带单
	cardType = BuildCards([]byte{0x02, 0x12, 0x22, 0x03, 0x13, 0x23, 0x26, 0x27})
	t.Logf("飞机带单 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DAN, cardType.Weight() == 0, cardType.Value() == 0)

	//飞机带单
	cardType = BuildCards([]byte{0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x26, 0x27})
	t.Logf("飞机带单 %v,%v，%v", cardType.Kind() == FEI_JI_DAI_DAN, cardType.Weight() == 1, cardType.Value() == 5)

	//飞机带对
	cardType = BuildCards([]byte{0x03, 0x13, 0x23, 0x05, 0x15, 0x25, 0x26, 0x36, 0x27, 0x37})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	//飞机带对
	cardType = BuildCards([]byte{0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x26, 0x36, 0x27, 0x37})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() == FEI_JI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 5)

	//飞机带对
	cardType = BuildCards([]byte{0x02, 0x12, 0x22, 0x03, 0x13, 0x23, 0x26, 0x36, 0x27, 0x37})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	//四带单
	cardType = BuildCards([]byte{0x02, 0x12, 0x22, 0x32, 0x06, 0x16})
	t.Logf("四带单 %v,%v，%v", cardType.Kind() == SI_DAI_DAN, cardType.Weight() == 1, cardType.Value() == GetValue(0x02))

	//四带单
	cardType = BuildCards([]byte{0x06, 0x16, 0x07, 0x17, 0x27, 0x37})
	t.Logf("四带单 %v,%v，%v", cardType.Kind() == SI_DAI_DAN, cardType.Weight() == 1, cardType.Value() == 7)

	//四带两对
	cardType = BuildCards([]byte{0x02, 0x12, 0x22, 0x32, 0x06, 0x36, 0x0a, 0x3a})
	t.Logf("四带两对 %v,%v，%v", cardType.Kind() == SI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == GetValue(0x02))

	cardType = BuildCards([]byte{0x0b, 0x0c, 0x2d, 0x1e, 0xFF})
	t.Logf("顺子 %v,%v，%v", cardType.Kind() != DAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)

	//t.Log(a,b,c, ConvertValue([]byte{0x0b, 0x0c, 0x2d, 0x1e, 0xFF}))
	cardType = BuildCards([]byte{0x02, 0x03, 0x24, 0x15, 0x36})
	t.Logf("顺子 %v,%v，%v", cardType.Kind() != DAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)

	// 三顺
	cardType = BuildCards([]byte{0x12, 0x02, 0x12, 0x28, 0x08, 0x18})
	t.Logf("三顺 %v,%v，%v", cardType.Kind() != SAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x12, 0x02, 0x12, 0x23, 0x03, 0x13})
	t.Logf("三顺 %v,%v，%v", cardType.Kind() != SAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x23, 0x03, 0x13, 0x14, 0x04, 0x14})
	t.Logf("三顺 %v,%v，%v", cardType.Kind() == SAN_SHUN_ZI, cardType.Weight() == 1, cardType.Value() == 4)

	cardType = BuildCards([]byte{0x23, 0x03, 0x14, 0x15, 0x05, 0x16})
	t.Logf("三顺 %v,%v，%v", cardType.Kind() != SAN_SHUN_ZI, cardType.Weight() == 0, cardType.Value() )

	cardType = BuildCards([]byte{0x2d, 0x0d, 0x1d, 0x1e, 0x0e, 0x1e})
	t.Logf("三顺 %v,%v，%v", cardType.Kind() == SAN_SHUN_ZI, cardType.Weight() == 1, cardType.Value() == 0xe)

	//飞机带两个

	cardType = BuildCards([]byte{0x02, 0x12, 0x22, 0x05, 0x15, 0x25, 0x26, 0x27})
	t.Logf("飞机带两个 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DAN, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x02, 0x12, 0x22, 0x03, 0x13, 0x23, 0x26, 0x27})
	t.Logf("飞机带两个 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DAN, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x26, 0x27})
	t.Logf("飞机带两个 %v,%v，%v", cardType.Kind() == FEI_JI_DAI_DAN, cardType.Weight() == 1, cardType.Value() == 0x5)

	//飞机带对
	//飞机带对
	cardType = BuildCards([]byte{0x03, 0x13, 0x23, 0x05, 0x15, 0x25, 0x26, 0x36, 0x27, 0x37})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	//飞机带对
	cardType = BuildCards([]byte{0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x26, 0x36, 0x27, 0x37})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() == FEI_JI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 0x5)

	//飞机带对
	cardType = BuildCards([]byte{0x02, 0x12, 0x22, 0x03, 0x13, 0x23, 0x26, 0x36, 0x27, 0x37})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	//飞机带对
	cardType = BuildCards([]byte{0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x26, 0x36, 0xEF, 0xFF})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() != FEI_JI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	//飞机带对
	cardType = BuildCards([]byte{0x02, 0x02, 0x03, 0x13, 0x23, 0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x26, 0x36, 0x27, 0x37})
	t.Logf("飞机带对 %v,%v，%v", cardType.Kind() == FEI_JI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 5)

	//四带两对
	cardType = BuildCards([]byte{0x38, 0x08, 0x18, 0x28, 0xEF, 0xFF, 0x1a, 0x2a})
	t.Logf("四带两对 %v,%v，%v", cardType.Kind() != SI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0xFF, 0x08, 0x18, 0x28, 0x38, 0xEF, 0x1e, 0x2e})
	t.Logf("四带两对 %v,%v，%v", cardType.Kind() != SI_DAI_DUI, cardType.Weight() == 0, cardType.Value() == 0)

	cardType = BuildCards([]byte{0x06, 0x16, 0x26, 0x07, 0x17, 0x27, 0x36, 0x37})
	t.Logf("四带两对 %v,%v，%v", cardType.Kind() != SI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 0x7)



	cardType = BuildCards([]byte{0x06, 0x16, 0x36, 0x08, 0x18, 0x28, 0x26, 0x38})
	t.Logf("四带两对 %v,%v，%v", cardType.Kind() == SI_DAI_DUI, cardType.Weight() == 1, cardType.Value() == 0x8)

	//三带一对
	cardType = BuildCards([]byte{0x1a, 0x08, 0x18, 0x28, 0x0a})
	t.Logf("三带一对 %v,%v，%v", cardType.Kind() == SAN_DAI_YI_DUI, cardType.Weight() == 1, cardType.Value() == 0x8)

	cardType = BuildCards([]byte{0xEF, 0x08, 0x18, 0x28, 0xFF})
	t.Logf("三带一对 %v,%v，%v", cardType.Kind() != SAN_DAI_YI_DUI, cardType.Weight() == 0, cardType.Value() == 0)
	//连对
	cardType = BuildCards([]byte{0x23, 0x13, 0x14, 0x04, 0x05, 0x15, 0x06, 0x36})
	t.Logf("连对 %v,%v，%v", cardType.Kind() == LIAN_DUI, cardType.Weight() == 1, cardType.Value() == 0x6)

	cardType = BuildCards([]byte{0x23, 0x13, 0x14, 0x04, 0x05, 0x15, 0x07, 0x37})
	t.Logf("连对 %v,%v，%v", cardType.Kind() != LIAN_DUI, cardType.Weight() == 0, cardType.Value() == 0,)

	cardType = BuildCards([]byte{0x0b, 0x0c, 0x2d, 0x1e, 0xff, 0x0b, 0x0c, 0x2d, 0x1e, 0xff})
	t.Logf("连对 %v,%v，%v", cardType.Kind() != LIAN_DUI, cardType.Weight() == 0, cardType.Value() == 0)

}

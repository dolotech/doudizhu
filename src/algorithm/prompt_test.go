package algorithm

import "testing"

func TestPromptFeijiDaiDui(t *testing.T) {

	//handcards := []byte{0x03, 0x05, 0x15, 0x25, 0x35, 0x06, 0x16, 0x26, 0x36, 0x08, 0x38, 0x27, 0x18, 0x28, 0x2D}
	//discards := []byte{0x23, 0x13, 0x03, 0x04, 0x24, 0x14, 0x29, 0x19, 0x29, 0x19}
	//cardType := SelectCards(handcards, discards)
	//
	//t.Logf("%#+v 飞机带对 %#+v %#+v ", yapaitest(handcards, discards, 4, cardType), handcards, discards)
	//

	handcards := []byte{0x06, 0x16, 0x26, 0x09, 0x19, 0x29, 0x27, 0x37, 0x07, 0x18, 0x29, 0x08, 0x38, 0x2D, 0x3D}
	discards := []byte{0x23, 0x13, 0x03, 0x04, 0x24, 0x14, 0x15, 0x25, 0x35, 0x2e, 0x1e, 0x2e, 0x1e, 0x2a, 0x1a}
	ty := &Cards{}

	ty.AnalyseUnSort(handcards)

	cardTypes := ty.Prompt(discards)

	for _, v := range cardTypes {
		t.Logf("%#+v 飞机带对 ", v.Cards)
		/*for i:= THREE;i<=DA_WANG;i++{
			t.Logf("%#v %v ",i,v.analyseCards.Get(i))
		}*/
	}

	t.Logf("%#+v 飞机带对 %#+v %#+v ", yapaitest(handcards, discards, 2, cardTypes), handcards, discards)
}

func TestPromptFeijiDaiDan(t *testing.T) {

	handcards := []byte{0x03, 0x05, 0x15, 0x25, 0x35, 0x06, 0x16, 0x26, 0x36, 0x18, 0x38, 0x27, 0x08, 0x28, 0x2D}
	discards := []byte{0x23, 0x13, 0x03, 0x33, 0x24, 0x14, 0x04, 0x19}
	ty := &Cards{}

	ty.AnalyseUnSort(handcards)

	cardTypes := ty.Prompt(discards)
	for _, v := range cardTypes {
		t.Logf("%#+v 飞机带单 ", v.Cards)
	}

	t.Logf("%#+v 飞机带单 %#+v %#+v ", yapaitest(handcards, discards, 4, cardTypes), handcards, discards)

	handcards = []byte{0x02, 0x15, 0x25, 0x35, 0x06, 0x16, 0x26, 0x36}
	discards = []byte{0x23, 0x13, 0x03, 0x33, 0x24, 0x14, 0x04, 0x19}
	ty = &Cards{}

	ty.AnalyseUnSort(handcards)

	cardTypes = ty.Prompt(discards)
	for _, v := range cardTypes {
		t.Logf("%#+v 飞机带单 ", v.Cards)
	}

	t.Logf("%#+v 飞机带单 %#+v %#+v ", yapaitest(handcards, discards, 2, cardTypes), handcards, discards)
}

func TestPromptSiDaiDui(t *testing.T) {
	handcards := []byte{0x03, 0x05, 0x15, 0x25, 0x35, 0x06, 0x16, 0x26, 0x36, 0x18, 0x08, 0x27, 0x38, 0x28, 0x2D}
	discards := []byte{0x23, 0x13, 0x03, 0x33, 0x24, 0x14, 0x29, 0x19}
	ty := &Cards{}
	ty.AnalyseUnSort(handcards)
	/*cardTypes := ty.Prompt(discards)
	for _,v:=range  cardTypes{
		t.Logf("%#+v 四带对 ",v.Cards)
	}
	for k,v:=range ty.analyseCards{
		t.Log(len(v),k,v)
	}
	t.Logf("%#+v 四带对 %#+v %#+v ", yapaitest(handcards, discards, 6, cardTypes), handcards, discards)*/

	handcards = []byte{0x18, 0x8, 0x38, 0x28, 0x5, 0x15, 0x6, 0x16}
	discards = []byte{0x23, 0x13, 0x3, 0x33, 0x24, 0x14, 0x29, 0x19}
	ty = &Cards{}

	ty.AnalyseUnSort(handcards)

	cardTypes := ty.Prompt(discards)

	for _, v := range cardTypes {
		t.Logf("%#+v 四带对 %#+v", v.Cards, v.Kind)
	}

	t.Logf("%#+v 四带对 %#+v %#+v ", yapaitest(handcards, discards, 2, cardTypes), handcards, discards)
}

func TestPromptSiDaiDan(t *testing.T) {

	handcards := []byte{0x03, 0x05, 0x15, 0x25, 0x35, 0x06, 0x16, 0x26, 0x36, 0x18, 0x38, 0x27, 0x08, 0x28, 0x2D}
	discards := []byte{0x23, 0x13, 0x03, 0x33, 0x24, 0x17}
	ty := &Cards{}

	ty.AnalyseUnSort(handcards)

	cardTypes := ty.Prompt(discards)
	for _, v := range cardTypes {
		t.Logf("%#+v 四带单 ", v.Cards)
	}

	t.Logf("%#+v 四带单 %#+v %#+v ", yapaitest(handcards, discards, 6, cardTypes), handcards, discards)
}

func TestPromptSanShun(t *testing.T) {

	handcards := []byte{0x05, 0x15, 0x25, 0x16, 0x26, 0x36, 0x18, 0x27, 0x18, 0x28, 0x2D}
	discards := []byte{0x23, 0x13, 0x03, 0x24, 0x14, 0x04}
	ty := &Cards{}

	ty.AnalyseUnSort(handcards)

	cardTypes := ty.Prompt(discards)
	for _, v := range cardTypes {
		t.Logf("%#+v 三顺 ", v.Cards)
	}

	t.Logf("%#+v 三顺 %#+v %#+v ", yapaitest(handcards, discards, 1, cardTypes), handcards, discards)
}

func TestPromptPromptDanTiao(t *testing.T) {
	handcards := []byte{0xEF, 0xFF, 0xFF, 0xFF, 0xFF, 0x02, 0x12, 0x22, 0x32, 0x9, 0x19, 0x23}

	var an AnalyseCards
	an.Set(handcards)

	//t.Logf("%b", an)

	an.SetValue(17)
	an.SetKind(14)
	//t.Logf("%b", an)
	//t.Logf("%v %v", an.GetKind(), an.GetValue())
	arr := make([]AnalyseCards, 0)
	an.PromptDanTiao(0x4, &arr)
	t.Logf("%v %b", len(arr) == 4, an)

	for _, v := range arr {
		t.Logf("%v %#v", v.GetKind(), v.GetValue())
	}

}
func TestPromptLianDui(t *testing.T) {

	handcards := []byte{0x25, 0x25, 0x26, 0x26, 0x17, 0x27, 0x18, 0x28, 0x2D}
	discards := []byte{0x04, 0x14, 0x05, 0x15, 0x06, 0x16}
	ty := &Cards{}

	ty.AnalyseUnSort(handcards)

	cardTypes := ty.Prompt(discards)
	for _, v := range cardTypes {
		t.Logf("%#+v 连对 ", v.Cards)
	}

	t.Logf("%#+v 连对 %#+v %#+v ", yapaitest(handcards, discards, 2, cardTypes), handcards, discards)

	handcards = []byte{0x15, 0x05, 0x16, 0x26, 0x18, 0x27, 0x08, 0x28, 0x2D}
	discards = []byte{0x04, 0x14, 0x05, 0x15, 0x06, 0x16}
	ty = &Cards{}

	ty.AnalyseUnSort(handcards)

	cardTypes = ty.Prompt(discards)
	for _, v := range cardTypes {
		t.Logf("%#+v 连对 ", v.Cards)
	}

	rel := yapaitest(handcards, discards, 0, cardTypes)
	t.Logf("%#+v 连对 %#+v %#+v ", rel, handcards, discards)
}

func TestPromptShunzi(t *testing.T) {

	handcards := []byte{0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x2D}
	discards := []byte{0x04, 0x05, 0x06, 0x07, 0x08}
	ty := &Cards{}

	ty.AnalyseUnSort(handcards)

	cardTypes := ty.Prompt(discards)
	for _, v := range cardTypes {
		t.Logf("%#+v 顺子 ", v.Cards)
	}
	rel := yapaitest(handcards, discards, 5, cardTypes)
	t.Logf("%#+v 顺子 %#+v %#+v ", rel, handcards, discards)

	handcards = []byte{0x03, 0x04, 0x25, 0x15, 0x16, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x2D}
	discards = []byte{0x03, 0x04, 0x05, 0x06, 0x07, 0x08}
	ty = &Cards{}

	ty.AnalyseUnSort(handcards)

	cardTypes = ty.Prompt(discards)

	for _, v := range cardTypes {
		t.Logf("%#+v 顺子 ", v.Cards)
	}
	rel = yapaitest(handcards, discards, 1, cardTypes)
	t.Logf("%#+v 顺子 %#+v %#+v ", rel, handcards, discards)

	handcards = []byte{0x03, 0x04, 0x15, 0x25, 0x16, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x2D}
	discards = []byte{0x03, 0x04, 0x05, 0x06, 0x07}
	ty = &Cards{}

	ty.AnalyseUnSort(handcards)

	cardTypes = ty.Prompt(discards)

	for _, v := range cardTypes {
		t.Logf("%#+v 顺子 ", v.Cards)
	}
	rel = yapaitest(handcards, discards, 2, cardTypes)
	t.Logf("%#+v 顺子 %#+v %#+v ", rel, handcards, discards)
}

func TestPrompt(t *testing.T) {
	handcards := []byte{0xEF, 0xFF, 0x02, 0x12, 0x22, 0x32}
	discards := []byte{0x04, 0x14}
	ty := &Cards{Cards: handcards}
	ty.AnalyseUnSort(handcards)
	cardTypes := ty.Prompt(discards)
	rel := yapaitest(handcards, discards, 3, cardTypes)
	t.Logf("%#+v 对子 %#+v %#+v ", rel, handcards, discards)

	handcards = []byte{0xEF, 0xFF, 0x02, 0x12, 0x22, 0x32}
	discards = []byte{0x04, 0x14, 0x24, 0x34}
	ty = &Cards{Cards: handcards}
	ty.AnalyseUnSort(handcards)
	cardTypes = ty.Prompt(discards)
	rel = yapaitest(handcards, discards, 2, cardTypes)
	t.Logf("%#+v 炸弹 %#+v %#+v ", rel, handcards, discards)

	handcards = []byte{0xEF, 0xFF, 0x02, 0x12, 0x22, 0x32, 0x09, 0x19, 0x29, 0x39}
	discards = []byte{0x04, 0x14, 0x24, 0x34}
	ty = &Cards{}
	ty.AnalyseUnSort(handcards)
	cardTypes = ty.Prompt(discards)
	rel = yapaitest(handcards, discards, 3, cardTypes)
	t.Logf("%#+v 炸弹 %#+v %#+v %+v ", rel, handcards, discards, len(cardTypes))

	handcards = []byte{0xEF, 0xFF, 0x02, 0x12, 0x22, 0x32}
	discards = []byte{0x04, 0x14, 0x24}
	ty = &Cards{Cards: handcards}
	ty.AnalyseUnSort(handcards)
	cardTypes = ty.Prompt(discards)
	rel = yapaitest(handcards, discards, 3, cardTypes)
	t.Logf("%#+v 三条 %#+v %#+v ", rel, handcards, discards)

	handcards = []byte{0xEF, 0xFF, 0x02, 0x12, 0x22, 0x32, 0x9, 0x19, 0x29}
	discards = []byte{0x04, 0x14}
	ty = &Cards{Cards: handcards}
	ty.AnalyseUnSort(handcards)
	cardTypes = ty.Prompt(discards)
	rel = yapaitest(handcards, discards, 4, cardTypes)

	for _, v := range cardTypes {
		t.Logf("%#+v 对子 ", v.Cards)
	}

	t.Logf("%#+v 对子 %#+v %#+v ", rel, handcards, discards)

	handcards = []byte{0xEF, 0xFF, 0x02, 0x12, 0x22, 0x32, 0x9, 0x19, 0x29}
	discards = []byte{0x04}
	ty = &Cards{}
	ty.AnalyseUnSort(handcards)
	cardTypes = ty.Prompt(discards)

	for _, v := range cardTypes {
		t.Logf("cards: %#v %#v", v.Cards, v.Kind)
	}

	rel = yapaitest(handcards, discards, 6, cardTypes)
	t.Logf("%#+v 单条 %#+v %#+v ", rel, handcards, discards)

	handcards = []byte{0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x2D}
	discards = []byte{0x04, 0x05, 0x06, 0x07, 0x08}
	ty = &Cards{}
	ty.AnalyseUnSort(handcards)
	cardTypes = ty.Prompt(discards)
	rel = yapaitest(handcards, discards, 5, cardTypes)
	t.Logf("%#+v 顺子 %#+v %#+v ", rel, handcards, discards)
}

func TestPromptSanDaiDan(t *testing.T) {
	handcards := []byte{0x03, 0x05, 0x15, 0x25, 0x16, 0x26, 0x36, 0x18, 0x27, 0x28, 0x08, 0x2D}
	discards := []byte{0x23, 0x13, 0x03, 0x24}

	ty := &Cards{}
	ty.AnalyseUnSort(handcards)
	cardTypes := ty.Prompt(discards)

	for _, v := range cardTypes {
		t.Logf("%#+v 三带单 ", v.Cards)
	}
	rel := yapaitest(handcards, discards, 3, cardTypes)
	t.Logf("%#+v 三带单 %#+v %#+v ", rel, handcards, discards)
}

func TestPrompSanDaiDui(t *testing.T) {

	handcards := []byte{0x03, 0x05, 0x15, 0x25, 0x16, 0x26, 0x36, 0x18, 0x27, 0x08, 0x28, 0x2D}
	discards := []byte{0x23, 0x13, 0x03, 0x24, 0x24}

	ty := &Cards{}
	ty.AnalyseUnSort(handcards)
	cardTypes := ty.Prompt(discards)

	for _, v := range cardTypes {
		t.Logf("%#+v 三带对 ", v.Cards)
	}

	rel := yapaitest(handcards, discards, 3, cardTypes)
	t.Logf("%#+v 三带对 %#+v %#+v ", rel, handcards, discards)
}

func TestPrompDanTiao4(t *testing.T) {
	handcards := []byte{0xEF, 0xFF, 0x02, 0x12, 0x22, 0x32, 0x9, 0x19, 0x29}
	discards := []byte{0x04}
	ty := &Cards{}
	ty.AnalyseUnSort(handcards)
	cardTypes := ty.Prompt(discards)

	for _, v := range cardTypes {
		t.Logf("cards: %#v %v", v.Cards, v.Kind)
	}

	rel := yapaitest(handcards, discards, 6, cardTypes)
	t.Logf("%#+v 单条 %#+v %#+v ", rel, handcards, discards)
}
func TestPrompDanTiao(t *testing.T) {
	//handcards := []byte{0xEF, 0xFF, 0x02, 0x12, 0x22, 0x32, 0x9, 0x19, 0x29}
	handcards := []byte{0xEF, 0xFF, 0x0a, 0x1a, 0x2a, 0x3a, 0x9, 0x19, 0x29}
	discards := []byte{0x04}
	ty := &Cards{}
	ty.AnalyseUnSort(handcards)
	cardTypes := ty.Prompt(discards)

	for _, v := range cardTypes {
		t.Logf("cards: %#v %v", v.Cards, v.Kind)
	}

}

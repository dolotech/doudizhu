package algorithm

import (
	"testing"
)

func TestAssembleAllSanDaiDan(t *testing.T) {
	var an AnalyseCards
	hand := []byte{0x03, 0x13, 0x23, 0x33, 0x04, 0x14, 0x24, 0x34, 0x15, 0x25, 0x06, 0x16, 0x07, 0x17}
	an.Set(hand)
	arr := make([]AnalyseCards, 0)
	an.AllSanDaiDan(&arr)

	t.Logf("%v ", len(arr) == 8)
	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		cards := v.ColorRecover(hand)
		SortCards(cards, 0, int8(len(cards)-1))
		t.Logf("%#v %v ", cards, v.Weight())
	}

	hand = []byte{0x03, 0x13, 0x23, 0x33, 0x04, 0x14, 0x24, 0x34, 0x15, 0x25, 0x06, 0x16, 0x07, 0x17}
	an.Set(hand)
	arr = make([]AnalyseCards, 0)
	an.AllSanDaiDui(&arr)

	t.Logf("%v ", len(arr) == 8)
	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		cards := v.ColorRecover(hand)
		SortCards(cards, 0, int8(len(cards)-1))
		t.Logf("%#v %v ", cards, v.Weight())
	}

	return
}
func TestAssembleAllSiDaiDan(t *testing.T) {
	var an AnalyseCards
	hand := []byte{0x03, 0x13, 0x23, 0x33, 0x04, 0x14, 0x24, 0x34, 0x15, 0x25, 0x06, 0x16, 0x07, 0x17}
	an.Set(hand)
	arr := make([]AnalyseCards, 0)
	an.AllSiDaiDan(&arr)

	t.Logf("%v ", len(arr) == 20)
	/*for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		cards := v.ColorRecover(hand)
		SortCards(cards, 0, int8(len(cards)-1))
		t.Logf("%#v %v ", cards, v.Weight())
	}*/

	return
}
func TestAssembleAllSiDaiDui(t *testing.T) {
	var an AnalyseCards
	hand := []byte{0x03, 0x13, 0x23, 0x33, 0x04, 0x14, 0x24, 0x34, 0x15, 0x25, 0x06, 0x16, 0x07, 0x17}
	an.Set(hand)
	arr := make([]AnalyseCards, 0)
	an.AllSiDaiDui(&arr)

	t.Logf("%v ", len(arr) == 8)

	/*for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		cards := v.ColorRecover(hand)
		SortCards(cards, 0, int8(len(cards)-1))
		t.Logf("%#v %v ", cards, v.Weight())
	}
*/
	return
}
func TestAssembleAllFeijiDaiDui(t *testing.T) {
	var an AnalyseCards
	hand := []byte{0x03, 0x13, 0x04, 0x14, 0x24, 0x15, 0x25, 0x05, 0x06, 0x16, 0x07, 0x17}
	an.Set(hand)
	arr := make([]AnalyseCards, 0)
	an.AllFeijiDaiDui(&arr)

	t.Logf("%v ", len(arr) == 3)

	/*for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		cards := v.ColorRecover(hand)
		SortCards(cards, 0, int8(len(cards)-1))
		t.Logf("%#v %v ", cards, v.Weight())
	}
*/
	return
}
func TestAssembleAllFeijiDaiDan(t *testing.T) {
	var an AnalyseCards
	hand := []byte{0x05, 0x03, 0x04, 0x14, 0x24, 0x15, 0x25, 0x06, 0x16, 0x07} // 不拆对子
	an.Set(hand)
	arr := make([]AnalyseCards, 0)
	an.AllFeijiDaiDan(&arr)

	t.Logf("%v ", len(arr) == 6)
	/*
		for i := len(arr) - 1; i >= 0; i-- {
			v := arr[i]
			cards := v.ColorRecover(hand)
			SortCards(cards, 0, int8(len(cards)-1))
			t.Logf("%#v %v ", cards, v.Weight())
		}*/

	return
}
func TestAssemble(t *testing.T) {
	var an AnalyseCards
	//hand:=[]byte{ 0x02, 0x12, 0x22, 0x32 ,0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08}
	//hand:=[]byte{ 0x02, 0x12, 0x22, 0x32 ,0x03}
	//hand:=[]byte{ 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E}

	//hand := []byte{0x03, 0x04, 0x14, 0x24, 0x15, 0x25, 0x05, 0x06, 0x16, 0x07} // 不拆对子
	//hand := []byte{0x05, 0x34, 0x04, 0x14, 0x24, 0x15, 0x25, 0x07} // 拆炸弹做飞机
	hand := []byte{0x05, 0x25} // 对子不能拆开打

	an.Set(hand)
	/*
		arr := an.GetGroup()
		for i := len(arr) - 1; i >= 0; i-- {
			v := arr[i]
			cards := v.ColorRecover(hand)
			SortCards(cards, 0, int8(len(cards)-1))
			t.Logf("%#v %v ", cards, v.Weight())
		}

		return*/
	var anList []AnalyseCards

	wei := AnalyseWeightMax(&anList, an)
	//wei := AnalyseWeightMax(an)
	t.Logf("%v %v ", wei, len(anList))

	for i := 0; i < len(anList); i++ {
		v := anList[i]
		cards := v.ColorRecover(hand)
		SortCards(cards, 0, int8(len(cards)-1))
		t.Logf("%#v %v ", cards, v.Weight())
	}
}

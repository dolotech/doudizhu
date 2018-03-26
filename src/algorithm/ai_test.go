package algorithm

import (
	"testing"
	"github.com/golang/glog"
)

func yapaitest(handcards, discards []byte, length int, cardType []*Cards) bool {
	if length != len(cardType) {
		return false
	}
	for _, v := range cardType {
		if !v.Follow(discards) {
			//glog.Errorf("%#v %#v %#v", v.Cards, discards, v.Kind())
			glog.Errorf("%#v %v ", v.Value(), v.Kind())
			return false
		}

		for _, card := range v.Cards {
			exist := false
			for _, hcard := range handcards {
				if card == hcard {
					exist = true
					break
				}
			}

			if !exist {
				return false
			}

		}
	}
	return true
}

func TestSelectCards(t *testing.T) {

	//handcards := []byte{0x03, 0x05, 0x15, 0x25, 0x35, 0x06, 0x16, 0x26, 0x36, 0x18, 0x38, 0x27, 0x08, 0x28, 0x2D}
	//ty := &Cards{}
	//ty.AnalyseUnSort(handcards)
	//cardTypes := ty.GetGroup()
	//for _,v:=range  cardTypes{
	//	t.Logf("%#+v ",v.Cards)
	//}
	//for k,v:=range ty.analyseCards{
	//	t.Log(len(v),k,v)
	//}
	//t.Logf("%#+v  %+v  %#+v  ", len(cardTypes) ==3 , len(cardTypes),handcards)

	handcards := []byte{0x03, 0x05, 0x15, 0x35, 0x06, 0x26, 0x36, 0x36, 0x17, 0x37, 0x27, 0x08, 0x28, 0x29}
	ty := &Cards{}
	ty.AnalyseUnSort(handcards)
	cardTypes := ty.GetGroup()
	//cardTypes = ty.Prompt([]byte{0x03, 0x04, 0x15,  0x37, 0x06})
	for _, v := range cardTypes {

		//t.Logf("%#+v 权重 %v ", v.Cards, v.GetWeight())
		t.Logf("权重 %v ", v)
	}

	//t.Logf("%#+v  %+v  %#+v  ", len(cardTypes) == 3, len(cardTypes), handcards)
}

// 找出手牌中所有的顺子
func TestSearchAllShunZi(t *testing.T) {

	cardType := &Cards{}
	cardType.AnalyseUnSort([]byte{0x03, 0x04, 0x15, 0x35, 0x06, 0x2e, 0x3d, 0x3c, 0x17, 0x3b, 0x2a, 0x08, 0x28, 0x29})
	arr := make([]AnalyseCards, 0)
	cardType.analyseCards.AllShunZi(&arr)
	t.Logf("%+v  %#+v", len(arr) == 36, arr)

	cardType = &Cards{}
	cardType.AnalyseUnSort([]byte{0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x29})

	arr = make([]AnalyseCards, 0)
	cardType.analyseCards.AllShunZi(&arr)
	t.Logf("%+v  %#+v", len(arr) == 6, arr)
}

// 找出手牌中所有的连对
func TestSearchAllLianDui(t *testing.T) {

	//cardType:=&Cards{}
	//cardType.AnalyseUnSort([]byte{0x03, 0x04, 0x15, 0x35, 0x06, 0x2e, 0x3d, 0x3c, 0x17, 0x3b, 0x2a, 0x08, 0x28, 0x29})
	//t.Logf("%+v  %#+v",len(cardType.AllShunZi())==8,cardType.AllShunZi())
	arr := make([]AnalyseCards, 0)
	cardType := &Cards{}
	cardType.AnalyseUnSort([]byte{0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x29, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x29})
	cardType.analyseCards.AllLianDui(&arr)
	t.Logf("%+v  %#+v", len(arr) == 15, arr)
}

// 找出手牌中所有的三顺
func TestSearchAllSanShun(t *testing.T) {
	arr := make([]AnalyseCards, 0)
	cardType := &Cards{}
	cardType.AnalyseUnSort([]byte{0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08})
	cardType.analyseCards.AllSanShun(&arr)
	t.Logf("%+v  %#+v", len(arr) == 15, arr)
}

// 找出手牌中所有的飞机带单
func TestSearchAllFeiJIDaiDan(t *testing.T) {
	cardType := &Cards{}
	arr := make([]AnalyseCards, 0)
	hand := []byte{0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08}
	cardType.AnalyseUnSort(hand)
	cardType.analyseCards.AllFeijiDaiDan(&arr)

	for _, v := range arr {
		t.Logf("%#+v", v.ColorRecover(hand))
	}
	t.Logf("%+v  %#+v", len(arr) == 216, arr)
}

// 找出手牌中所有的飞机带对
func TestSearchAllAllFeiJIDaiDui(t *testing.T) {
	cardType := &Cards{}
	arr := make([]AnalyseCards, 0)
	cardType.AnalyseUnSort([]byte{0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08})
	cardType.analyseCards.AllFeijiDaiDui(&arr)
	t.Logf("%+v  %#+v", len(arr) == 5, arr)
}

func TestSubAndAdd(t *testing.T) {
	var an AnalyseCards
	an.Set([]byte{0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08})
	orLen := an.Len()
	t.Logf(" %v ", orLen == 18)
	var bb AnalyseCards
	bb.Set([]byte{0x03, 0x04, 0x15, 0x06, 0x17, 0x08})

	an.Sub(bb)
	an.Add(bb)
	t.Logf(" %v ", an.Len() == orLen)

}

func TestSearchMax2(t *testing.T) {
	var an AnalyseCards
	hand := []byte{0x02, 0x12, 0x22, 0x32}
	an.Set(hand)

	//discards:=[]byte{ 0x06, 0x16, 0x26 }

	var anList []AnalyseCards
	max := AnalyseWeightMax(&anList, an)
	t.Logf("%v %v ", max, len(anList))

	for _, v := range anList {
		cards := v.ColorRecover(hand)
		SortCards(cards, 0, int8(len(cards)-1))
		t.Logf("%#v %v ", cards, v.Weight())
	}
}

func TestSearchMax(t *testing.T) {
	var an AnalyseCards
	an.Set([]byte{0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08})

	array := an.GetGroup()

	var anList []AnalyseCards

	max := AnalyseWeightMax(&anList, an)
	t.Logf("%v %v %v", len(array) == 78, max == 3, len(anList) == 1)

}
func TestSearch(t *testing.T) {

	var an AnalyseCards
	an.SetKind(12)

	t.Logf("%+v ", an.GetKind() == 12)

	b := an

	//b = 0
	b.SetKind(3)
	an.SetKind(4)

	t.Logf("%+v %+v %+v", b.GetKind() == 3, an.GetKind() == 4, an != b)

}

func TestSearchMax1(t *testing.T) {
	var an AnalyseCards
	//hand:=[]byte{ 0x02, 0x12, 0x22, 0x32 ,0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08, 0x03, 0x04, 0x15, 0x06, 0x17, 0x08}
	//hand:=[]byte{ 0x02, 0x12, 0x22, 0x32 ,0x03,0x13,0x13,0x13}
	//hand:=[]byte{ 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E}

	hand := []byte{0x05, 0x03, 0x04, 0x14, 0x24, 0x15, 0x25, 0x06, 0x16,0x26, 0x07} // 不拆对子

	//hand := []byte{0x05, 0x34, 0x04, 0x14, 0x24, 0x15, 0x25, 0x07} // 拆炸弹做飞机
	//hand:=[]byte{ 0x05,0x25}  // 对子不能拆开打
	//hand := []byte{0x04, 0x04, 0x05, 0x06, 0x07, 0x08} // 对子不能拆开打

	an.Set(hand)

/*	arr := an.GetGroup()
	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]
		cards := v.ColorRecover(hand)
		SortCards(cards, 0, int8(len(cards)-1))
		t.Logf("%#v %v ", cards, v.Weight())
	}
	t.Logf("%v ",  len(arr))*/
	//return

	var anList []AnalyseCards
	max := AnalyseWeightMax(&anList, an)
	t.Logf("%v %v ",  len(anList),max)

	for i := 0; i < len(anList); i++ {
		v := anList[i]
		cards := v.ColorRecover(hand)
		SortCards(cards, 0, int8(len(cards)-1))
		t.Logf("%#v %v ", cards, v.Weight())
	}
}

type  AAA struct {
	Value int
} 

func TestZuHe(t *testing.T) {

	//t.Log(-300,-300&0xF)



}

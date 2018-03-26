package algorithm

import (
	"testing"
	"github.com/golang/glog"
)


// 拷贝测试
func  TestCopy(t *testing.T) {
	hand:=[]byte{0x02,0x03,0x04,0x05}
	a:=Copy(hand)
	glog.Errorf("a is %v,addressA is %v,addresshand is %v,addressA content is %v,addressHand is %v",
		a,&a[0],&hand[0],&a,&hand)
}

func TestLegal(t *testing.T) {
	//for _, v := range CARDS {
	//t.Logf("%v  %#x", Legal(v), v)
	//}
}


func TestValue(t *testing.T) {
	//for _, v := range CARDS {
	//t.Logf("  %#x %v ", v, GetValue(v))
	//}
}

func TestAnalyseCardArray(t *testing.T) {
	a := Cards{}
	//a.AnalyseUnSort( []byte{0x05, 0x15, 0x25,0x06, 0x16, 0x26,  0x27,0x37,0x07,  0x18, 0x28,0x08, 0x38, 0x2D, 0x2D})

	//t.Log(len(a.Cards))

	handcards := []byte{0xEF, 0xFF, 0x0a, 0x1a, 0x2a, 0x3a, 0x9, 0x19, 0x29}
	//discards := []byte{0x04}
	//a.AnalyseUnSort([]byte{0xEF,0xFF,0x02, 0x12, 0x03, 0x13, 0x23, 0x04, 0x14, 0x24, 0x05, 0x15, 0x25, 0x26, 0x36, 0x27, 0x37})
	a.AnalyseUnSort(handcards)

	//for k,v:=range a.analyseCards{
	//	t.Log(k,v)
	//}
	////
	t.Log(a.analyseCards, "====")
	for i := uint8(0); i < 21; i++ {
		//t.Log(i, a.Get(i))
	}

	//t.Log(a.Get(232))
}

func TestInitArray(t *testing.T) {

	//a.InitCardArray()

}

func TestShuffle(t *testing.T) {
	t.Logf("洗牌： %#x", Shuffle())
}


func TestDEAL(t *testing.T) {
	for i:=0;i<5;{
		glog.Errorf("第 %v 次:\n",i)
		a, b, c, d := Deal()
		glog.Errorf("发牌 %#v,%#v，%#v %#v", a, b, c, d)
		glog.Errorf("发牌 %v,%v，%v %v", len(a) == 17, len(b) == 17, len(c) == 17, len(d) == 3)
		i++
	}
}

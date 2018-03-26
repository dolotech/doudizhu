package algorithm

import (
	"testing"
	"github.com/golang/glog"
)

func  TestCard_GetValue(t *testing.T) {
	//hand:=[]byte{0x01,0x11,0x21,0x31,}
	for i,v:=range CARDS {

		glog.Errorf("i is %v,before is %v",i,v)
		a:=GetValue(v)
		glog.Errorf("i is %v,after is %v,card is %v",i,a,v)
	}
}


func TestCard_Legal(t *testing.T) {

}

func TestCard_GetSuits(t *testing.T) {

}
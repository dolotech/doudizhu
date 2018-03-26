package algorithm

import (
	"utils"
)

// 验证给定的牌是否是有效的牌值
func Legal(card byte) bool {
	if card == 0xEF || card == 0xFF {
		return true
	}
	if card>>4 > 4 || card>>4 < 0 {
		return false
	}

	if card&0xF > 0x0E || card&0xF < 0x02 {
		return false
	}
	return true
}

type Card byte

// 获取牌的花色
func GetSuits(card byte) int {
	return int(card >> 4)
}

// 获取牌的点数大小
func GetValue(card byte) byte {
	if card&0xF == 0x2 {
		return TWO
	} else if card == 0xEF {
		return XIAO_WANG
	} else if card == 0xFF {
		return DA_WANG
	}
	return card & 0xF
}

//拷贝一份完整牌
func Copy(c []byte) []byte {
	cards := make([]byte, len(c))
	copy(cards, c)
	return cards
}

// 指定的手牌里是否包括指定的牌型
func VerifyCards(source, target []byte) bool {
	for _, tv := range target {
		b := false
		for _, sv := range source {
			if tv == sv {
				b = true
				break
			}
		}
		if !b {
			return false
		}
	}
	return true
}

//洗牌
func Shuffle() []byte {
	d := Copy(CARDS)
	for n := 0; n < 3; n++ {
		for i := range d {
			j := utils.RandInt32N(int32(len(d)))
			d[i], d[j] = d[j], d[i]
		}
	}
	return d
}

package algorithm

//扑克牌54张，分别包含普通牌52张 2-10、J、Q、K、A （以上每种牌4个花色 黑桃、梅花、红心、方块）和小王大王
var CARDS = []byte{
	//方块
	0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E,
	//梅花
	0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E,
	//红桃
	0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x2A, 0x2B, 0x2C, 0x2D, 0x2E,
	//黑桃
	0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x3A, 0x3B, 0x3C, 0x3D, 0x3E,
	//小大王
	0xEF, 0xFF,
}

// 牌型
const (
	UNKNOW         = 0  // 未知牌型
	DAN_TIAO       = 1  //单条----从3(最小)A<2<小王<大王(最大)
	DUI_ZI         = 2  //对子----两张大小相同的牌，从3(最小)到2(最大)
	DAN_SHUN_ZI    = 3  //顺子----至少5张连续大小(从3到A，2和王不能用)的牌，例如8-9-10-J-Q；
	LIAN_DUI       = 4  //连对----至少3个连续大小(从3到A，2和王不能用)的对子，例如10-10-J-J-Q-Q-K-K；
	SAN_TIAO       = 5  //三条----三张大小相同的牌
	SAN_DAI_YI     = 6  //三带单----三张并带上任意一张牌，例如6-6-6-8，根据三张的大小来比较，例如9-9-9-3盖过8-8-8-A；
	SAN_DAI_YI_DUI = 7  //三带对----三张并带上一对，类似扑克中的副路(Full House)，根据三张的大小来比较，例如Q-Q-Q-6-6盖过10-10-10-K-K；
	SAN_SHUN_ZI    = 8  //三顺----至少2个连续大小(从3到A)的三张，例如4-4-4-5-5-5；
	FEI_JI_DAI_DAN = 9  //飞机带单----“三顺”带同数量的单牌，例如7-7-7-8-8-8-3-6，尽管三张2不能用，但能够带上单张2和王；
	FEI_JI_DAI_DUI = 10 //飞机带对----“三顺”带同数量的对子，例如8-8-8-9-9-9-4-4-J-J，尽管三张2不能用，但能够带上一对2，三张带上的单张和一对不能是混合的，例如3-3-3-4-4-4-6-7-7就是不合法的；
	SI_DAI_DAN     = 11 //四带两个----四张点数相同的牌带任意两张单牌
	SI_DAI_DUI     = 12 //四带两对----四张点数相同的牌带任意两张两个对子
	ZHA_DAN        = 13 //炸弹----四张大小相同的牌，炸弹能盖过除火箭外的其他牌型，大的炸弹能盖过小的炸弹
	WANG_ZHA       = 14 //王炸----一对王，这是最大的组合，能够盖过包括炸弹在内的任何牌型
)

// 牌型权重映射表
var Weights = map[uint8]uint8{
	DAN_TIAO:       1,
	DAN_SHUN_ZI:    1,
	WANG_ZHA:       3,
	DUI_ZI:         1,
	LIAN_DUI:       1,
	SAN_TIAO:       1,
	SAN_DAI_YI:     1,
	SAN_DAI_YI_DUI: 1,
	FEI_JI_DAI_DAN: 1,
	FEI_JI_DAI_DUI: 1,
	ZHA_DAN:        2,
	SI_DAI_DAN:     1,
	SI_DAI_DUI:     1,
	SAN_SHUN_ZI:    1,
}

const (
	DA_WANG   byte = 17
	XIAO_WANG byte = 16
	TWO       byte = 15
	A         byte = 14
	THREE     byte = 3
)

const (
	SEAT                = 3
	TOTAL               = 54
	HAND_CARDS          = 17
	HAND_CARDS_LANDLORD = 20
)

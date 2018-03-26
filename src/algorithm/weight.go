package algorithm

func (this *AnalyseCards) Weight() int {
	value := int(this.GetValue())
	length := this.Len()
	w := int(length) - 17

	var nValue = 0
	switch this.GetKind() {
	case DAN_TIAO: // 飞机带单
		nValue = value - 10
	case DUI_ZI: //对牌类型 (-7~7)
		nValue = value - 10
	case SAN_DAI_YI: //三带一单 (-7~7)
		nValue = value - 10
	case SAN_DAI_YI_DUI: //三带一对 (-7~7)
		nValue = value - 10
	case SAN_TIAO: //三条类型(-7~7)
		nValue = value - 10
	case DAN_SHUN_ZI: //单连类型(-6~6)
		nValue = value - 10 + 1
	case LIAN_DUI: //对连类型(-6~6)
		nValue = value - 10 + 1
	case SAN_SHUN_ZI: //三顺(0~6)
		nValue = (value - 3 + 1) / 2
	case FEI_JI_DAI_DAN: //飞机带单(0~6)
		nValue = (value - 3 + 1) / 2
	case FEI_JI_DAI_DUI: //飞机带对(0~6)
		nValue = (value - 3 + 1) / 2
	case SI_DAI_DAN: //四带两单(0~6)
		nValue = (value - 3 ) / 2
	case SI_DAI_DUI: //四带两对(0~6)
		nValue = (value - 3 ) / 2
	case ZHA_DAN: //炸弹类型(7~19)
		nValue = value - 3 + 7
	case WANG_ZHA: //王炸类型
		nValue = 20
	}
	return nValue - 20 + w
}

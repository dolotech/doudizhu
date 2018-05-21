# Go版本实现斗地主核心算法
#### 紧凑的数据结构高效核心算法
#### Author:Michael 
#### Email:<dolotech@163.com>

```go
                                type AnalyseCards uint64

                                func (this *AnalyseCards) Set(cards []byte) {
                                        *this = 0
                                        l := uint8(len(cards))
                                        for i := uint8(0); i < l; i++ {
                                                card := (GetValue(cards[i]) - 3) * 3
                                                count := ((*this >> card) & 0x07) + 1
                                                *this &= (^(0x07 << card))
                                                *this |= (count << card)
                                        }

                                        this.SetStartEnd()
                                }

                                func (this *AnalyseCards) SetSortCards(cards []byte) {
                                        *this = 0
                                        l := uint8(len(cards))
                                        for i := uint8(0); i < l; i++ {
                                                card := (cards[i] - 3) * 3
                                                count := ((*this >> card) & 0x07) + 1
                                                *this &= (^(0x07 << card))
                                                *this |= (count << card)
                                        }
                                        this.SetStartEnd()
                                }

                                func (this *AnalyseCards) Get(card byte) uint8 {
                                        return uint8((*this >> ((card - 3) * 3 )) & 0x07)
                                }

                                func (this *AnalyseCards) SetStartEnd() {
                                        var start, end uint8
                                        for i := uint8(THREE); i <= DA_WANG; i++ {
                                                if start == 0 && this.Get(i) > 0 {
                                                        start = i
                                                }

                                                if end == 0 && this.Get(DA_WANG+THREE-i) > 0 {
                                                        end = DA_WANG + THREE - i
                                                }

                                                if start > 0 && end > 0 {
                                                        break
                                                }
                                        }

                                        card := uint8(46)
                                        *this &= (^(0xF << card))
                                        *this |= (AnalyseCards(start-3) << card)

                                        card = uint8(50)
                                        *this &= (^(0xF << card))
                                        *this |= (AnalyseCards(end-3) << card)
                                }

                                func (this *AnalyseCards) Reset() {
                                        this.SetValue(0)
                                        this.SetKind(0)
                                        this.SetStartEnd()
                                }

                                func (this *AnalyseCards) End() uint8 {
                                        return byte((*this>>50 )&0xF) + 3
                                }
                                func (this *AnalyseCards) Start() uint8 {
                                        return byte((*this>>46 )&0xF) + 3
                                }

                                func (this *AnalyseCards) GetValue() byte {
                                        return byte((*this >> 59 ) & 0x1F)
                                }

                                func (this *AnalyseCards) SetValue(value byte) {
                                        card := uint8(59)
                                        *this &= (^(0x1F << card))
                                        *this |= (AnalyseCards(value) << card)
                                }

                                func (this *AnalyseCards) GetKind() uint8 {
                                        return byte((*this >> 54 ) & 0x1F)
                                }

                                func (this *AnalyseCards) SetKind(kind uint8) {
                                        card := uint8(54)
                                        *this &= (^(0x1F << card))
                                        *this |= (AnalyseCards(kind) << card)
                                }

                                func (this *AnalyseCards) GetWeight() uint8 {
                                        return Weights[this.GetKind()]
                                }

                                func (this *AnalyseCards) ColorRecover(handcards []byte) []byte {
                                        analyseCards := *this
                                        cards := make([]byte, 0, 10)
                                        for _, v := range handcards {
                                                if analyseCards > 0 {
                                                        card := GetValue(v)
                                                        if analyseCards.Get(card) > 0 {
                                                                cards = append(cards, v)
                                                                analyseCards.Incr(card, -1)
                                                        }
                                                }
                                        }
                                        return cards
                                }

                                func (this *AnalyseCards) Flat() []byte {
                                        cards := make([]byte, 0, this.Len())
                                        start := this.Start()
                                        end := this.End()
                                        for i := start; i <= end; i++ {
                                                count := this.Get(i)
                                                if count > 0 {
                                                        for j := uint8(0); j < count; j++ {
                                                                cards = append(cards, i)
                                                        }
                                                }
                                        }
                                        return cards
                                }
                                func (this *AnalyseCards) Len() uint8 {
                                        var count uint8
                                        start := this.Start()
                                        end := this.End()
                                        for i := start; i <= end; i++ {
                                                count += this.Get(i)  
                                        }
                                        return count
                                }

                                func (this *AnalyseCards) Sub(cards AnalyseCards) AnalyseCards {
                                        t := *this
                                        start := cards.Start()
                                        end := cards.End()
                                        for i := start; i <= end; i++ {
                                                count := cards.Get(i)
                                                if count > 0 {
                                                        t.Incr(i, -int8(count))
                                                }
                                        }
                                        t.SetValue(0)
                                        t.SetKind(0)
                                        return t
                                }
                                func (this *AnalyseCards) Add(cards AnalyseCards) AnalyseCards {
                                        t := *this
                                        start := cards.Start()
                                        end := cards.End()
                                        for i := start; i <= end; i++ {
                                                count := cards.Get(i)
                                                if count > 0 {
                                                        t.Incr(i, int8(count))
                                                }
                                        }
                                        t.SetValue(0)
                                        t.SetKind(0)
                                        return t
                                }

                                func (this *AnalyseCards) Incr(card byte, num int8) {
                                        card = (card - 3) * 3
                                        count1 := ((*this >> card) & 0x07)
                                        count := int64(count1) + int64(num)
                                        *this &= (^(0x07 << card))
                                        if count > 0 {
                                                *this |= ((AnalyseCards(count & 0x07) ) << card)
                                        }
                                        this.SetStartEnd()
                                }

                                func (this *AnalyseCards) Padding(count uint8, length uint8, card byte) {
                                        for m := uint8(0); m < count; m++ {
                                                c := card - m
                                                this.Incr(c, int8(length))
                                        }
                                        this.SetValue(0)
                                        this.SetKind(0)
                                }
```

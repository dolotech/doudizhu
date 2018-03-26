package algorithm

// 对牌值从小到大排序，采用快速排序算法
func SortCards(arr []byte, start, end int8) {
	if start < end {
		i, j := start, end
		card := arr[(start+end)/2]
		key := GetValue(card)
		suit := card >> 4
		for i <= j {
			for GetValue(arr[i]) < key || (GetValue(arr[i]) == key && arr[i]>>4 < suit) {
				i++
			}
			for GetValue(arr[j]) > key || (GetValue(arr[j]) == key && arr[j]>>4 > suit) {

				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			SortCards(arr, start, j)
		}
		if end > i {
			SortCards(arr, i, end)
		}
	}
}

// 对牌值从大到小排序
func BigSortCards(arr []byte, start, end int8) {
	if start < end {
		i, j := start, end
		card := arr[(start+end)/2]
		key := GetValue(card)
		suit := card >> 4
		for i <= j {
			for GetValue(arr[i]) > key || (GetValue(arr[i]) == key && arr[i]>>4 < suit) {
				i++
			}
			for GetValue(arr[j]) < key || (GetValue(arr[j]) == key && arr[j]>>4 > suit) {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if start < j {
			BigSortCards(arr, start, j)
		}
		if end > i {
			BigSortCards(arr, i, end)
		}
	}
}

func Sort(cards []byte, start, end int8) {
	if start < end {
		i, j := start, end
		card := cards[(start+end)/2]
		for i <= j {
			for cards[i] < card {
				i++
			}
			for cards[j] > card {
				j--
			}
			if i <= j {
				cards[i], cards[j] = cards[j], cards[i]
				i++
				j--
			}
		}
		if start < j {
			Sort(cards, start, j)
		}
		if end > i {
			Sort(cards, i, end)
		}
	}
}

func SearchMaxWeight(cards []AnalyseCards) AnalyseCards {
	leng := len(cards)
	var tmp = cards[0]
	var wei = tmp.Weight()
	for i := 1; i < leng; i++ {
		twei := cards[i].Weight()
		if twei > wei {
			tmp = cards[i]
			wei = twei
		}
	}
	return tmp
}
func SortAnalyseCardsByWeigth(cards []AnalyseCards, start, end int) {
	if start < end {
		i, j := start, end
		card := cards[(start+end)/2].Weight()
		for i <= j {
			for cards[i].Weight() < card {
				i++
			}
			for cards[j].Weight() > card {
				j--
			}
			if i <= j {
				cards[i], cards[j] = cards[j], cards[i]
				i++
				j--
			}
		}
		if start < j {
			SortAnalyseCardsByWeigth(cards, start, j)
		}
		if end > i {
			SortAnalyseCardsByWeigth(cards, i, end)
		}
	}
}

func SortAnalyseCardsByLen(cards []AnalyseCards, start, end int) {
	if start < end {
		i, j := start, end
		card := cards[(start+end)/2].Len()
		for i <= j {
			for cards[i].Len() < card {
				i++
			}
			for cards[j].Len() > card {
				j--
			}
			if i <= j {
				cards[i], cards[j] = cards[j], cards[i]
				i++
				j--
			}
		}
		if start < j {
			SortAnalyseCardsByLen(cards, start, j)
		}
		if end > i {
			SortAnalyseCardsByLen(cards, i, end)
		}
	}
}

func NextSeat(seat uint32) uint32 {
	if seat == SEAT {
		seat = 1
	}
	return seat + 1
}

func combineloop(arr []byte, rel *[][]byte, r []byte, i int, n int) {
	if n <= 0 {
		return
	}
	rlen := len(r) - n
	alen := len(arr)
	for j := i; j < alen; j++ {
		r[rlen] = arr[j]
		if n == 1 {
			or := make([]byte, len(r))
			copy(or, r)
			*rel = append(*rel, or)
		} else {
			combineloop(arr, rel, r, j+1, n-1)
		}
	}
}

func CombineUnique(cards []byte, n int) [][]byte {
	rel := Combine(cards, n)
	for i := 0; i < len(rel)-1; i++ {
		for j := i + 1; j < len(rel); j++ {
			count:=0
			for k := 0; k < n; k++ {
				if rel[i][k] != rel[j][k] {
					break
				}
				count++
			}
			if count == n{
				rel = append(rel[:i], rel[i+1:]...)
			}
		}
	}
	return rel[:]
}

func Combine(cards []byte, n int) [][]byte {
	arr := make([]byte, n)
	rel := make([][]byte, 0)
	combineloop(cards[:], &rel, arr[:], 0, n)
	return rel[:]
}

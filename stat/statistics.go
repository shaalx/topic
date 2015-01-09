package stat

import (
	"fmt"
	"github.com/shaalx/topic/seg"
	"sort"
	"strings"
)

type Stats map[string]int32

type Cell struct {
	Word string
	Freq int32
}

type CellSlice []*Cell

var sentence string
var threshold = int32(3)

func init() {
	seg.SetDictionary("dict.txt") // 设定字典
	sentence = seg.ReadAll("file.txt")
}

func Stating() (Stats, []string) {
	stats := make(Stats, 1)
	segs := seg.Cut(sentence, false, true)
	for _, it := range segs {
		stats[it]++
	}
	return stats, segs
}

func (s Stats) Map2Slice() CellSlice {
	cellSlice := make(CellSlice, 0)
	for k, v := range s {
		if "，" == k {
			continue
		}
		if "。" == k {
			continue
		}
		if "　" == k {
			continue
		}
		if "\n\t" == k {
			continue
		}
		r := []rune(k)
		if 13 == r[0] && 10 == r[1] {
			continue
		}
		cell := Cell{k, v}
		cellSlice = append(cellSlice, &cell)
	}
	return cellSlice
}

func (c CellSlice) Len() int {
	return len(c)
}

func (c CellSlice) Less(i, j int) bool {
	return c[i].Freq < c[j].Freq
}

func (c CellSlice) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c CellSlice) String() {
	for i, v := range c {
		fmt.Println(i, v)
	}
}

func FirstStep() {
	freq, segs := Stating()
	fmt.Println(freq, segs)
	cells := freq.Map2Slice()
	sort.Sort(sort.Reverse(cells))
	cells.String()
	Search(cells, segs)
}

func Search(cells CellSlice, segs []string) {
	flags := make([]int, 0)
	for _, v := range cells {
		if v.Freq < threshold {
			break
		}
		index := strings.Index(sentence, v.Word)
		flags = append(flags, index)
	}
	sort.Ints(flags)
	fmt.Println(flags)
	for i := 0; i < len(flags)-1; i += 2 {
		fmt.Println(segs[flags[i]:flags[i+1]])
	}
}

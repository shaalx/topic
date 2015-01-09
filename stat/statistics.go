package stat

import (
	"fmt"
	"github.com/shaalx/topic/seg"
	"regexp"
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
var threshold = int32(4)
var filter []string = []string{
	"的", "在", "和", "了", "也", "上", "还", "是", "年", "有", "，", "。", " ", "都", "而", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "", "",
}

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
		filtered := false
		for _, it := range filter {
			if it == k {
				filtered = true
			}
		}
		if filtered {
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
		if threshold >= v.Freq {
			continue
		}
		fmt.Println(i, v)
	}
}

func FirstStep() {
	freq, segs := Stating()
	// fmt.Println(freq, segs)
	cells := freq.Map2Slice()
	sort.Sort(sort.Reverse(cells))
	cells.String()
	Search(cells, segs)
}

func Search(cells CellSlice, segs []string) {
	flags := make([]int, 0)
	// flagsend := make([]int, 0)
	senbs := []byte(sentence)
	for _, v := range cells {
		if v.Freq < threshold {
			break
		}
		// index := strings.Index(sentence, v.Word)
		// flags = append(flags, index)
		rege := regexp.MustCompile(v.Word)
		indexs := rege.FindAllIndex(senbs, -1)
		for _, it := range indexs {
			flags = append(flags, it[0])
		}
	}

	// // 标点句号 结尾
	// regend := regexp.MustCompile("。")
	// indexsend := regend.FindAllIndex(senbs, -1)
	// for _, it := range indexsend {
	// 	flagsend = append(flagsend, it[0])
	// }
	result := make([]string, 0)

	sort.Ints(flags)
	// fmt.Println(flags)
	for i := 0; i < len(flags); i++ {
		// fmt.Println(segs[flags[i]:flags[i+1]])
		// 查找行首以 H 开头，以。结尾的字符串
		reg := regexp.MustCompile(`^.*。`)
		ins := reg.FindAllString(sentence[flags[i]:], -1)
		for _, in := range ins {
			if resContains(result, in) {
				continue
			}
			if strings.Count(in, "。") > 2 {
				continue
			}
			fmt.Printf("%v\n", in)
			result = append(result, in)
		}
	}
	// fmt.Printf("%v\n", result)
}

func resContains(result []string, in string) bool {
	for _, it := range result {
		if strings.Contains(it, in) {
			return true
		}
	}
	return false
}

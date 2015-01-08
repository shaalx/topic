package main

import (
	"fmt"
	"github.com/shaalx/topic/seg"
	"strings"
)

var sentence = "我来到北京清华大学"

func main() {
	seg.SetDictionary("dict.txt") // 设定字典
	fmt.Printf("【全模式】: %s\n\n", strings.Join(seg.Cut(sentence, true, true), "/ "))
	fmt.Printf("【精确模式】: %s\n\n", strings.Join(seg.Cut(sentence, false, true), "/ "))
	fmt.Printf("【新词识别】：%s\n\n", strings.Join(seg.Cut("他来到了网易杭研大厦", false, true), ", "))
	fmt.Printf("【搜索引擎模式】：%s\n\n", strings.Join(seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true), ", "))
}

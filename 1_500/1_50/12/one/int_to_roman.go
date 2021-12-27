package one

import "strings"

// 通用逻辑
// 基础数据
var (
	nums  = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman = map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
)

func IntToRoman(num int) string {
	ans := strings.Builder{}
	index := 0 // 记录当前到达的数字
	for num > 0 {
		if num >= nums[index] {
			ans.WriteString(roman[nums[index]])
			num -= nums[index]
			continue
		}
		index++
	}
	return ans.String()
}

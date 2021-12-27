package two

import "strings"

// 本题明确说明num的范围是1-3999
// 那么可以直接枚举
var roman = [4][10]string{
	{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}, // 1-9
	{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}, // 10,20...90
	{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"}, // 100,200...900
	{"", "M", "MM", "MMM"}, // 1000,2000,3000
}

func IntToRoman(num int) string {
	ans := strings.Builder{}
	ans.WriteString(roman[3][num/1000])     // 1000
	ans.WriteString(roman[2][num%1000/100]) // 100
	ans.WriteString(roman[1][num%100/10])   // 10
	ans.WriteString(roman[0][num%10])       // 1
	return ans.String()
}

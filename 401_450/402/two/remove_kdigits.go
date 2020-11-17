package two

import "strings"

func RemoveKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	// 递增栈
	stack := make([]byte, 0, len(num))
	for i := range num {
		for k > 0 && len(stack) > 0 && num[i] < stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
			k--
		}
		stack = append(stack, num[i])
	}
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if len(ans) == 0 {
		return "0"
	}
	return ans
}

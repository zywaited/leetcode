package one

func LongestPalindrome(s string) int {
	// 最大长度就是所有的偶数+1（如果有多余的字符）
	nums := make(map[byte]int, len(s))
	for index := range s {
		nums[s[index]]++
	}
	ans := 0
	for _, num := range nums {
		if num%2 == 0 {
			ans += num
			continue
		}
		ans += num - 1
	}
	if ans < len(s) {
		ans++
	}
	return ans
}

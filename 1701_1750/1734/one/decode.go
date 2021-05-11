package one

func Decode(encoded []int) []int {
	n := len(encoded) + 1
	f := 0
	for i := 1; i <= n; i++ {
		f ^= i
	}
	for i := 1; i < len(encoded); i += 2 {
		f ^= encoded[i]
	}
	ans := []int{f}
	for _, e := range encoded {
		ans = append(ans, ans[len(ans)-1]^e)
	}
	return ans
}

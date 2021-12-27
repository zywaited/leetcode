package one

// O(N)
func MinSubArrayLen(s int, nums []int) int {
	// 滑动窗口
	windows := make([]int, 0, len(nums))
	wi := 0 // 为了解决缩减数组的性能消耗
	sum := 0
	ans := len(nums) + 1
	for _, num := range nums {
		sum += num
		windows = append(windows, num)
		for sum >= s {
			ans = min(ans, len(windows)-wi)
			sum -= windows[wi]
			wi++
		}
	}
	if ans > len(nums) {
		ans = 0
	}
	return ans
}

func min(f, s int) int {
	if f <= s {
		return f
	}
	return s
}

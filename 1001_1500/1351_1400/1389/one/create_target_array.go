package one

func CreateTargetArray(nums []int, index []int) []int {
	ans := make([]int, 0, len(nums))
	for i := 0; i < len(nums); i++ {
		// 增加一位
		ans = append(ans, 0)
		for ni := len(ans) - 1; ni > index[i]; ni-- {
			ans[ni] = ans[ni-1]
		}
		ans[index[i]] = nums[i]
	}
	return ans
}

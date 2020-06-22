package one

func Shuffle(nums []int, n int) []int {
	ans := make([]int, n<<1)
	ansIndex := 0
	for index := 0; index < n; index++ {
		ans[ansIndex] = nums[index]
		ansIndex++
		ans[ansIndex] = nums[index+n]
		ansIndex++
	}
	return ans
}

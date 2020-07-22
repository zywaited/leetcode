package one

func FindMin(nums []int) int {
	ans := nums[0]
	for index := 1; index < len(nums); index++ {
		if nums[index] < ans {
			ans = nums[index]
		}
	}
	return ans
}

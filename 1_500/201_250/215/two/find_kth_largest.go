package two

// 类似冒泡
func FindKthLargest(nums []int, k int) int {
	index := 0
	for index < k {
		mi := index
		for i := index + 1; i < len(nums); i++ {
			if nums[i] > nums[mi] {
				mi = i
			}
		}
		if mi != index {
			nums[index], nums[mi] = nums[mi], nums[index]
		}
		index++
	}
	return nums[index-1]
}

package one

func FindUnsortedSubarray(nums []int) int {
	index := 1
	for ; index < len(nums) && nums[index] >= nums[index-1]; index++ {
	}
	if index == len(nums) {
		return 0
	}
	start := index - 1
	end := index
	max := nums[start]
	for index = end; index < len(nums); index++ {
		if nums[index] >= max {
			max = nums[index]
			continue
		}
		end = index
		for ; start >= 0 && nums[start] > nums[index]; start-- {
		}
	}
	return end - start
}

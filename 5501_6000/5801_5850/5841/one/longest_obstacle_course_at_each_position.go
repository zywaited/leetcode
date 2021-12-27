package one

import "sort"

func LongestObstacleCourseAtEachPosition(obstacles []int) []int {
	nums := make([]int, 0, len(obstacles))
	ans := make([]int, 0, len(obstacles))
	for _, obstacle := range obstacles {
		index := sort.Search(len(nums), func(i int) bool {
			return nums[i] > obstacle
		})
		ans = append(ans, index+1)
		if index == len(nums) {
			nums = append(nums, obstacle)
			continue
		}
		nums[index] = obstacle
	}
	return ans
}

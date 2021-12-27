package one

func FindSubsequences(nums []int) [][]int {
	ans := make([][]int, 0)
	findCurrentSubsequences(nums, 0, nil, &ans)
	return ans
}

func findCurrentSubsequences(nums []int, index int, cns []int, ans *[][]int) {
	dm := make(map[int]bool)
	for nextIndex := index; nextIndex < len(nums); nextIndex++ {
		if !dm[nums[nextIndex]] && (index == 0 || nums[nextIndex] >= nums[index-1]) {
			cns = append(cns, nums[nextIndex])
			if len(cns) > 1 {
				*ans = append(*ans, append(make([]int, 0, len(cns)+1), cns...))
			}
			dm[nums[nextIndex]] = true
			findCurrentSubsequences(nums, nextIndex+1, cns, ans)
			cns = cns[:len(cns)-1]
		}
	}
}

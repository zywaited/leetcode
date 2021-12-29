package one

func CountQuadruplets(nums []int) int {
	ans := 0
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for k := j + 1; k < len(nums)-1; k++ {
				for l := k + 1; l < len(nums); l++ {
					if nums[i]+nums[j]+nums[k] == nums[l] {
						ans++
					}
				}
			}
		}
	}
	return ans
}

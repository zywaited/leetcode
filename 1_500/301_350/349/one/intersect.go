package one

func Intersect(nums1 []int, nums2 []int) []int {
	nn := make(map[int]bool, len(nums1))
	ans := make([]int, 0, len(nums2))
	for _, num := range nums1 {
		nn[num] = true
	}
	for _, num := range nums2 {
		if nn[num] {
			ans = append(ans, num)
			nn[num] = false
		}
	}
	return ans
}

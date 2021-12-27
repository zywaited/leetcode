package one

func Intersect(nums1 []int, nums2 []int) []int {
	nn := make(map[int]int, len(nums1))
	ans := make([]int, 0, len(nums2))
	for _, num := range nums1 {
		nn[num]++
	}
	for _, num := range nums2 {
		if nn[num] > 0 {
			ans = append(ans, num)
			nn[num]--
		}
	}
	return ans
}

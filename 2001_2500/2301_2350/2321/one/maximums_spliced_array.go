package one

func MaximumsSplicedArray(nums1 []int, nums2 []int) int {
	maxAddNum1 := 0
	currAddNum1 := 0
	sums1 := 0
	sums2 := 0
	maxAddNum2 := 0
	currAddNum2 := 0
	for i := 0; i < len(nums1); i++ {
		sums1 += nums1[i]
		sums2 += nums2[i]
		currAddNum1 += nums2[i] - nums1[i]
		currAddNum1 = max(currAddNum1, 0)
		maxAddNum1 = max(maxAddNum1, currAddNum1)
		currAddNum2 += nums1[i] - nums2[i]
		currAddNum2 = max(currAddNum2, 0)
		maxAddNum2 = max(maxAddNum2, currAddNum2)
	}
	return max(sums1+maxAddNum1, sums2+maxAddNum2)
}

func max(f, s int) int {
	if s <= f {
		return f
	}
	return s
}

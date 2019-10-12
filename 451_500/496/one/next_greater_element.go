package one

func NextGreaterElement(nums1 []int, nums2 []int) []int {
	stack := make([]int, 0, len(nums2))
	indexes := make(map[int]int, len(nums2))
	for index := len(nums2) - 1; index >= 0; index-- {
		// 从尾到头一个个判断是否有比当前值大的数字
		for len(stack) > 0 && stack[len(stack)-1] <= nums2[index] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			// 没有
			indexes[nums2[index]] = -1
		} else {
			// 最后一个就是比当前大的值
			indexes[nums2[index]] = stack[len(stack)-1]
		}
		// 把当前值加入栈
		stack = append(stack, nums2[index])
	}
	rs := make([]int, len(nums1))
	for index, num := range nums1 {
		rs[index] = indexes[num]
	}
	return rs
}

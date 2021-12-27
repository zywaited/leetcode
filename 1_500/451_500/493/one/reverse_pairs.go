package one

func ReversePairs(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	n1 := append([]int{}, nums[0:len(nums)>>1]...)
	n2 := append([]int{}, nums[len(nums)>>1:]...)
	cn := ReversePairs(n1) + ReversePairs(n2)
	// 计算数量(左边对右边)
	n := 0
	for _, num := range n1 {
		for n < len(n2) && num > n2[n]<<1 {
			n++
		}
		cn += n
	}
	// 归并排序
	i1, i2 := 0, 0
	for i := 0; i < len(nums); i++ {
		if i1 < len(n1) && (i2 == len(n2) || n1[i1] < n2[i2]) {
			nums[i] = n1[i1]
			i1++
			continue
		}
		nums[i] = n2[i2]
		i2++
	}
	return cn
}

package two

func MoveZeroes(nums []int) {
	index := 0 // 不是0的索引
	for i, num := range nums {
		if num != 0 {
			if i != index {
				nums[index] = num
			}
			index++
		}
	}
	for ; index < len(nums); index++ {
		nums[index] = 0
	}
}

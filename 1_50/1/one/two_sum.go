package one

func TwoSum(nums []int, target int) []int {
	var response []int
	num := len(nums)
	if num < 2 {
		return response
	}

	// 一次hash执行，即便输入数组存在相同的值，也能找到对应的索引
	numIndex := make(map[int]int)
	for index, num := range nums {
		another := target - num
		if i, ok := numIndex[another]; ok && i != index {
			response = append(append(response, i), index)
			break
		}

		numIndex[num] = index
	}

	return response
}

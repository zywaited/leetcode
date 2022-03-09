package two

func BestRotation(nums []int) int {
	// 前缀和【改造第一种方法即可】
	sums := make([]int, len(nums)+1)
	for index, num := range nums {
		if index < num {
			sums[index+1] += 1
			sums[len(nums)-(num-index)+1] -= 1
			continue
		}
		sums[0]++
		if num == 0 {
			continue
		}
		sums[index-num+1] -= 1
		sums[index+1] += 1
	}
	max := sums[0]
	ans := 0
	for index := 1; index < len(nums); index++ {
		sums[index] += sums[index-1]
		if max < sums[index] {
			max = sums[index]
			ans = index
		}
	}
	return ans
}

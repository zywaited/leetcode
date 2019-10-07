package one

func SubsetsWithDup(nums []int) [][]int {
	// 思路
	// 计算出每个数字的个数，先按照没有重复数字的情况列出组合
	// 然后逐渐添加重复数从1-N

	// 例如: 1,2,2
	// 最开始排列[]
	// [] [1]([]+1)
	// [] [1] [2]([]+2) [1,2]([1]+2) [2,2]([]+2,2) [1,2,2]([1]+2,2)
	rs := make([][]int, 0, 1<<uint(len(nums)))
	rs = append(rs, []int{}) // 空集合
	end := 1                 // rs上一次的集合最终点
	counts := make(map[int]int, len(nums))
	for _, num := range nums {
		counts[num]++
	}
	for num, count := range counts {
		for i := 0; i < end; i++ {
			for j := 1; j <= count; j++ {
				tmp := make([]int, len(rs[i]), len(rs[i])+j)
				copy(tmp, rs[i])
				for k := 0; k < j; k++ {
					tmp = append(tmp, num)
				}
				rs = append(rs, tmp)
			}
		}
		end = len(rs)
	}
	return rs
}

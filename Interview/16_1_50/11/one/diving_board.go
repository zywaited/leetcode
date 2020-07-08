package one

func DivingBoard(shorter int, longer int, k int) []int {
	min := shorter * k
	max := longer * k
	ans := make([]int, 0)
	if min > 0 {
		for min <= max {
			ans = append(ans, min)
			if shorter == longer {
				break
			}
			min += longer - shorter
		}
	}
	return ans
}

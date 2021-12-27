package one

func CountMaxOrSubsets(nums []int) int {
	curr := map[int]int{}
	var next map[int]int
	max := 0
	for _, num := range nums {
		max |= num
		next = map[int]int{num: 1}
		for b, c := range curr {
			next[b|num] += c
			next[b] += c
		}
		curr = next
	}
	return curr[max]
}

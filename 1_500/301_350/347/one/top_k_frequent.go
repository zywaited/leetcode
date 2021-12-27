package one

// O(NK)
func TopKFrequent(nums []int, k int) []int {
	mp := make(map[int]int)
	for _, num := range nums {
		mp[num]++
	}
	rs := make([]int, k)
	// 每次取出最大的
	i := 0
	for i < k {
		mn, mt := 0, 0
		for n, t := range mp {
			if t > mt {
				mt = t
				mn = n
			}
		}
		rs[i] = mn
		delete(mp, mn)
		i++
	}
	return rs
}

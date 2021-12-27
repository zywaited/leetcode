package two

import "sort"

func ChalkReplacer(chalk []int, k int) int {
	sums := make([]int, len(chalk)+1)
	for i := 1; i <= len(chalk); i++ {
		sums[i] = sums[i-1] + chalk[i-1]
		if sums[i] > k {
			return i - 1
		}
	}
	remainder := k % sums[len(chalk)]
	return sort.Search(len(sums), func(i int) bool {
		return sums[i] > remainder
	}) - 1
}

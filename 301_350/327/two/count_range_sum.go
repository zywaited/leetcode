package two

func CountRangeSum(nums []int, lower int, upper int) int {
	sums := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	tmp := make([]int, 0, len(sums))
	var merge func(s, e int) int
	merge = func(s, e int) int {
		// sums 包含了一个额外的0，这里就不能算本身
		if s == e {
			return 0
		}
		m := s + (e-s)>>1
		cnt := merge(s, m) + merge(m+1, e)
		tmp = tmp[:0]
		i := s
		j := m + 1
		l := j
		r := j
		for i <= m || j <= e {
			if j <= e && (i > m || sums[j] <= sums[i]) {
				tmp = append(tmp, sums[j])
				j++
				continue
			}
			tmp = append(tmp, sums[i])
			for ; l <= e && sums[l]-sums[i] < lower; l++ {
			}
			for ; r <= e && sums[r]-sums[i] <= upper; r++ {
			}
			cnt += r - l
			i++
		}
		i = s
		j = 0
		for ; j < len(tmp); j++ {
			sums[i] = tmp[j]
			i++
		}
		return cnt
	}
	return merge(0, len(sums)-1)
}

package one

func LongestMountain(A []int) int {
	ans := 0
	prev := 0
	flag := false
	for index := 0; index < len(A); index++ {
		if !flag {
			if index+1 == len(A) || A[index] < A[index+1] {
				continue
			}
			if A[index] == A[index+1] || index == prev {
				prev = index + 1
				continue
			}
			flag = true
			continue
		}
		ans = max(ans, index-prev+1)
		if index+1 == len(A) {
			continue
		}
		if A[index] <= A[index+1] {
			flag = false
			prev = index
			if A[index] == A[index+1] {
				prev++
			}
		}
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

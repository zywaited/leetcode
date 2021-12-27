package one

func CanThreePartsEqualSum(A []int) bool {
	sum := 0
	for _, num := range A {
		sum += num
	}
	if sum%3 > 0 {
		return false
	}
	avg := sum / 3
	i, j := 0, len(A)-1
	l, r := A[i], A[j]
	// i+1 < j保证中间一定有数据
	for i+1 < j {
		if l == avg && r == avg {
			return true
		}
		if l != avg {
			i++
			l += A[i]
		}
		if r != avg {
			j--
			r += A[j]
		}
	}
	return false
}

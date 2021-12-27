package one

func MinKBitFlips(A []int, K int) int {
	ans := 0
	i := 0
	rt := 0 // 是否反转
	for ; i < len(A); i++ {
		if i >= K && A[i-K] > 1 {
			rt ^= 1 // 下一个区域还原上一次
			A[i-K] -= 2
		}
		if rt != A[i] {
			continue
		}
		// 这时候代表A[i]已经变成0
		if len(A)-i < K {
			break
		}
		rt ^= 1
		A[i] += 2
		ans++
	}
	if i < len(A) {
		ans = -1
	}
	return ans
}

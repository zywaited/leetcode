package two

func FindKthBit(n int, k int) byte {
	if n == 1 {
		return '0'
	}
	m := 1 << uint(n-1)
	if k == m {
		return '1'
	}
	if k < m {
		return FindKthBit(n-1, k)
	}
	// S5 = "0111001 1 0110001 1 0111001 0 0110001"
	// 转成左边计算，但k为镜像
	return FindKthBit(n-1, (m<<1)-k) ^ 1
}

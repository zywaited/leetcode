package one

func FindKthBit(n int, k int) byte {
	bs := make([]byte, k)
	bs[0] = '0'
	i := 1
	for i < k {
		bs[i] = '1'
		i++
		for j := i - 2; j >= 0 && i < k; j-- {
			bs[i] = bs[j] ^ 1
			i++
		}
	}
	return bs[k-1]
}

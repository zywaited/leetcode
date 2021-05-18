package one

func CountTriplets(arr []int) int {
	xor := make([]int, len(arr)+1)
	for index, num := range arr {
		xor[index+1] = xor[index] ^ num
	}
	ans := 0
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if xor[i] == xor[j+1] {
				ans += j - i
			}
		}
	}
	return ans
}

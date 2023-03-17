package one

func BeautifulSubarrays(nums []int) int64 {
	ans := int64(0)
	bitNum := map[int]int64{0: 1}
	bit := 0
	for _, num := range nums {
		bit ^= num
		ans += bitNum[bit]
		bitNum[bit]++
	}
	return ans
}

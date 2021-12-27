package one

func TotalHammingDistance(nums []int) int {
	ans := 0
	bs := [32][2]int{}
	for _, num := range nums {
		for i := 31; i >= 0; i-- {
			ans += bs[i][num>>uint(i)&1^1]
			bs[i][num>>uint(i)&1]++
		}
	}
	return ans
}

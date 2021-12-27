package one

func CheckSubarraySum(nums []int, k int) bool {
	sumIndex := map[int]int{}
	sum := 0
	for index, num := range nums {
		sum = (sum + num) % k
		if (sum == 0 && index > 0) || (sumIndex[sum] > 0 && index-sumIndex[sum] > 0) {
			return true
		}
		if sumIndex[sum] == 0 {
			sumIndex[sum] = index + 1
		}
	}
	return false
}

package one

import(
	"math"
)

func maxSubarraySum(nums []int, k int) int64 {
    type maxSum struct {
        sum int64
        max int64
    }
    maxSums := make([]maxSum, len(nums)+1)
    max := int64(math.MinInt64)
    sum := int64(0)
    for index, num := range nums {
        sum += int64(num)
        curr := int64(0)
        prevIndex := index+1-k
        if prevIndex >= 0 {
            curr = sum-maxSums[prevIndex].sum
            if maxSums[prevIndex].max > 0 {
                curr += maxSums[prevIndex].max
            }
            if curr > max {
                max = curr
            }
        }
        maxSums[index+1] = maxSum{sum: sum, max: curr}
    }
    return max
}
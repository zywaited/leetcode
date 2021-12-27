package _260

import (
	"leetcode/1_500/251_300/260/one"
	"testing"

	"github.com/stretchr/testify/require"
)

/**
 * 先把数组所有的数字进行异或处理，因为有两个数不同，所以一定不会为0，即这两个数二进制一定有一位值不同
 * 就利用这不同的一位把数组分成两个数组分别进行异或处理，这样就得到这两个数了
 */

func TestSingleNumber(t *testing.T) {
	nums := one.SingleNumber([]int{1, 2, 1, 3, 2, 5})
	expected := []int{3, 5}
	require.Len(t, nums, len(expected))
	for _, num := range expected {
		require.Contains(t, nums, num)
	}
}

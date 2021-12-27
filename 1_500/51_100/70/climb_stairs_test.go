package _70

import (
	"leetcode/1_500/51_100/70/one"
	"leetcode/1_500/51_100/70/three"
	"leetcode/1_500/51_100/70/two"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestClimbStairs(t *testing.T) {
	for n := 1; n <= 200; n++ {
		nums := one.ClimbStairs(n)
		require.Equal(t, nums, two.ClimbStairs(n), n)
		require.Equal(t, nums, three.ClimbStairs(n), n)
	}
}

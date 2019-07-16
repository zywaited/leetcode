package _76

import (
	"leetcode/51_100/76/one"

	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinWindow(t *testing.T) {
	require.Equal(t, "BANC", one.MinWindow("ADOBECODEBANC", "ABC"))
}

package _458

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/451_500/458/one"
)

func TestPoorPigs(t *testing.T) {
	require.Equal(t, 5, one.PoorPigs(1000, 15, 60))
}

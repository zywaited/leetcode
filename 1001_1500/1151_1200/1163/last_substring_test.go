package _1163

import (
	"leetcode/1001_1500/1151_1200/1163/one"
	"leetcode/1001_1500/1151_1200/1163/two"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLastSubstringTest(t *testing.T) {
	expects := []string{
		"bab",
		"tcode",
		"cac",
		"aaaaaaaa",
	}
	ins := []string{
		"abab",
		"leetcode",
		"bcac",
		"aaaaaaaa",
	}
	for i, in := range ins {
		require.Equal(t, expects[i], one.LastSubstring(in))
		require.Equal(t, expects[i], two.LastSubstring(in))
		require.Equal(t, expects[i], two.AnotherLastSubstring(in))
	}
}

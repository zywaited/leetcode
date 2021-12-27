package _1125

import (
	"leetcode/1001_1500/1101_1150/1125/one"
	"leetcode/1001_1500/1101_1150/1125/three"
	"leetcode/1001_1500/1101_1150/1125/two"
	"testing"

	"github.com/stretchr/testify/require"
)

type smallestSufficientTeam func([]string, [][]string) []int

func test(t *testing.T, fn smallestSufficientTeam) {
	finds := fn(
		[]string{"java", "nodejs", "reactjs"},
		[][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}},
	)
	require.Len(t, finds, 2)
	for _, num := range []int{0, 2} {
		require.Contains(t, finds, num)
	}

	finds = fn(
		[]string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
		[][]string{
			{"algorithms", "math", "java"},
			{"algorithms", "math", "reactjs"},
			{"java", "csharp", "aws"},
			{"reactjs", "csharp"},
			{"csharp", "math"},
			{"aws", "java"},
		},
	)
	for _, num := range []int{1, 2} {
		require.Contains(t, finds, num)
	}

	finds = fn(
		[]string{"gvp", "jgpzzicdvgxlfix", "kqcrfwerywbwi", "jzukdzrfgvdbrunw", "k"},
		[][]string{
			{}, {}, {}, {},
			{"jgpzzicdvgxlfix"},
			{"jgpzzicdvgxlfix", "k"},
			{"jgpzzicdvgxlfix", "kqcrfwerywbwi"},
			{"gvp"},
			{"jzukdzrfgvdbrunw"},
			{"gvp", "kqcrfwerywbwi"},
		},
	)
	require.Len(t, finds, 3)
	for _, num := range []int{5, 8, 9} {
		require.Contains(t, finds, num)
	}

	finds = fn(
		[]string{"wieaul", "cxota", "frq", "knngtpip", "junne", "ctniuowumlcrhh"},
		[][]string{
			{"wieaul"},
			{"wieaul", "frq", "junne", "ctniuowumlcrhh"},
			{"wieaul"},
			{},
			{"knngtpip"},
			{},
			{"frq", "knngtpip"},
			{"cxota"},
			{"ctniuowumlcrhh"},
			{},
		},
	)
	require.Len(t, finds, 3)
	//for _, num := range []int{1, 4, 7} {
	//	require.Contains(t, finds, num)
	//}
}

func TestSmallestSufficientTeam(t *testing.T) {
	test(t, one.SmallestSufficientTeam)
	test(t, two.SmallestSufficientTeam)
	test(t, three.SmallestSufficientTeam)
	test(t, three.OptimizeSmallestSufficientTeam)
}

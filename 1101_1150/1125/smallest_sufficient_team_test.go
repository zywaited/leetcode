package one

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/1101_1150/1125/one"
)

func TestSmallestSufficientTeam(t *testing.T) {
	for _, num := range []int{0, 2} {
		require.Contains(
			t,
			one.SmallestSufficientTeam(
				[]string{"java", "nodejs", "reactjs"},
				[][]string{{"java"}, {"nodejs"}, {"nodejs", "reactjs"}},
			),
			num,
		)
	}

	for _, num := range []int{1, 2} {
		require.Contains(
			t,
			one.SmallestSufficientTeam(
				[]string{"algorithms", "math", "java", "reactjs", "csharp", "aws"},
				[][]string{
					{"algorithms", "math", "java"},
					{"algorithms", "math", "reactjs"},
					{"java", "csharp", "aws"},
					{"reactjs", "csharp"},
					{"csharp", "math"},
					{"aws", "java"},
				},
			),
			num,
		)
	}

	for _, num := range []int{5, 8, 9} {
		require.Contains(
			t,
			one.SmallestSufficientTeam(
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
			),
			num,
		)
	}

	for _, num := range []int{1, 4, 7} {
		require.Contains(
			t,
			one.SmallestSufficientTeam(
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
			),
			num,
		)
	}
}

package _149

import (
	"testing"

	"github.com/stretchr/testify/require"

	"leetcode/101_150/149/one"
)

func TestMaxPoints(t *testing.T) {
	// NOTE 找的评论区里的测试用例
	type args struct {
		points [][]int // 参数结构
	}
	tests := []struct {
		name string // 测试名称
		args args   // 传入参数
		want int    // 预期数据
	}{
		{name: "没有点", args: args{points: [][]int{}}, want: 0},
		{name: "只有一个点", args: args{points: [][]int{
			{1, 1},
		}}, want: 1},
		{name: "只有两个点", args: args{points: [][]int{
			{1, 1},
			{0, 0},
		}}, want: 2},
		{name: "有两个(1,1)的点", args: args{points: [][]int{
			{1, 1},
			{0, 0},
			{1, 1},
		}}, want: 3},
		{name: "有两个点存在重复的情况(1,1), (0,0)", args: args{points: [][]int{
			{1, 1},
			{0, 0},
			{1, 1},
			{0, 0},
			{2, 3},
		}}, want: 4},
		{name: "正常测试1", args: args{points: [][]int{
			{1, 1},
			{2, 2},
			{3, 3},
		}}, want: 3},
		{name: "正常测试2", args: args{points: [][]int{
			{1, 1},
			{3, 2},
			{5, 3},
			{4, 1},
			{2, 3},
			{1, 4},
		}}, want: 4},
		{name: "三个点相同的点", args: args{points: [][]int{
			{1, 1},
			{1, 1},
			{1, 1},
		}}, want: 3},
		{name: "斜率为0", args: args{points: [][]int{
			{2, 3},
			{3, 3},
			{-5, 3},
			{2, 1},
		}}, want: 3},
		{name: "斜率为无穷大", args: args{points: [][]int{
			{3, -1},
			{3, 2},
			{3, 1},
			{2, 1},
		}}, want: 3},
		{name: "k除不尽", args: args{points: [][]int{
			{84, 250},
			{0, 0},
			{1, 0},
			{0, -70},
			{0, -70},
			{1, -1},
			{21, 10},
			{42, 90},
			{-42, -230},
		}}, want: 6},
		{name: "浮点数精度丢失", args: args{points: [][]int{
			{94911152, 94911151},
			{0, 0},
			{94911151, 94911150},
		}}, want: 2},
	}
	for _, test := range tests {
		require.Equal(t, test.want, one.MaxPoints(test.args.points), test.name)
	}
}

package one

func CanEat(candiesCount []int, queries [][]int) []bool {
	sum := []int{0}
	for _, c := range candiesCount {
		sum = append(sum, sum[len(sum)-1]+c)
	}
	ans := make([]bool, len(queries))
	for i, q := range queries {
		ans[i] = q[1]+1 <= sum[q[0]+1] && (q[1]*q[2] >= sum[q[0]] || sum[q[0]]-q[1]*q[2] < q[2])
	}
	return ans
}

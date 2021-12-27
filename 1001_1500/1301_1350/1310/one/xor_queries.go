package one

func XorQueries(arr []int, queries [][]int) []int {
	l := make([]int, len(arr)+1)
	for index, num := range arr {
		l[index+1] = l[index] ^ num
	}
	ans := make([]int, len(queries))
	for index, query := range queries {
		ans[index] = l[query[1]+1] ^ l[query[0]]
	}
	return ans
}

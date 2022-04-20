package one

func LengthLongestPath(input string) int {
	var sums []int
	ans := 0
	index := 0
	for index < len(input) {
		if input[index] == '\n' {
			index++
		}
		tab := 0
		for ; index < len(input) && input[index] == '\t'; index++ {
			tab++
		}
		sums = sums[:tab]
		length := 0
		file := false
		for ; index < len(input) && input[index] != '\n'; index++ {
			length++
			file = file || input[index] == '.'
		}
		if len(sums) > 0 {
			length += sums[len(sums)-1] + 1
		}
		sums = append(sums, length)
		if file && ans < sums[len(sums)-1] {
			ans = sums[len(sums)-1]
		}
	}
	return ans
}

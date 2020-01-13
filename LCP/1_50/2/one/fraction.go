package one

func Fraction(cont []int) []int {
	ans := []int{cont[len(cont)-1], 1}
	for index := len(cont) - 2; index >= 0; index-- {
		if ans[1]%ans[0] == 0 {
			ans[0], ans[1] = cont[index]+ans[1]/ans[0], 1
			continue
		}
		ans[0], ans[1] = cont[index]*ans[0]+ans[1], ans[0]
	}
	return ans
}

package one

func MaxTurbulenceSize(arr []int) int {
	ans := 1
	f := false
	l := 1
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			l = 0
			continue
		}
		if f {
			if (i&1 == 0 && arr[i] < arr[i+1]) || (i&1 == 1 && arr[i] > arr[i+1]) {
				l++
				ans = max(ans, l)
				continue
			}
			l = 2
			f = false
			ans = max(ans, l)
			continue
		}
		if (i&1 == 0 && arr[i] > arr[i+1]) || (i&1 == 1 && arr[i] < arr[i+1]) {
			l++
			ans = max(ans, l)
			continue
		}
		l = 2
		f = true
		ans = max(ans, l)
		continue
	}
	return ans
}

func max(f, s int) int {
	if f >= s {
		return f
	}
	return s
}

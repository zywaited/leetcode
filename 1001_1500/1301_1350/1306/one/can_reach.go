package one

func CanReach(arr []int, start int) bool {
	used := make(map[int]byte)
	queue := []int{start}
	for len(queue) > 0 {
		index := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if arr[index] == 0 {
			return true
		}
		// 往后
		if index+arr[index] < len(arr) && used[index+arr[index]] == 0 {
			queue = append(queue, index+arr[index])
			used[index+arr[index]] = 1
		}
		// 往前
		if index-arr[index] >= 0 && used[index-arr[index]] == 0 {
			queue = append(queue, index-arr[index])
			used[index-arr[index]] = 1
		}
	}
	return false
}

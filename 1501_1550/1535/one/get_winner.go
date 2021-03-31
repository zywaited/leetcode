package one

func GetWinner(arr []int, k int) int {
	wn := arr[0]
	kn := 0
	for i := 1; i < len(arr); i++ {
		if kn == k {
			break
		}
		if arr[i] > wn {
			wn = arr[i]
			kn = 1
			continue
		}
		kn++
	}
	// 走一回合后剩下的那个数一定是赢家
	return wn
}

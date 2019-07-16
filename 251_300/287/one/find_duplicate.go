package one

func FindDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[slow]
	// 第一次相遇
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// 快指针是慢指针两倍速度
	// 数学可证明快指针从头一步一步走等于慢指针在在环形行走n圈后加上第一次相遇点到环形起点的距离和
	// 第一次相遇慢指针一定未走完一圈
	// 设起点都环形起点的距离为SA
	// 设第一次相遇点到环形起点距离为X
	// 设环形长度为R
	// 设第一次相遇点到环形起点剩余距离为Y (R - X)

	// 第一次相遇：
	// 慢指针走的距离为S = SA + X
	// 快指针走的距离为2S = SA + N*R + X (快指针可能走了好几圈)

	// 代入
	// 2*(SA + X)= SA + N*R + X
	// SA = N*R - X = (N-1)*R + R - X = (N-1)*R + Y (所以慢指针走了N圈后与快指针在起点相遇)

	fast = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}

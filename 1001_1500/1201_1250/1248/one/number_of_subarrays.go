package one

import "container/list"

func NumberOfSubarrays(nums []int, k int) int {
	queue := list.New()
	// 保留满足K个的时候前面可舍弃长度
	prev := 0
	ans := 0
	odd := 0
	for _, num := range nums {
		odd += num & 1
		queue.PushBack(num & 1)
		if odd < k {
			continue
		}
		ans++
		if odd > k {
			prev = 0
			// 去除第一个
			queue.Remove(queue.Front())
			odd--
		}
		// 计算prev的值
		for queue.Front() != nil {
			if queue.Front().Value.(int) == 0 {
				prev++
				queue.Remove(queue.Front())
				continue
			}
			break
		}
		ans += prev
	}
	return ans
}

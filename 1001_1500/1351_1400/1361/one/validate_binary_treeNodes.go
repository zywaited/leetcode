package one

func ValidateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	// 找根节点(题目没有说0就是根节点，最开始理解出错)
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		if leftChild[i] != -1 {
			nums[leftChild[i]]++
		}
		if rightChild[i] != -1 {
			nums[rightChild[i]]++
		}
	}
	// 如果root有多个就代表是两棵树(这里就不循环所有了，最后的判断会处理)
	root := -1
	for i := 0; i < n; i++ {
		if nums[i] == 0 {
			root = i
			break
		}
	}
	if root == -1 {
		return false
	}
	queue := make([]int, 0, n)
	queue = append(queue, root)
	used := make(map[int]byte) // 用过的节点
	used[root] = 1
	for len(queue) > 0 {
		i := queue[0]
		queue = queue[1:]
		used[i] = 1
		if leftChild[i] != -1 {
			if used[leftChild[i]] == 1 {
				return false
			}
			used[leftChild[i]] = 1
			queue = append(queue, leftChild[i])
		}
		if rightChild[i] != -1 {
			if used[rightChild[i]] == 1 {
				return false
			}
			used[rightChild[i]] = 1
			queue = append(queue, rightChild[i])
		}
	}
	// 节点数与实际一致才行
	return len(used) == n
}

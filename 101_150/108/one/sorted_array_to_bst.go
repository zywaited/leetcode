package one

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	node := &TreeNode{Val: nums[len(nums)>>1]}
	node.Left = SortedArrayToBST(nums[:len(nums)>>1])
	node.Right = SortedArrayToBST(nums[len(nums)>>1+1:])
	return node
}
